package stub

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var (
	addr               = flag.String("addr", "localhost:50051", "the address to connect to")
	RPC                *grpc.ClientConn
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing option metadata")
	err                error
)

func ConnectRPC() (*grpc.ClientConn, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Provide the correct paths to your TLS certificate file and the server's hostname
	creds, err := credentials.NewClientTLSFromFile(filepath.Join(currentDir, "certs", "server", "cert.pem"), "localhost")
	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	// Connect to the gRPC server
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
