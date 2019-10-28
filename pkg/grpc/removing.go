package grpc

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/azbshiri/beers-proto/pkg/pb"
	"github.com/azbshiri/beers-srv/pkg/removing"
)

func (srv *server) Remove(ctx context.Context, r *pb.BeerRemoveRequest) (*pb.BeerRemoveResponse, error) {
	id := r.GetId()
	if id < 0 {
		return &pb.BeerRemoveResponse{
			Status: pb.Status_BAD,
			ErrMsg: "invalid id",
		}, nil
	}

	beersChan := make(chan *removing.Beer, 1)
	errors := hystrix.Go("beers", func() error {
		beer, err := srv.remover.RemoveBeer(removing.Beer{Id: int(id)})
		if err != nil {
			return err
		}

		beersChan <- beer
		return nil
	}, nil)

	select {
	case beer := <-beersChan:
		return &pb.BeerRemoveResponse{
			Status: pb.Status_OK,
			Id:     int32(beer.Id),
			Name:   beer.Name,
		}, nil
	case err := <-errors:
		return &pb.BeerRemoveResponse{
			Status: pb.Status_BAD,
			ErrMsg: err.Error(),
		}, nil
	}

}
