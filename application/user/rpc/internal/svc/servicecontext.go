package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zh-go-zero/application/user/rpc/internal/config"
	"zh-go-zero/application/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
