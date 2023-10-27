package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dailytravel/x/account/config"
	"github.com/dailytravel/x/account/graph"
	"github.com/dailytravel/x/account/internal/controllers"
	"github.com/dailytravel/x/account/pkg/auth"
	"github.com/dailytravel/x/account/pkg/database"
	"github.com/dailytravel/x/account/pkg/database/migrations"
	"github.com/dailytravel/x/account/pkg/queuing/producer"
	"github.com/dailytravel/x/account/scheduler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	routingKey   = flag.String("key", "test-key", "AMQP routing key")
	body         = flag.String("body", "foobar", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
)

func init() {
	flag.Parse()
	os.Setenv("GIN_MODE", "release")
	os.Setenv("PORT", "4001")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "account")
	os.Setenv("MONGODB_URI", "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.10.1")
	os.Setenv("ISSUER", "https://api.trip.express")
	os.Setenv("AUDIENCE", "https://api.trip.express/graphql")
	os.Setenv("JWKS_URI", "https://api.trip.express/.well-known/jwks.json")
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
	// connect MongoDB
	client, err := database.ConnectDB()
	failOnError(err, "Failed to connect to MongoDB")

	defer func() {
		err := client.Disconnect(context.Background())
		failOnError(err, "Failed to disconnect from MongoDB")
	}()

	database.Database = client.Database(os.Getenv("DB_NAME"))
	database.Redis = database.ConnectRedis()
	database.Client = database.ConnectTypesense()

	err = migrations.AutoMigrate()
	failOnError(err, "Failed to migrate database")

	database.Client.Collection("users").Delete()
	// start scheduler jobs
	scheduler.SyncUsersJob()

	// need restart the server if drop or create a new collection in mongodb, else will not work
	for _, name := range []string{} {
		stream, err := database.Database.Collection(name).Watch(context.Background(), mongo.Pipeline{})
		failOnError(err, "Failed to watch collection")
		waitGroup.Add(1)
		go controllers.IndexStream(&waitGroup, stream, name)
	}

	err = producer.Publish(*uri, *exchangeName, *exchangeType, *routingKey, []byte(*body), *reliable)
	failOnError(err, "Failed to publish a message")

	log.Printf("published %dB OK", len(*body))

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

	// Start the server with error handling using a simple channel
	errCh := make(chan error)
	go func() {
		errCh <- r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	}()

	// Wait for the server to start or throw an error
	err = <-errCh
	failOnError(err, "Failed to start server")

	// Wait for the waitGroup to finish
	waitGroup.Wait()
}
