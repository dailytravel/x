package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/dailytravel/x/certs"
	"github.com/dailytravel/x/proto/account"
	"github.com/dailytravel/x/proto/activity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ConnectDB() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.10.1").SetServerAPIOptions(serverAPI)
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

func (s *server) Authorization(ctx context.Context, in *account.Request) (*account.Response, error) {
	// Convert the input message (token ID) to an ObjectID
	id, err := primitive.ObjectIDFromHex(in.Message)
	if err != nil {
		log.Printf("Failed to convert to ObjectID: %v", err)
		return nil, err
	}

	// Find the token by ID
	result := s.db.Database("account").Collection("tokens").FindOneAndUpdate(ctx,
		primitive.M{"_id": id, "revoked": false, "expires": bson.M{"$gt": primitive.Timestamp{T: uint32(time.Now().Unix())}}},
		bson.M{"$set": bson.M{"last_used": primitive.Timestamp{T: uint32(time.Now().Unix())}}})

	if result.Err() != nil {
		log.Printf("Failed to find and update token: %v", result.Err())
		return nil, result.Err()
	}

	var token map[string]interface{}
	if err := result.Decode(&token); err != nil {
		log.Printf("Failed to decode result: %v", err)
		return nil, err
	}

	// Find the user by user ID from the token
	err = s.db.Database("account").Collection("users").FindOne(ctx, primitive.M{"_id": token["uid"]}).Err()
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		return nil, err
	}

	// Construct the response message
	response := &account.Response{
		Status: "authenticated",
	}

	return response, nil
}

func (s *server) Activity(ctx context.Context, in *activity.Request) (*activity.Response, error) {
	uid, err := primitive.ObjectIDFromHex(in.Uid)
	if err != nil {
		log.Printf("Failed to convert to ObjectID: %v", err)
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(in.ObjectId)
	if err != nil {
		log.Printf("Failed to convert to ObjectID: %v", err)
		return nil, err
	}

	doc := bson.M{
		"uid":         uid,
		"action":      in.Action,
		"object._id":  objectId,
		"object.type": in.ObjectType,
		"timestamp":   primitive.Timestamp{T: uint32(in.Timestamp)},
	}

	if in.TargetId != "" {
		targetId, err := primitive.ObjectIDFromHex(in.TargetId)
		if err != nil {
			log.Printf("Failed to convert to ObjectID: %v", err)
			return nil, err
		}
		doc["target._id"] = targetId
		doc["target.type"] = in.TargetType
	}

	res, err := s.db.Database("insight").Collection("activities").InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Failed to insert activity: %v", err)
		return nil, err
	}

	response := &activity.Response{
		Status:  "success",
		Message: fmt.Sprintf("Inserted %v documents into activity collection", res.InsertedID),
	}

	return response, nil
}

type server struct {
	db *mongo.Client
	account.UnimplementedAccountServer
	activity.UnimplementedActivityServer
}

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(certs.Path("x509/server_cert.pem"), certs.Path("x509/server_key.pem"))
	failOnError(err, "Failed to load server certificates")

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	creds := credentials.NewTLS(tlsConfig)

	client, err := ConnectDB()
	failOnError(err, "Failed to connect to MongoDB")
	defer func() {
		failOnError(client.Disconnect(context.Background()), "Failed to disconnect from MongoDB")
	}()

	s := grpc.NewServer(grpc.Creds(creds))
	account.RegisterAccountServer(s, &server{db: client})
	activity.RegisterActivityServer(s, &server{db: client})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	failOnError(err, "Failed to listen")
	defer lis.Close()

	fmt.Printf("gRPC server listening at %v\n", lis.Addr())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		errCh := make(chan error)
		errCh <- s.Serve(lis)

		select {
		case err := <-errCh:
			failOnError(err, "Failed to serve")
		case <-ctx.Done():
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	select {
	case <-signalChan:
		cancel()
	case <-ctx.Done():
	}
}
