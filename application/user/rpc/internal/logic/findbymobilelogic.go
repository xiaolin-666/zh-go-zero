package logic

import (
	"context"
	"zh-go-zero/application/user/rpc/internal/svc"
	"zh-go-zero/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *service.FindByMobileRequest) (*service.FindByMobileResponse, error) {
	u, err := l.svcCtx.UserModel.FindByMobile(l.ctx, in.Mobile)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return &service.FindByMobileResponse{}, err
	}
	return &service.FindByMobileResponse{
		UserId:   u.Id,
		Mobile:   u.Mobile,
		Username: u.Username,
		Avatar:   u.Avatar,
	}, nil
}
