package logic

import (
	"context"

	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *user.FollowListReq) (*user.FollowListResp, error) {
	var followerList []*user.User
	l.svcCtx.Engine.Table("user").Join("LEFT", "follow", "user.id = follow.user_id").Select(
		"user.*").Where("follow.fan_id = ?", in.UserId).Find(&followerList)

	return &user.FollowListResp{
		StatusCode: 0,
		StatusMsg: "success",
		UserList: followerList,
	}, nil
}
