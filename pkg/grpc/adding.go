package grpc

import (
	"context"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"github.com/azbshiri/beers-srv/pkg/adding"
)

func (srv *server) Add(ctx context.Context, r *pb.BeerAddRequest) (*pb.BeerAddResponse, error) {
	name := r.GetName()
	if name == "" {
		return &pb.BeerAddResponse{
			Status: pb.Status_BAD,
			ErrMsg: "name cannot be blank",
		}, nil
	}

	beer, err := srv.adder.AddBeer(adding.Beer{Name: name})
	if err != nil {
		return &pb.BeerAddResponse{
			Status: pb.Status_BAD,
			ErrMsg: err.Error(),
		}, nil
	}

	return &pb.BeerAddResponse{
		Status: pb.Status_OK,
		Id:     int32(beer.Id),
		Name:   beer.Name,
	}, nil
}
