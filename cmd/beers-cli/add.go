package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"google.golang.org/grpc"
)

func strip(str string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}

	return reg.ReplaceAllString(str, ""), nil
}

func add(ctx context.Context, cc *grpc.ClientConn, client pb.BeersClient, name *string) (*pb.BeerAddResponse, error) {
	n, err := strip(*name)
	if err != nil {
		return nil, fmt.Errorf("strip: %s\n", err)
	}

	req := &pb.BeerAddRequest{
		Name: n,
	}

	resp, err := client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("grpc client: %s\n", err)
	}

	return resp, nil
}
