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

type BlockInfo struct {
	height int64
	hash   []byte
}

func main() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	fetch_chan := make(chan BlockInfo)
	block_queue := []BlockInfo{}

	go func() {
		for {
			select {
			case <-ticker.C:
				client := tmclient.Dial("localhost:8080")
				response, err := (*client).GetLatestBlock(context.Background(), &tmservice.GetLatestBlockRequest{})
				if err != nil {
					log.Fatal("Unable to get latest block")
					return
				}
				block_hash := response.BlockId.Hash
				block_height := response.Block.Header.Height
				fetch_chan <- BlockInfo{hash: block_hash, height: block_height}
			}
		}
	}()

	previous_block_height := 0
	for true {
		block := <-fetch_chan
		if block.height == int64(previous_block_height) {
			continue
		}

		previous_block_height = int(block.height)
		block_queue = append(block_queue, block)
		if len(block_queue) > 5 {
			block_queue = block_queue[1:]
		}

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
