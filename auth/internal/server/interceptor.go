package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func UnaryValidationInterceptor(srv server) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if m, ok := req.(proto.Message); ok {
			if err := srv.validator.Validate(m); err != nil {
				return nil, rpcValidationError(err)
			}
		}
		return handler(ctx, req)
	}
}
