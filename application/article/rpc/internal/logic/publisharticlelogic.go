package logic

import (
	"context"
	"time"
	"zh-go-zero/application/article/rpc/internal/model"
	"zh-go-zero/application/user/rpc/types"

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
	result, err := l.svcCtx.ArticleModel.Insert(l.ctx, &model.Article{
		Title:       in.Title,
		Description: in.Description,
		Content:     in.Content,
		AuthorId:    in.UserId,
		Cover:       in.Cover,
		PublishTime: time.Now(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		Status:      types.ArticleStatusVisible,
	})
	if err != nil {
		l.Logger.Errorf("publish insert req: %v error: %v", in, err)
		return nil, err
	}
	artId, err := result.LastInsertId()
	if err != nil {
		l.Logger.Errorf("LastInsertId error: %v", err)
		return nil, err
	}
	return &service.PublishArticleResponse{ArticleId: artId}, nil
}
