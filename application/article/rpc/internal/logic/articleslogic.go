package logic

import (
	"context"

	"zh-go-zero/application/article/rpc/internal/svc"
	"zh-go-zero/application/article/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesLogic {
	return &ArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticlesLogic) Articles(in *service.ArticlesRequest) (*service.ArticlesResponse, error) {
	// todo: add your logic here and delete this line

	return &service.ArticlesResponse{}, nil
}
