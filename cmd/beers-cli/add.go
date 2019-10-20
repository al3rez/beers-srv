package main

import (
	"context"
	"fmt"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"google.golang.org/grpc"
)

func add(ctx context.Context, cc *grpc.ClientConn, client pb.BeersClient, name *string) (*pb.BeerAddResponse, error) {
	req := &pb.BeerAddRequest{
		Name: *name,
	}

	resp, err := client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("grpc client: %s\n", err)
	}

	return resp, nil
}
