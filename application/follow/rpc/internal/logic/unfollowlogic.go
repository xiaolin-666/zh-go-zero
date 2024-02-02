package logic

import (
	"context"

	"zh-go-zero/application/follow/rpc/internal/svc"
	"zh-go-zero/application/follow/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnFollowLogic) UnFollow(in *service.UnFollowRequest) (*service.UnFollowResponse, error) {
	// todo: add your logic here and delete this line

	return &service.UnFollowResponse{}, nil
}
