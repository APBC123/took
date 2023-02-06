package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"model_video/core/internal/logic"
	"model_video/core/internal/svc"
)

func GetVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetVideoListLogic(r.Context(), svcCtx)
		err := l.GetVideoList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
