package grpc

import (
	"context"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"github.com/azbshiri/beers/pkg/removing"
)

func (srv *server) Remove(ctx context.Context, r *pb.BeerRemoveRequest) (*pb.BeerRemoveResponse, error) {
	id := r.GetId()
	if id < 0 {
		return &pb.BeerRemoveResponse{
			Status: pb.Status_BAD,
			ErrMsg: "invalid id",
		}, nil
	}

	beer, err := srv.remover.RemoveBeer(removing.Beer{Id: int(id)})
	if err != nil {
		return &pb.BeerRemoveResponse{
			Status: pb.Status_BAD,
			ErrMsg: err.Error(),
		}, nil
	}

	return &pb.BeerRemoveResponse{
		Status: pb.Status_OK,
		Id:     int32(beer.Id),
		Name:   beer.Name,
	}, nil
}
