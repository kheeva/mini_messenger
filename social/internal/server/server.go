package server

import (
	"context"
	"fmt"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "mini_messenger/social/pkg/api/v1"

	"buf.build/go/protovalidate"
)

const (
	PendingStatus  = "PENDING"
	AcceptedStatus = "ACCEPTED"
	DeclinedStatus = "DECLINED"
)

type server struct {
	// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
	pb.UnsafeSocialServiceServer
	healthpb.UnimplementedHealthServer

	validator protovalidate.Validator
}

func NewServer() (*server, error) {
	srv := &server{}
	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&pb.SendFriendRequest{},
			&pb.ListRequestsRequest{},
			&pb.AcceptFriendRequest{},
			&pb.DeclineFriendRequest{},
			&pb.DeclineFriendRequest{},
			&pb.RemoveFriendRequest{},
			&pb.ListFriendsRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator
	return srv, nil
}

func (s *server) SendFriend(ctx context.Context, req *pb.SendFriendRequest) (*pb.SendFriendResponse, error) {
	return &pb.SendFriendResponse{RequestId: "1", Status: PendingStatus}, nil
}

func (s *server) ListRequests(ctx context.Context, req *pb.ListRequestsRequest) (*pb.ListRequestsResponse, error) {
	fr := &pb.FriendRequest{RequestId: "1", Status: PendingStatus}
	return &pb.ListRequestsResponse{Requests: []*pb.FriendRequest{fr}}, nil
}

func (s *server) AcceptFriend(ctx context.Context, req *pb.AcceptFriendRequest) (*pb.AcceptFriendResponse, error) {
	return &pb.AcceptFriendResponse{RequestId: "1", Status: AcceptedStatus}, nil
}

func (s *server) DeclineFriend(ctx context.Context, req *pb.DeclineFriendRequest) (*pb.DeclineFriendResponse, error) {
	return &pb.DeclineFriendResponse{RequestId: "1", Status: DeclinedStatus}, nil
}

func (s *server) RemoveFriend(ctx context.Context, req *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	return &pb.RemoveFriendResponse{}, nil
}

func (s *server) ListFriends(ctx context.Context, req *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	return &pb.ListFriendsResponse{FriendUserIds: []string{"1"}}, nil
}

func (s *server) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

// Watch implements the HealthCheckServiceServer Watch method.
func (s *server) Watch(in *healthpb.HealthCheckRequest, stream healthpb.Health_WatchServer) error {
	return stream.Send(&healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING})
}
