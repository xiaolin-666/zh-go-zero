package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"zh-go-zero/application/follow/rpc/internal/config"
	"zh-go-zero/application/follow/rpc/internal/model"
	"zh-go-zero/pkg/orm"
)

type ServiceContext struct {
	Config           config.Config
	BizRedis         *redis.Redis
	DB               *gorm.DB
	FollowModel      *model.FollowModel
	FollowCountModel *model.FollowCountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustMysqlDB(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxLifeTime:  c.DB.ConnMaxLifetime,
	})
	followModel := model.NewFollowModel(db)
	followCountModel := model.NewFollowCountModel(db)

	bizRedis := redis.MustNewRedis(c.BizRedis)

	return &ServiceContext{
		Config:           c,
		BizRedis:         bizRedis,
		DB:               db,
		FollowModel:      followModel,
		FollowCountModel: followCountModel,
	}
}
