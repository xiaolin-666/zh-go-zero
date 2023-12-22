package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"zh-go-zero/application/like/mq/internal/config"
	"zh-go-zero/application/like/mq/internal/logic"
	"zh-go-zero/application/like/mq/internal/svc"
)

var configFile = flag.String("f", "etc/like-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, c)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	serviceGroup.Add(logic.Consumer(ctx, svcCtx))

	serviceGroup.Start()

}
