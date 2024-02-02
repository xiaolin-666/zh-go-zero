package logic

import (
	"context"

	"zh-go-zero/application/follow/rpc/internal/svc"
	"zh-go-zero/application/follow/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FansListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansListLogic {
	return &FansListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FansListLogic) FansList(in *service.FansListRequest) (*service.FansListResponse, error) {
	// todo: add your logic here and delete this line

	return &service.FansListResponse{}, nil
}
