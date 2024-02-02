package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource      string
		MaxIdleConns    int `json:"default=10"`
		MaxOpenConns    int `json:"default=100"`
		ConnMaxLifetime int `json:"default=3600"`
	}
	BizRedis redis.RedisConf
}
