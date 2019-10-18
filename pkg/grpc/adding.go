package grpc

import (
	"context"

	"github.com/azbshiri/beers/pkg/adding"
	pb "github.com/azbshiri/beers/pkg/grpc/proto/beers"
)

func (srv *server) Add(ctx context.Context, r *pb.BeerAddRequest) (*pb.BeerAddResponse, error) {
	name := r.GetName()
	if name == "" {
		return &pb.BeerAddResponse{
			Status: pb.Status_OK,
			ErrMsg: "Name cannot be blank",
		}, nil
	}

	beer, err := srv.adder.AddBeer(adding.Beer{Name: name})
	if err != nil {
		return nil, err
	}

	return &pb.BeerAddResponse{
		Status: pb.Status_OK,
		Id:     int32(beer.Id),
		Name:   beer.Name,
	}, nil
}
