// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package server

import (
	"context"

	"zh-go-zero/application/article/rpc/internal/logic"
	"zh-go-zero/application/article/rpc/internal/svc"
	"zh-go-zero/application/article/rpc/service"
)

type ArticleServer struct {
	svcCtx *svc.ServiceContext
	service.UnimplementedArticleServer
}

func NewArticleServer(svcCtx *svc.ServiceContext) *ArticleServer {
	return &ArticleServer{
		svcCtx: svcCtx,
	}
}

func (s *ArticleServer) PublishArticle(ctx context.Context, in *service.PublishArticleRequest) (*service.PublishArticleResponse, error) {
	l := logic.NewPublishArticleLogic(ctx, s.svcCtx)
	return l.PublishArticle(in)
}

func (s *ArticleServer) Articles(ctx context.Context, in *service.ArticlesRequest) (*service.ArticlesResponse, error) {
	l := logic.NewArticlesLogic(ctx, s.svcCtx)
	return l.Articles(in)
}

func (s *ArticleServer) ArticleDelete(ctx context.Context, in *service.ArticleDeleteRequest) (*service.ArticleDeleteResponse, error) {
	l := logic.NewArticleDeleteLogic(ctx, s.svcCtx)
	return l.ArticleDelete(in)
}

func (s *ArticleServer) ArticleDetail(ctx context.Context, in *service.ArticleDetailRequest) (*service.ArticleDetailResponse, error) {
	l := logic.NewArticleDetailLogic(ctx, s.svcCtx)
	return l.ArticleDetail(in)
}
