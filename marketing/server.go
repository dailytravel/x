package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dailytravel/x/marketing/config"
	"github.com/dailytravel/x/marketing/graph"
	"github.com/dailytravel/x/marketing/pkg/auth"
	"github.com/dailytravel/x/marketing/pkg/database"
	"github.com/dailytravel/x/marketing/pkg/database/migrations"
	"github.com/dailytravel/x/marketing/pkg/stub"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func init() {
	flag.Parse()
	os.Setenv("GIN_MODE", "release")
	os.Setenv("PORT", "4009")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "marketing")
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
	stub.RPC, err = stub.ConnectRPC()

	failOnError(err, "Failed to connect to RPC")

	err = migrations.AutoMigrate()
	failOnError(err, "Failed to migrate")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(auth.Middleware(stub.RPC))
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
}
