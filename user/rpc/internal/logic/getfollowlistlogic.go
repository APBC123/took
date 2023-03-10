package logic

import (
	"context"

	"took/user/model"
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
	var followList []*user.User
	l.svcCtx.Engine.Table("user").Join("LEFT", "follow", "user.id = follow.user_id").Select(
		"user.*").Where("follow.fan_id = ?", in.ToUserId).Find(&followList)

	for i := range followList {
		isFollow, _ := l.svcCtx.Engine.Exist(&model.Follow{
			UserId: followList[i].Id,
			FanId: in.UserId,
		})
		followList[i].IsFollow = isFollow
	}

	return &user.FollowListResp{
		StatusCode: 0,
		StatusMsg: "success",
		UserList: followList,
	}, nil
}
