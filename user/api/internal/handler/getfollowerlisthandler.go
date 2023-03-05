package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"model_user/api/internal/logic"
	"model_user/api/internal/svc"
	"model_user/api/internal/types"
)

func GetFollowerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFollowerListLogic(r.Context(), svcCtx)
		resp, err := l.GetFollowerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
