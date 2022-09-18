package main

import (
	"log"

	tmservice "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
)

func dial(addr string) *tmservice.ServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while connecting %v", err)
		return nil
	}
	client := tmservice.NewServiceClient(conn)
	return &client
}
