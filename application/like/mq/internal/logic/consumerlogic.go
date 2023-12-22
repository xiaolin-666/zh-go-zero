package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"zh-go-zero/application/like/mq/internal/svc"
)

type ConsumerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConsumerLogic {
	//queue := kq.MustNewQueue(svcCtx.Config.KqConsumerConf, ConsumerLogic)
	return &ConsumerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumerLogic) Consume(key, value string) error {
	fmt.Printf("get key: %v, val: %v", key, value)
	return nil
}

func Consumer(ctx context.Context, svcCtx *svc.ServiceContext) service.Service {
	return kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewConsumerLogic(ctx, svcCtx))
}
