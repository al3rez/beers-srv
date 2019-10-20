package grpc

import (
	"context"

	v1 "github.com/azbshiri/beers-proto/pkg/pb/grpc_health_v1"
)

func (srv *server) Check(ctx context.Context, r *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {
	return &v1.HealthCheckResponse{
		Status: v1.HealthCheckResponse_SERVING,
	}, nil
}

func (srv *server) Watch(*v1.HealthCheckRequest, v1.Health_WatchServer) error {
	return nil
}

