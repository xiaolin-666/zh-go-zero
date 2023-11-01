package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strings"
	"zh-go-zero/application/applet/internal/code"
	"zh-go-zero/application/applet/internal/svc"
	"zh-go-zero/application/applet/internal/types"
	"zh-go-zero/application/user/rpc/user"
	"zh-go-zero/pkg/encrypt"
	"zh-go-zero/pkg/jwt"
	"zh-go-zero/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}
	err = checkVerificationCode(l.svcCtx.RedisCli, req.Mobile, req.VerificationCode)
	if err != nil {
		return nil, err
	}
	// 手机号密文保存
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		return nil, err
	}
	u, err := l.svcCtx.UserRpc.FindByMobile(l.ctx, &user.FindByMobileRequest{Mobile: mobile})
	if err != nil {
		return nil, err
	}
	if u.GetUserId() == 0 {
		return nil, xcode.AccessDenied
	}
	token, err := jwt.BuildToken(jwt.TokenOption{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": u.UserId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		UserId: u.GetUserId(),
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}

func checkVerificationCode(rds *redis.Redis, mobile string, code string) (err error) {
	cacheCode, err := getVerificationCode(rds, mobile)
	if err != nil {
		return err
	}
	if cacheCode == "" {
		return errors.New("verification code expired")
	}
	if cacheCode != code {
		return errors.New("verification code does not match")
	}
	return nil
}
