package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func UnaryValidationInterceptor(srv server) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if m, ok := req.(proto.Message); ok {
			log.Printf("Received request: %s, payload: %+v", info.FullMethod, req)
			if err := srv.validator.Validate(m); err != nil {
				return nil, rpcValidationError(err)
			}
		}
		return handler(ctx, req)
	}
}
