package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"zh-go-zero/application/applet/internal/config"
	"zh-go-zero/application/user/rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  user.User
	RedisCli *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		RedisCli: redis.MustNewRedis(c.RedisCli),
	}
}
