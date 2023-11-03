package code

import "zh-go-zero/pkg/xcode"

var (
	RegisterMobileEmpty   = xcode.New(100001, "注册手机号不能为空")
	VerificationCodeEmpty = xcode.New(100002, "验证码不能为空")
	MobileHasRegistered   = xcode.New(100003, "手机号已经注册")
	LoginMobileEmpty      = xcode.New(100003, "手机号不能为空")
	RegisterPasswdEmpty   = xcode.New(100004, "密码不能为空")
	RegisterUsernameEmpty = xcode.New(100005, "用户名不能为空")
)
