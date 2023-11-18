package logic

import (
	"context"
	"strings"
	"zh-go-zero/application/applet/internal/code"
	"zh-go-zero/application/applet/internal/svc"
	"zh-go-zero/application/applet/internal/types"
	"zh-go-zero/application/user/rpc/user"
	"zh-go-zero/pkg/encrypt"
	"zh-go-zero/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	req.Name = strings.TrimSpace(req.Name)
	if len(req.Name) == 0 {
		return nil, code.RegisterUsernameEmpty
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, code.RegisterPasswdEmpty
	} else {
		req.Password = encrypt.EncPasswd(req.Password)
	}
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	} else {
		mobile, err := encrypt.EncMobile(req.Mobile)
		if err != nil {
			return nil, err
		}
		req.Mobile = mobile
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}
	err = checkVerificationCode(l.svcCtx.RedisCli, req.Mobile, req.VerificationCode)
	if err != nil {
		return nil, err
	}
	u, err := l.svcCtx.UserRpc.FindByMobile(l.ctx, &user.FindByMobileRequest{Mobile: req.Mobile})
	if err != nil {
		return nil, err
	}
	if u.GetUserId() > 0 {
		return nil, code.MobileHasRegistered
	}
	regRet, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Mobile:   req.Mobile,
		Username: req.Name,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	_ = deleteVerificationCode(l.svcCtx.RedisCli, req.Mobile)

	token, err := jwt.BuildToken(jwt.TokenOption{
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		Fields: map[string]interface{}{
			"userId": regRet.GetUserId(),
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{
		UserId: regRet.GetUserId(),
		ToKen: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil

}
