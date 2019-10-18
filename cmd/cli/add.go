package main

import (
	"context"
	"fmt"
	"os"

	"github.com/azbshiri/beers/pkg/grpc/proto/beers"
	"google.golang.org/grpc"
)

func add(ctx context.Context, cc *grpc.ClientConn, client beers.BeersClient, name *string) {
	req := &beers.BeerAddRequest{
		Name: *name,
	}

	resp, err := client.Add(ctx, req)
	if err != nil {
		fmt.Printf("grpc client: %q\n", err)
		os.Exit(1)
	}

	fmt.Printf("added: %v\n", resp)
}
