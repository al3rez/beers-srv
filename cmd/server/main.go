package main

import (
	"fmt"
	"net"
	"os"

	"log"

	"github.com/azbshiri/beers/pkg/grpc"
)

func main() {
	s := grpc.New()
	listener, err := net.Listen("tcp", fmt.Sprintf(":80"))
	if err != nil {
		fmt.Printf("failed to listen: %s\n", err)
		os.Exit(1)
	}
	log.Fatal(s.Serve(listener))
}
