package tmclient

import (
	"log"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
)

// Create gRPC client and connect to the server
func Dial(addr string) *tmservice.ServiceClient {

	// Dial to the server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while connecting %v", err)
		return nil
	}
	// Create tmservice client from the connected client
	client := tmservice.NewServiceClient(conn)
	return &client
}
