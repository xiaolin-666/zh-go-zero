package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
	"zh-go-zero/application/article/rpc/internal/config"
	"zh-go-zero/application/article/rpc/internal/model"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
	BizRedis     *redis.Redis
	SingleFlight singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.BizRedis)

	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.Datasource), c.CacheRedis),
		BizRedis:     rds,
	}
}
