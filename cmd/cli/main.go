package main

import (
	"context"
	"flag"
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

		add(ctx, cc, client, addCmdName)
	}
}
