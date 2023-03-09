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

	resp := user.FollowResp{}
	if in.ActionType == 1 {
		l.svcCtx.Engine.Insert(&model.Follow{
			FanId: in.UserId,
			UserId: in.ToUserId,
		})
		resp.StatusMsg = "关注成功"
	} else {
		l.svcCtx.Engine.Table("follow").Where("user_id = ? AND fan_id = ?", in.ToUserId, in.UserId).Delete()
		resp.StatusMsg = "取消关注成功"
	}

	return &resp, nil
}
