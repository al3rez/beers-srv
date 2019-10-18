package main

import (
	"fmt"
	"net"
	"os"

	"github.com/azbshiri/beers/pkg/grpc"
	"github.com/prometheus/common/log"
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
