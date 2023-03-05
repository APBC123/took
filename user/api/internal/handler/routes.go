// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"model_user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/user",
				Handler: GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/relation/action",
				Handler: FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follow/list",
				Handler: GetFollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/realtion/follower/list",
				Handler: GetFollowerListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/friend/list",
				Handler: GetFriendListHandler(serverCtx),
			},
		},
	)
}