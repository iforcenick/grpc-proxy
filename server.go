package main

import (
	"context"

	tmservice "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
)

type grpcServer struct {
	tmservice.UnimplementedServiceServer
	client *tmservice.ServiceClient
}

func (s *grpcServer) GetNodeInfo(ctx context.Context, req *tmservice.GetNodeInfoRequest) (*tmservice.GetNodeInfoResponse, error) {
	return (*s.client).GetNodeInfo(ctx, req)
}
func (s *grpcServer) GetSyncing(ctx context.Context, req *tmservice.GetSyncingRequest) (*tmservice.GetSyncingResponse, error) {
	return (*s.client).GetSyncing(ctx, req)
}
func (s *grpcServer) GetLatestBlock(ctx context.Context, req *tmservice.GetLatestBlockRequest) (*tmservice.GetLatestBlockResponse, error) {
	return (*s.client).GetLatestBlock(ctx, req)
}

func (s *grpcServer) GetBlockByHeight(ctx context.Context, req *tmservice.GetBlockByHeightRequest) (*tmservice.GetBlockByHeightResponse, error) {
	return (*s.client).GetBlockByHeight(ctx, req)
}
func (s *grpcServer) GetLatestValidatorSet(ctx context.Context, req *tmservice.GetLatestValidatorSetRequest) (*tmservice.GetLatestValidatorSetResponse, error) {
	return (*s.client).GetLatestValidatorSet(ctx, req)
}
func (s *grpcServer) GetValidatorSetByHeight(ctx context.Context, req *tmservice.GetValidatorSetByHeightRequest) (*tmservice.GetValidatorSetByHeightResponse, error) {
	return (*s.client).GetValidatorSetByHeight(ctx, req)
}
