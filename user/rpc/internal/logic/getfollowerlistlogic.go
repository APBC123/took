package logic

import (
	"context"
	// "fmt"

	//"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *user.FollowerListReq) (*user.FollowerListResp, error) {
	var followerList []*user.User
	l.svcCtx.Engine.Table("user").Join("LEFT", "follow", "user.id = follow.fan_id").Select(
		"user.*").Where("follow.user_id = ?", in.UserId).Find(&followerList)

	// println("<============>")
	// fmt.Printf("%#v", followerList)

	return &user.FollowerListResp{
		StatusCode: 0,
		StatusMsg: "success",
		UserList: followerList,
	}, nil
}
