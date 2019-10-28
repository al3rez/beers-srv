package grpc

import (
	"context"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"github.com/azbshiri/beers-srv/pkg/adding"

	"github.com/afex/hystrix-go/hystrix"
)

func (srv *server) Add(ctx context.Context, r *pb.BeerAddRequest) (*pb.BeerAddResponse, error) {
	name := r.GetName()
	if name == "" {
		return &pb.BeerAddResponse{
			Status: pb.Status_BAD,
			ErrMsg: "name cannot be blank",
		}, nil
	}

	beersChan := make(chan *adding.Beer, 1)
	errors := hystrix.Go("beers", func() error {
		beer, err := srv.adder.AddBeer(adding.Beer{Name: name})
		if err != nil {
			return err
		}

		beersChan <- beer
		return nil
	}, nil)

	select {
	case beer := <-beersChan:
		return &pb.BeerAddResponse{
			Status: pb.Status_OK,
			Id:     int32(beer.Id),
			Name:   beer.Name,
		}, nil
	case err := <-errors:
		return &pb.BeerAddResponse{
			Status: pb.Status_BAD,
			ErrMsg: err.Error(),
		}, nil
	}
}
