package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "mini_messenger/auth/pkg/api/v1"
)

func run() error {
	grpcEndpoint := flag.String("auth-grpc-endpoint", "auth:8000", "auth gRPC endpoint")
	httpAddr := flag.String("http-addr", ":8080", "HTTP listen address")
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *grpcEndpoint, opts)
	if err != nil {
		return err
	}

	log.Printf("Starting HTTP gateway on %s, proxying to gRPC %s", *httpAddr, *grpcEndpoint)
	return http.ListenAndServe(*httpAddr, mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run gateway: %v", err)
	}
}
