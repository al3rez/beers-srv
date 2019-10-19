package grpc

import (
	"context"
	"fmt"

	"github.com/azbshiri/beers/pkg/adding"
	pb "github.com/azbshiri/beers/pkg/serializing/protobuf/beers"
)

func (srv *server) Add(ctx context.Context, r *pb.BeerAddRequest) (*pb.BeerAddResponse, error) {
	name := r.GetName()
	if name == "" {
		return &pb.BeerAddResponse{
			Status: pb.Status_BAD,
			ErrMsg: "name cannot be blank",
		}, fmt.Errorf("grpc add: name cannot be blank\n")
	}

	beer, err := srv.adder.AddBeer(adding.Beer{Name: name})
	if err != nil {
		return &pb.BeerAddResponse{
			Status: pb.Status_BAD,
			ErrMsg: err.Error(),
		}, fmt.Errorf("service add: %s\n", err)
	}

	return &pb.BeerAddResponse{
		Status: pb.Status_OK,
		Id:     int32(beer.Id),
		Name:   beer.Name,
	}, nil
}
