package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"took/server/service/core/internal/logic"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
)

func SendChatMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendChatMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendChatMessageLogic(r.Context(), svcCtx)
		resp, err := l.SendChatMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
