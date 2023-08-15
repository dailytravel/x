package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dailytravel/x/community/auth"
	"github.com/dailytravel/x/community/config"
	"github.com/dailytravel/x/community/db"
	"github.com/dailytravel/x/community/db/migrations"
	"github.com/dailytravel/x/community/graph"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func init() {
	os.Setenv("GIN_MODE", "release")
	os.Setenv("PORT", "4003")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "community")
	os.Setenv("MONGODB_URI", "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.10.1")
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
	resolver := graph.NewResolver(db.Database, db.Redis, db.Client)
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
	client, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Failed to close MongoDB connection: ", err)
		}
	}()

	db.Database = client.Database(os.Getenv("DB_NAME"))
	db.Redis = db.ConnectRedis()
	db.Client = db.ConnectTypesense()

	if err := migrations.AutoMigrate(); err != nil {
		log.Fatal("Error running migrations: ", err)
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
}
