package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zh-go-zero/application/article/rpc/internal/config"
	"zh-go-zero/application/article/rpc/internal/model"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.Datasource), c.CacheRedis),
	}
}
