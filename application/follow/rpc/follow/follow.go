// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package follow

import (
	"context"

	"zh-go-zero/application/follow/rpc/service"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FansItem           = service.FansItem
	FansListRequest    = service.FansListRequest
	FansListResponse   = service.FansListResponse
	FollowItem         = service.FollowItem
	FollowListRequest  = service.FollowListRequest
	FollowListResponse = service.FollowListResponse
	FollowRequest      = service.FollowRequest
	FollowResponse     = service.FollowResponse
	UnFollowRequest    = service.UnFollowRequest
	UnFollowResponse   = service.UnFollowResponse

	Follow interface {
		Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
		UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error)
		FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
		FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error)
	}

	defaultFollow struct {
		cli zrpc.Client
	}
)

func NewFollow(cli zrpc.Client) Follow {
	return &defaultFollow{
		cli: cli,
	}
}

func (m *defaultFollow) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	client := service.NewFollowClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

func (m *defaultFollow) UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error) {
	client := service.NewFollowClient(m.cli.Conn())
	return client.UnFollow(ctx, in, opts...)
}

func (m *defaultFollow) FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	client := service.NewFollowClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

func (m *defaultFollow) FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error) {
	client := service.NewFollowClient(m.cli.Conn())
	return client.FansList(ctx, in, opts...)
}
