package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zh-go-zero/application/article/api/internal/logic"
	"zh-go-zero/application/article/api/internal/svc"
)

func UploadCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadCoverLogic(r.Context(), svcCtx)
		resp, err := l.UploadCover(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
