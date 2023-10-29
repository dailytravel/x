package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/dailytravel/x/proto/account"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	Database *mongo.Database
	port     = flag.Int("port", 50051, "The server port")
)

func ConnectDB() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Check the connection by pinging the server
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return client, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	account.UnimplementedAccountServer
}

func (s *server) Authorization(ctx context.Context, in *account.Request) (*account.Response, error) {
	// Connect to MongoDB
	client, err := ConnectDB()
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// Convert the input message (token ID) to an ObjectID
	id, err := primitive.ObjectIDFromHex(in.Message)
	if err != nil {
		log.Printf("Failed to convert to ObjectID: %v", err)
		return nil, err
	}

	// Find the token by ID
	var token map[string]interface{}
	err = client.Database("accounts").Collection("tokens").FindOne(context.Background(), primitive.M{"_id": id}).Decode(&token)
	if err != nil {
		log.Printf("Failed to find token: %v", err)
		return nil, err
	}

	// Find the user by user ID from the token
	var user map[string]interface{}
	err = client.Database("accounts").Collection("users").FindOne(context.Background(), primitive.M{"_id": token["uid"]}).Decode(&user)
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		return nil, err
	}

	// Construct the response message
	response := &account.Response{
		Status:  "authenticated",
		Message: fmt.Sprintf("Hello %s", user["name"]),
	}

	return response, nil
}

func main() {
	flag.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Load your TLS certificate and key
	cert, err := tls.LoadX509KeyPair(
		filepath.Join(currentDir, "certs", "x509", "server_cert.pem"), filepath.Join(currentDir, "certs", "x509", "server_key.pem"))
	failOnError(err, "Failed to load key pair")

	// Create a new gRPC server with TLS credentials
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	creds := credentials.NewTLS(tlsConfig)

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	failOnError(err, "Failed to listen")

	s := grpc.NewServer(opts...)
	account.RegisterAccountServer(s, &server{})
	fmt.Printf("gRPC server listening at %v\n", lis.Addr())

	err = s.Serve(lis)
	failOnError(err, "Failed to serve")
}
