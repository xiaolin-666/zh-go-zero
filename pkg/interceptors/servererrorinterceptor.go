package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"zh-go-zero/pkg/xcode"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		return resp, xcode.FromError(err).Err()
	}
}
