package xcode

import (
	"net/http"
	"zh-go-zero/pkg/xcode/types"
)

func ErrHandle(err error) (int, any) {
	// 将err断言为xcode
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}
