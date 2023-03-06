package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"took/server/service/core/internal/logic"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
)

func GetVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetVideoLogic(r.Context(), svcCtx)
		resp, err := l.GetVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
