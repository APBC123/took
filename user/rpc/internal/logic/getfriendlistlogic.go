package logic

import (
	"context"

	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *user.FriendListReq) (*user.FriendListResp, error) {
	var friendList []*user.User
	l.svcCtx.Engine.Table("user").Alias("u").Join(
		"INNER", []string{"follow", "f1"}, "u.id = f1.user_id").Join(
		"INNER", []string{"follow", "f2"}, "f1.user_id = f2.fan_id").Select("u.*").Where(
		"f1.fan_id = ? AND f1.fan_id = f2.user_id", in.UserId).Find(&friendList)

	for i := range friendList {
		friendList[i].IsFollow = true
	}

	return &user.FriendListResp{
		StatusCode: 0,
		StatusMsg: "success",
		UserList: friendList,
	}, nil
}
