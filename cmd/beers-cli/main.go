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

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func printDefaults(cmds ...*flag.FlagSet) {
	for _, cmd := range cmds {
		fmt.Fprintf(cmd.Output(), "Usage of %s:\n", cmd.Name())
		cmd.PrintDefaults()
	}
}

func main() {
	port := fmt.Sprintf(":%s", getEnv("PORT", "8000"))
	cc, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc client: %v\n", err)
		os.Exit(1)
	}

	client := pb.NewBeersClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmdName := addCmd.String("name", "", "Beer name")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeCmdId := removeCmd.Int("id", -1, "Beer id")

	if len(os.Args) < 2 {
		printDefaults(addCmd, removeCmd)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
	case "remove":
		removeCmd.Parse(os.Args[2:])
	default:
		printDefaults(addCmd, removeCmd)
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
	} else if removeCmd.Parsed() {
		if *removeCmdId < 0 {
			printDefaults(removeCmd)
			os.Exit(1)
		}

		resp, err := remove(ctx, cc, client, removeCmdId)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("removed: %s\n", resp)
	}
}
