package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zh-go-zero/application/applet/internal/logic"
	"zh-go-zero/application/applet/internal/svc"
	"zh-go-zero/application/applet/internal/types"
)

func VerificationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerificationRequrst
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVerificationLogic(r.Context(), svcCtx)
		resp, err := l.Verification(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
