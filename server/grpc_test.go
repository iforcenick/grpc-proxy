package main

import (
	"context"
	"reflect"
	"testing"
	tmclient "tmclient"

	tmservice "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
)

// Test the equality of two responses from localhost and osmosis public endpoint *GetLatestBlock*
func TextGetLatestBlock(t *testing.T) {
	local_client := tmclient.Dial("localhost:8080")
	local_response, local_err := (*local_client).GetLatestBlock(context.Background(), &tmservice.GetLatestBlockRequest{})
	if local_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	remote_client := tmclient.Dial("grpc.osmosis.zone:9090")
	remote_response, remote_err := (*remote_client).GetLatestBlock(context.Background(), &tmservice.GetLatestBlockRequest{})
	if remote_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	if !reflect.DeepEqual(local_response, remote_response) {
		t.Errorf("Expected %v, got %v", remote_response, local_response)
	}
}

// Test the equality of two responses from localhost and osmosis public endpoint *GetBlockByHeight*
func TestGetBlock(t *testing.T) {
	local_client := tmclient.Dial("localhost:8080")
	local_response, local_err := (*local_client).GetBlockByHeight(context.Background(), &tmservice.GetBlockByHeightRequest{Height: 5850706})
	if local_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	remote_client := tmclient.Dial("grpc.osmosis.zone:9090")
	remote_response, remote_err := (*remote_client).GetBlockByHeight(context.Background(), &tmservice.GetBlockByHeightRequest{Height: 5850706})
	if remote_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	if !reflect.DeepEqual(local_response, remote_response) {
		t.Errorf("Expected %v, got %v", remote_response, local_response)
	}
}

// Test the equality of two responses from localhost and osmosis public endpoint *GetSyncing*
func TestGetSyncing(t *testing.T) {
	local_client := tmclient.Dial("localhost:8080")
	local_response, local_err := (*local_client).GetSyncing(context.Background(), &tmservice.GetSyncingRequest{})
	if local_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	remote_client := tmclient.Dial("grpc.osmosis.zone:9090")
	remote_response, remote_err := (*remote_client).GetSyncing(context.Background(), &tmservice.GetSyncingRequest{})
	if remote_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	if !reflect.DeepEqual(local_response, remote_response) {
		t.Errorf("Expected %v, got %v", remote_response, local_response)
	}
}

// Test the equality of two responses from localhost and osmosis public endpoint *GetNodeInfo*
func TestGetNodeInfo(t *testing.T) {
	local_client := tmclient.Dial("localhost:8080")
	local_response, local_err := (*local_client).GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	if local_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	remote_client := tmclient.Dial("grpc.osmosis.zone:9090")
	remote_response, remote_err := (*remote_client).GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	if remote_err != nil {
		t.Errorf("Error occured with %v", local_err)
	}

	if !reflect.DeepEqual(local_response, remote_response) {
		t.Errorf("Expected %v, got %v", remote_response, local_response)
	}
}
