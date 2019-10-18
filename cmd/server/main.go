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
		log.Errorf("failed to listen: %v", err)
		os.Exit(1)
	}
	log.Fatal(s.Serve(listener))
}
