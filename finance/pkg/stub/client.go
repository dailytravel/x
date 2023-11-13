package stub

import (
	"flag"
	"log"

	"github.com/dailytravel/x/certs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	RPC  *grpc.ClientConn
)

func ConnectRPC() (*grpc.ClientConn, error) {

	// Provide the correct paths to your TLS certificate file and the server's hostname
	creds, err := credentials.NewClientTLSFromFile(certs.Path("x509/ca_cert.pem"), "x.trip.express")
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
