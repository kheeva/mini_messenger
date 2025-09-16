package server

import (
	"context"
	"fmt"

	pb "mini_messenger/auth/pkg/api"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"buf.build/go/protovalidate"
)

type server struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	pb.UnimplementedAuthServiceServer
	healthpb.UnimplementedHealthServer

	validator protovalidate.Validator
}

func NewServer() (*server, error) {
	srv := &server{}
	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&pb.RegisterRequest{},
			&pb.LoginRequest{},
			&pb.RefreshRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator
	return srv, nil
}

func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{UserId: "1"}, nil
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{UserId: "1", AccessToken: "access_token", RefreshToken: "refresh_token"}, nil
}

func (s *server) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return &pb.RefreshResponse{UserId: "1", AccessToken: "access_token", RefreshToken: "refresh_token"}, nil
}

func (s *server) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

// Watch implements the HealthCheckServiceServer Watch method.
func (s *server) Watch(in *healthpb.HealthCheckRequest, stream healthpb.Health_WatchServer) error {
	return stream.Send(&healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING})
}
