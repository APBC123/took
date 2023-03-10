package logic

import (
	"context"

	"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *user.FollowReq) (*user.FollowResp, error) {
	if in.UserId == in.ToUserId {
		return &user.FollowResp{
			StatusCode: 7,
			StatusMsg: "无法关注自己",
		}, nil
	}

	isFollow, _ := l.svcCtx.Engine.Exist(&model.Follow{
		FanId: in.UserId,
		UserId: in.ToUserId,
	})
	if isFollow && in.ActionType == 1 {
		return &user.FollowResp{
			StatusCode: 4,
			StatusMsg: "已关注该用户",
		}, nil
	} else if !isFollow && in.ActionType == 2 {
		return &user.FollowResp{
			StatusCode: 5,
			StatusMsg: "未关注该用户",
		}, nil
	}

	var fromUser, toUser model.User
	l.svcCtx.Engine.ID(in.UserId).Cols("follow_count", "follower_count").Get(&fromUser)
	l.svcCtx.Engine.ID(in.ToUserId).Cols("follow_count", "follower_count").Get(&toUser)

	resp := user.FollowResp{}
	if in.ActionType == 1 {
		l.svcCtx.Engine.Insert(&model.Follow{
			FanId: in.UserId,
			UserId: in.ToUserId,
		})
		fromUser.FollowCount++
		toUser.FollowerCount++
		resp.StatusMsg = "关注成功"
	} else {
		l.svcCtx.Engine.Table("follow").Where("user_id = ? AND fan_id = ?", in.ToUserId, in.UserId).Delete()
		fromUser.FollowCount--
		toUser.FollowerCount--
		resp.StatusMsg = "取消关注成功"
	}
	l.svcCtx.Engine.ID(in.UserId).Cols("follow_count", "follower_count").Update(fromUser)
	l.svcCtx.Engine.ID(in.ToUserId).Cols("follow_count", "follower_count").Update(toUser)

	return &resp, nil
}
