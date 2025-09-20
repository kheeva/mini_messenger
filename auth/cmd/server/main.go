package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	server "mini_messenger/auth/internal/server"
	pb "mini_messenger/auth/pkg/api/v1"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(server.UnaryValidationInterceptor(*srv)))
	healthpb.RegisterHealthServer(gRPCServer, srv)
	pb.RegisterAuthServiceServer(gRPCServer, srv)

	reflection.Register(gRPCServer)

	log.Printf("Auth service listening at %v", lis.Addr())
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// grpcurl -plaintext -d '{"email": "test@example.com", "password": "passWord1234"}' localhost:8000 github.com.kheeva.mini_messenger.auth.AuthService/Register
	// grpcurl -plaintext -d '{"email": "test@example.com", "password": "password"}' localhost:8000 github.com.kheeva.mini_messenger.auth.AuthService/Login
	// grpcurl -plaintext -d '{"refresh_token": "refresh_token"}' localhost:8000 github.com.kheeva.mini_messenger.auth.AuthService/Refresh
}
