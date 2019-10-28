package grpc

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/azbshiri/beers-proto/pkg/pb"
	v1 "github.com/azbshiri/beers-proto/pkg/pb/grpc_health_v1"
	"github.com/azbshiri/beers-srv/pkg/adding"
	"github.com/azbshiri/beers-srv/pkg/removing"
	"github.com/azbshiri/beers-srv/pkg/storage/mem"
	"google.golang.org/grpc"
)

type server struct {
	adder   adding.Service
	remover removing.Service
}

func New() *grpc.Server {
	hystrix.ConfigureCommand("beers", hystrix.CommandConfig{
		SleepWindow:            5000,
		RequestVolumeThreshold: 10,
	})

	storage := mem.Storage{}
	addr := adding.NewService(&storage)
	rmvr := removing.NewService(&storage)
	s := grpc.NewServer()
	srv := &server{addr, rmvr}
	pb.RegisterBeersServer(s, srv)
	v1.RegisterHealthServer(s, srv)
	return s
}
