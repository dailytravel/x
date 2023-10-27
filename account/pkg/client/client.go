package client

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func ConnectRPC() (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile("", "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		//grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	}

	driver, err := grpc.Dial(*addr, opts...)

	if err != nil {
		return nil, err
	}

	return driver, nil
}
