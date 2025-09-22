package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	server "mini_messenger/social/internal/server"
	pb "mini_messenger/social/pkg/api/v1"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("failed to create social server: %v", err)
	}

	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(server.UnaryValidationInterceptor(*srv)))
	healthpb.RegisterHealthServer(gRPCServer, srv)
	pb.RegisterSocialServiceServer(gRPCServer, srv)

	reflection.Register(gRPCServer)

	log.Printf("Social service listening at %v", lis.Addr())
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
