package grpc

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/azbshiri/beers-proto/pkg/pb"
	"github.com/azbshiri/beers-srv/pkg/adding"
)

type dummyRepo struct {
	err error
}

func (r *dummyRepo) AddBeer(b adding.Beer) (*adding.Beer, error) {
	if r.err != nil {
		return nil, r.err
	}

	return &b, nil
}

func Test_Add(t *testing.T) {
	s := &server{adder: adding.NewService(&dummyRepo{nil})}
	badS := &server{adder: adding.NewService(&dummyRepo{fmt.Errorf("pg cannot connect")})}

	type args struct {
		ctx context.Context
		r   *pb.BeerAddRequest
	}
	tests := []struct {
		name    string
		srv     *server
		args    args
		want    *pb.BeerAddResponse
		wantErr bool
	}{
		{"adds beer", s, args{
			context.Background(),
			&pb.BeerAddRequest{Name: "ali"}},
			&pb.BeerAddResponse{
				Status: pb.Status_OK,
				Id:     int32(0),
				Name:   "ali"},
			false,
		},
		{"checks against blank name", s, args{
			context.Background(),
			&pb.BeerAddRequest{Name: ""}},
			&pb.BeerAddResponse{
				Status: pb.Status_BAD,
				ErrMsg: "name cannot be blank"},
			false,
		},
		{"fails due to service", badS, args{
			context.Background(),
			&pb.BeerAddRequest{Name: "ali"}},
			&pb.BeerAddResponse{
				Status: pb.Status_BAD,
				ErrMsg: "pg cannot connect"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Add(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Log(got)
				t.Errorf("server.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
