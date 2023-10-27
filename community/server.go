package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dailytravel/x/community/config"
	"github.com/dailytravel/x/community/graph"
	"github.com/dailytravel/x/community/pkg/auth"
	"github.com/dailytravel/x/community/pkg/database"
	"github.com/dailytravel/x/community/pkg/database/migrations"
	"github.com/dailytravel/x/community/pkg/queuing/consumer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "test-key", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
	lifetime     = flag.Duration("lifetime", 5*time.Second, "lifetime of process before shutdown (0s=infinite)")
)

func init() {
	flag.Parse()
	os.Setenv("GIN_MODE", "release")
	os.Setenv("PORT", "4004")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "community")
	os.Setenv("MONGODB_URI", "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.10.1")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Defining the playgroundHandler handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	resolver := graph.NewResolver(database.Database, database.Redis, database.Client)
	c := graph.Config{Resolvers: resolver}
	config.Directives(&c)

	executableSchema := graph.NewExecutableSchema(c)

	server := handler.NewDefaultServer(executableSchema)
	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	return func(c *gin.Context) {
		if strings.ToLower(c.Request.Header.Get("Upgrade")) == "websocket" {
			log.Println("Websocket request")
		}

		server.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	// Connect to MongoDB
	client, err := database.ConnectDB()
	failOnError(err, "Failed to connect to MongoDB")

	defer func() {
		err := client.Disconnect(context.Background())
		failOnError(err, "Failed to disconnect from MongoDB")
	}()

	// Initialize database connections
	database.Database = client.Database(os.Getenv("DB_NAME"))
	database.Redis = database.ConnectRedis()
	database.Client = database.ConnectTypesense()

	// Run database migrations
	err = migrations.AutoMigrate()
	failOnError(err, "Failed to run database migrations")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(auth.Middleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// Start the server and handle errors using a goroutine
	errCh := make(chan error)
	go func() {
		errCh <- r.Run(":" + os.Getenv("PORT"))
	}()

	// Wait for the server to start or throw an error
	err = <-errCh
	failOnError(err, "Failed to start server")

	c, err := consumer.NewConsumer(*uri, *exchange, *exchangeType, *queue, *bindingKey, *consumerTag)
	failOnError(err, "Failed to connect to RabbitMQ")

	if *lifetime > 0 {
		log.Printf("running for %s", *lifetime)
		time.Sleep(*lifetime)
	} else {
		log.Printf("running forever")
		select {}
	}

	log.Printf("shutting down")

	err = c.Shutdown()
	failOnError(err, "Failed to shutdown consumer")

	// Wait for the server to shutdown
	waitGroup.Wait()
}
