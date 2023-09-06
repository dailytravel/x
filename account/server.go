package main

import (
	"context"
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
	"github.com/dailytravel/x/account/graph"
	"github.com/dailytravel/x/account/internal/controllers"
	"github.com/dailytravel/x/account/pkg/auth"
	"github.com/dailytravel/x/account/pkg/config"
	"github.com/dailytravel/x/account/pkg/database"
	"github.com/dailytravel/x/account/pkg/database/migrations"
	"github.com/dailytravel/x/account/scheduler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	os.Setenv("GIN_MODE", "release")
	os.Setenv("PORT", "4001")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "account")
	os.Setenv("MONGODB_URI", "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.10.1")
	os.Setenv("ISSUER", "https://api.trip.express")
	os.Setenv("AUDIENCE", "https://api.trip.express/graphql")
	os.Setenv("JWKS_URI", "https://api.trip.express/.well-known/jwks.json")
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
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Failed to close MongoDB connection: ", err)
		}
	}()

	database.Database = client.Database(os.Getenv("DB_NAME"))
	database.Redis = database.ConnectRedis()
	database.Client = database.ConnectTypesense()

	if err := migrations.AutoMigrate(); err != nil {
		log.Fatal("Error running migrations: ", err)
	}

	database.Client.Collection("users").Delete()
	// start scheduler jobs
	scheduler.SyncUsersJob()

	// need restart the server if drop or create a new collection in mongodb, else will not work
	for _, name := range []string{} {
		stream, err := database.Database.Collection(name).Watch(context.Background(), mongo.Pipeline{})
		if err != nil {
			panic(err)
		}
		waitGroup.Add(1)
		go controllers.IndexStream(&waitGroup, stream, name)
	}

	// setting up Gin
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
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	// Wait for the waitGroup to finish
	waitGroup.Wait()
}
