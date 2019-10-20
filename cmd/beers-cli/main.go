package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":80", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc client: %v\n", err)
		os.Exit(1)
	}

	client := pb.NewBeersClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmdName := addCmd.String("name", "", "Beer name")

	if len(os.Args) < 2 {
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addCmd.Parsed() {
		if *addCmdName == "" {
			addCmd.PrintDefaults()
			os.Exit(1)
		}

		resp, err := add(ctx, cc, client, addCmdName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("added: %s\n", resp)
	}
}
