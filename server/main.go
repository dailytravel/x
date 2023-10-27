package server

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port = flag.Int("port", 50051, "the port to serve on")
)

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("", "")
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	s := grpc.NewServer(opts...)

	// user.RegisterUserServer(s, &userService.Server{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		panic(err)
	}

	fmt.Printf("server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
