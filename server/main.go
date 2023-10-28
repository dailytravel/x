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

// server is used to implement helloworld.GreeterServer.
type server struct {
	account.UnimplementedAccountServer
}

func (s *server) Authorization(ctx context.Context, in *account.Request) (*account.Response, error) {
	log.Printf("Received: %v", in.Message)
	return &account.Response{Message: "Hello " + in.Message}, nil
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
