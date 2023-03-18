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

	isFollow, _ := l.svcCtx.FollowModel.Exist(l.ctx, &model.Follow{
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

	fromUser := model.User{Id: in.UserId}
	toUser := model.User{Id: in.ToUserId}
	l.svcCtx.UserModel.GetById(l.ctx, &fromUser)
	l.svcCtx.UserModel.GetById(l.ctx, &toUser)

	resp := user.FollowResp{}
	if in.ActionType == 1 {
		l.svcCtx.FollowModel.Insert(&model.Follow{
			FanId: in.UserId,
			UserId: in.ToUserId,
		})
		fromUser.FollowCount++
		toUser.FollowerCount++
		resp.StatusMsg = "关注成功"
	} else {
		l.svcCtx.FollowModel.Delete(&model.Follow{
			FanId: in.UserId,
			UserId: in.ToUserId,
		})
		fromUser.FollowCount--
		toUser.FollowerCount--
		resp.StatusMsg = "取消关注成功"
	}
	l.svcCtx.UserModel.Update(l.ctx, &fromUser)
	l.svcCtx.UserModel.Update(l.ctx, &toUser)
	return &resp, nil
}
