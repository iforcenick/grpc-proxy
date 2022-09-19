package main

import (
	"context"
	"encoding/json"
	"log"
	"time"
	tmclient "tmclient"

	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
)

// Block information retrieved from GetLatestBlock
type BlockInfo struct {
	height int64
	hash   []byte
}

// Program entrypoint
func main() {
	// Create a timer involed every 1 second
	ticker := time.NewTicker(1000 * time.Millisecond)

	// Connect to the gRPC server
	client := tmclient.Dial("localhost:8080")

	// Channel to pass data from fethcing thread to main thread
	fetch_chan := make(chan BlockInfo)

	// Block queue for 5 block reservation
	block_queue := []BlockInfo{}

	// Thread for fetching latest block info every 1 second
	// Pass retrieved info through fetch_chan
	go func() {
		for {
			select {
			case <-ticker.C:
				// Get latest block
				response, err := (*client).GetLatestBlock(context.Background(), &tmservice.GetLatestBlockRequest{})
				if err != nil {
					log.Fatal("Unable to get latest block")
					return
				}

				// Hash and height
				block_hash := response.BlockId.Hash
				block_height := response.Block.Header.Height

				// Pass data through channel
				fetch_chan <- BlockInfo{hash: block_hash, height: block_height}
			}
		}
	}()

	previous_block_height := 0
	for true {
		// Get fetched information
		block := <-fetch_chan

		// Check if new block is appended to the blockchain
		if block.height == int64(previous_block_height) {
			continue
		}

		previous_block_height = int(block.height)

		// Update block queue
		block_queue = append(block_queue, block)
		if len(block_queue) > 5 {
			block_queue = block_queue[1:]
		}

		// Save data in JSON format
		block_json := make([]map[string]interface{}, len(block_queue))
		for i := range block_queue {
			block_json[i] = map[string]interface{}{
				"height": block_queue[i].height,
				"hash":   block_queue[i].hash,
			}
		}
		buffer, err := json.MarshalIndent(map[string]interface{}{
			"test_result": block_json,
		}, "", " ")
		if err != nil {
			log.Fatalf("Unable to marshal json data: %v", err)
		}
		_ = ioutil.WriteFile("block_status.json", buffer, 0644)
		log.Print(buffer)
	}
}
