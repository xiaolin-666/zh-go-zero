package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"zh-go-zero/application/applet/internal/config"
	"zh-go-zero/application/user/rpc/user"
	"zh-go-zero/pkg/interceptors"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  user.User
	RedisCli *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义client拦截器
	userRpc := zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:   c,
		UserRpc:  user.NewUser(userRpc),
		RedisCli: redis.MustNewRedis(c.RedisCli),
	}
}
