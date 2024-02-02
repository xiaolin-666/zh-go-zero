package logic

import (
	"context"

	"zh-go-zero/application/follow/rpc/internal/svc"
	"zh-go-zero/application/follow/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowListLogic) FollowList(in *service.FollowListRequest) (*service.FollowListResponse, error) {
	// todo: add your logic here and delete this line

	return &service.FollowListResponse{}, nil
}
