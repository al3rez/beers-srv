package grpc

import (
	"github.com/azbshiri/beers-proto/pkg/pb"
	"github.com/azbshiri/beers/pkg/adding"
	"github.com/azbshiri/beers/pkg/removing"
	"github.com/azbshiri/beers/pkg/storage/mem"
	"google.golang.org/grpc"
)

type server struct {
	adder   adding.Service
	remover removing.Service
}

func New() *grpc.Server {
	storage := mem.Storage{}
	addr := adding.NewService(&storage)
	rmvr := removing.NewService(&storage)
	s := grpc.NewServer()
	pb.RegisterBeersServer(s, &server{addr, rmvr})
	return s
}
