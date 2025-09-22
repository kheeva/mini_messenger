package server

import (
	"context"
	"fmt"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "mini_messenger/user/pkg/api/v1"

	"buf.build/go/protovalidate"
)

type server struct {
	// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
	pb.UnimplementedUserServiceServer
	healthpb.UnimplementedHealthServer

	validator protovalidate.Validator
}

func NewServer() (*server, error) {
	srv := &server{}
	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&pb.CreateProfileRequest{},
			&pb.UpdateProfileRequest{},
			&pb.GetProfileByIDRequest{},
			&pb.GetProfileByNicknameRequest{},
			&pb.SearchByNicknameRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator
	return srv, nil
}

func (s *server) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	bio := "bio"
	avatarUrl := "http://example.com"
	return &pb.CreateProfileResponse{UserId: "1", Nickname: "zerocool", Bio: &bio, AvatarUrl: &avatarUrl}, nil
}

func (s *server) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	bio := "bio"
	avatarUrl := "http://example.com"
	return &pb.UpdateProfileResponse{UserId: "1", Nickname: "zerocool", Bio: &bio, AvatarUrl: &avatarUrl}, nil
}

func (s *server) GetProfileByID(ctx context.Context, req *pb.GetProfileByIDRequest) (*pb.GetProfileByIDResponse, error) {
	bio := "bio"
	avatarUrl := "http://example.com"
	return &pb.GetProfileByIDResponse{UserId: "1", Nickname: "zerocool", Bio: &bio, AvatarUrl: &avatarUrl}, nil
}

func (s *server) GetProfileByNickname(ctx context.Context, req *pb.GetProfileByNicknameRequest) (*pb.GetProfileByNicknameResponse, error) {
	bio := "bio"
	avatarUrl := "http://example.com"
	return &pb.GetProfileByNicknameResponse{UserId: "1", Nickname: "zerocool", Bio: &bio, AvatarUrl: &avatarUrl}, nil
}

func (s *server) SearchByNickname(ctx context.Context, req *pb.SearchByNicknameRequest) (*pb.SearchByNicknameResponse, error) {
	bio := "bio"
	avatarUrl := "http://example.com"
	userProfile := &pb.UserProfile{UserId: "1", Nickname: "zerocool", Bio: &bio, AvatarUrl: &avatarUrl}
	return &pb.SearchByNicknameResponse{Results: []*pb.UserProfile{userProfile}}, nil
}

func (s *server) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

// Watch implements the HealthCheckServiceServer Watch method.
func (s *server) Watch(in *healthpb.HealthCheckRequest, stream healthpb.Health_WatchServer) error {
	return stream.Send(&healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING})
}
