package logic

import (
	"context"
	"math/rand"

	"took/user/model"
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
	friendList := user.NewFriendList(l.svcCtx.FollowModel.GetFriendById(l.ctx, in.ToUserId))
	
	for i := range friendList {
		isFollow, _ := l.svcCtx.FollowModel.Exist(l.ctx, &model.Follow{
			UserId: friendList[i].Id,
			FanId: in.UserId,
		})
		friendList[i].IsFollow = isFollow
	}

	// 与好友的最新信息，待完善...
	text := []string{
		"最近在忙什么？",
		"明天还有个PPT展示",
		"hello",
		"今晚去哪吃饭？",
		"来打游戏不？",
		"来了来了",
		"后天吧，明天还有事",
		"搞快点，就差你一个了",
		"睡觉了，白白",
	}
	for i := range friendList {
		friendList[i].Message = text[rand.Intn(len(text))]
		friendList[i].MsgType = rand.Int63n(2)
	}

	return &user.FriendListResp{
		StatusCode: 0,
		StatusMsg: "success",
		UserList: friendList,
	}, nil
}
