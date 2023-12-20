package logic

import (
	"context"

	"zh-go-zero/application/article/rpc/internal/svc"
	"zh-go-zero/application/article/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishArticleLogic {
	return &PublishArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishArticleLogic) PublishArticle(in *service.PublishArticleRequest) (*service.PublishArticleResponse, error) {
	// todo: add your logic here and delete this line

	return &service.PublishArticleResponse{}, nil
}
