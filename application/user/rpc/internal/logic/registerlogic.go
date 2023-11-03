package logic

import (
	"context"
	"time"
	"zh-go-zero/application/user/rpc/internal/model"

	"zh-go-zero/application/user/rpc/internal/svc"
	"zh-go-zero/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.GetUsername(),
		Avatar:     in.GetAvatar(),
		Mobile:     in.GetMobile(),
		Password:   in.GetPassword(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &service.RegisterResponse{
		UserId: id,
	}, nil
}
