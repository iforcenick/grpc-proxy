package main

import (
	"log"
	"net"

	tmclient "tmclient"

	tmservice "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer() *grpc.Server {
	client := tmclient.Dial("grpc.osmosis.zone:9090")
	gserver := grpc.NewServer()
	srv := GrpcServer{
		client: client,
	}

	tmservice.RegisterServiceServer(gserver, &srv)
	return gserver
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed with error: %v", err)
	}
	log.Printf("Listening to port 8080")

	gserver := NewGRPCServer()
	reflection.Register(gserver)
	if err := gserver.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("gRPC server started.")
}
