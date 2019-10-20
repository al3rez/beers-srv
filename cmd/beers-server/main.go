package main

import (
	"fmt"
	"net"
	"os"

	"log"

	"github.com/azbshiri/beers/pkg/grpc"
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func main() {
	port := fmt.Sprintf(":%s", getEnv("PORT", "8000"))

	s := grpc.New()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %s\n", err)
		os.Exit(1)
	}

	log.Printf("running server on %s\n", port)
	log.Fatal(s.Serve(listener))
}
