package grpc

import (
	"github.com/azbshiri/beers/pkg/adding"
	"github.com/azbshiri/beers/pkg/grpc/proto/beers"
	"github.com/azbshiri/beers/pkg/storage/mem"
	"google.golang.org/grpc"
)

type server struct {
	adder adding.Service
}

func New() *grpc.Server {
	storage := mem.Storage{}
	srv := adding.NewService(&storage)
	s := grpc.NewServer()
	beers.RegisterBeersServer(s, &server{srv})
	return s
}
