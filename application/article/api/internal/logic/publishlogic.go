package logic

import (
	"context"
	"encoding/json"
	"zh-go-zero/application/article/api/code"
	"zh-go-zero/application/article/rpc/service"
	"zh-go-zero/pkg/xcode"

	"zh-go-zero/application/article/api/internal/svc"
	"zh-go-zero/application/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const minContentLen = 80

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishArticleRequest) (resp *types.PublishArticleResponse, err error) {
	if len(req.Cover) == 0 {
		return nil, code.ArticleCoverEmpty
	}
	if len(req.Title) == 0 {
		return nil, code.ArticleTitleEmpty
	}
	if len(req.Content) <= minContentLen {
		return nil, code.ArticleContentTooFewWords
	}

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		logx.Errorf("parse l.ctx.value userId err: %v", err)
		return nil, xcode.NoLogin
	}
	article, err := l.svcCtx.ArticleRpc.PublishArticle(l.ctx, &service.PublishArticleRequest{
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		Cover:       req.Cover,
		UserId:      userId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.ArticleRpc.PublishArticle err: %v, userId: %v, req: %v", err, userId, req)
		return nil, err
	}
	return &types.PublishArticleResponse{ArticleId: article.ArticleId}, nil
}
