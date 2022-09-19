package main

import (
	"log"
	"net"

	tmclient "tmclient"

	tmservice "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Create gRPC server and register endpoints
func NewGRPCServer() *grpc.Server {
	// Dial to the Osmosis mainnet public endpoint
	client := tmclient.Dial("grpc.osmosis.zone:9090")
	gserver := grpc.NewServer()
	srv := GrpcServer{
		client: client,
	}

	// Register gRPC endpoints
	tmservice.RegisterServiceServer(gserver, &srv)
	return gserver
}

// Program entrypoint - Create server and listen
func main() {
	// Create listening tcp port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed with error: %v", err)
	}
	log.Printf("Listening to port 8080")

	// Bootstrap gRPC server on the listening port
	gserver := NewGRPCServer()
	reflection.Register(gserver)
	if err := gserver.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("gRPC server started.")
}
