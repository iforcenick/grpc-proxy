package main

import (
	"context"
	"log"
	tmclient "tmclient"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
)

func main() {
	client := tmclient.Dial("localhost:8080")
	response, err := (*client).GetLatestBlock(context.Background(), &tmservice.GetLatestBlockRequest{})
	if err != nil {
		log.Fatal("Unable to get latest block")
		return
	}
	block_hash := response.BlockId.Hash
	block_height := response.Block.Header.Height

	log.Print(block_hash, block_height)
}
