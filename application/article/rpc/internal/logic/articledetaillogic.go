package logic

import (
	"context"

	"zh-go-zero/application/article/rpc/internal/svc"
	"zh-go-zero/application/article/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDetailLogic {
	return &ArticleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleDetailLogic) ArticleDetail(in *service.ArticleDetailRequest) (*service.ArticleDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &service.ArticleDetailResponse{}, nil
}
