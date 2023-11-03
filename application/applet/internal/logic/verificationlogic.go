package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"zh-go-zero/application/applet/internal/svc"
	"zh-go-zero/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const prefixActivation = "register"

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req *types.VerificationRequrst) (resp *types.VerificationResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

func getVerificationCode(rds *redis.Redis, mobile string) (code string, err error) {
	code, err = rds.Get(prefixActivation + mobile)
	return code, err
}

func deleteVerificationCode(rds *redis.Redis, mobile string) error {
	_, err := rds.Del(prefixActivation + mobile)
	return err
}
