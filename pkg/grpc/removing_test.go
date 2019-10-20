package grpc

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"github.com/azbshiri/beers-srv/pkg/removing"
)

func (r *dummyRepo) RemoveBeer(b removing.Beer) (*removing.Beer, error) {
	if r.err != nil {
		return nil, r.err
	}

	return &b, nil
}

func Test_Remove(t *testing.T) {
	s := &server{remover: removing.NewService(&dummyRepo{nil})}
	badS := &server{remover: removing.NewService(&dummyRepo{fmt.Errorf("pg cannot connect")})}

	type args struct {
		ctx context.Context
		r   *pb.BeerRemoveRequest
	}
	tests := []struct {
		name    string
		srv     *server
		args    args
		want    *pb.BeerRemoveResponse
		wantErr bool
	}{
		{"adds beer", s, args{
			context.Background(),
			&pb.BeerRemoveRequest{Id: 1}},
			&pb.BeerRemoveResponse{
				Status: pb.Status_OK,
				Id:     int32(1),
			},
			false,
		},
		{"checks against blank name", s, args{
			context.Background(),
			&pb.BeerRemoveRequest{Id: -1}},
			&pb.BeerRemoveResponse{
				Status: pb.Status_BAD,
				ErrMsg: "invalid id"},
			false,
		},
		{"fails due to service", badS, args{
			context.Background(),
			&pb.BeerRemoveRequest{Id: 1}},
			&pb.BeerRemoveResponse{
				Status: pb.Status_BAD,
				ErrMsg: "pg cannot connect"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Remove(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Log(got)
				t.Errorf("server.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
