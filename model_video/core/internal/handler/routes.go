// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"model_video/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/douyin/feed",
				Handler: GetVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/publish/action",
				Handler: PublishVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/publish/list",
				Handler: PublishListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/comment/list",
				Handler: GetCommentListHandler(serverCtx),
			},
		},
	)
}
