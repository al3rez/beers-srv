package main

import (
	"context"
	"fmt"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"google.golang.org/grpc"
)

func remove(ctx context.Context, cc *grpc.ClientConn, client pb.BeersClient, id *int) (*pb.BeerRemoveResponse, error) {
	req := &pb.BeerRemoveRequest{
		Id: int32(*id),
	}

	resp, err := client.Remove(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("grpc client: %s\n", err)
	}

	return resp, nil
}
