package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/azbshiri/beers/pkg/grpc/proto/beers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":80", grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("grpc client: %v\n", err)
		os.Exit(1)
	}

	client := beers.NewBeersClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &beers.BeerAddRequest{
		Name: "ali",
	}

	resp, err := client.Add(ctx, req)
	if err != nil {
		logrus.Errorf("grpc client: %v\n", err)
		os.Exit(1)
	}

	log.Println(resp)
}
