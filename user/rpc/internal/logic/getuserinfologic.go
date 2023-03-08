package logic

import (
	"context"

	"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	var usr model.User;
	l.svcCtx.Engine.Where("id=?", in.UserId).Get(&usr)

	return &user.UserInfoResp{
		StatusCode: 0,
		StatusMsg: "success",
		User: &user.User{
			Id: usr.Id,
			Username: usr.Username,
			FollowCount: usr.FollowCount,
			FollowerCount: usr.FollowerCount,
			FavoriteCount: usr.FavoriteCount,
			IsFollow: false, // TODO...
			Avatar: usr.Avatar,
			BackgroundImage: usr.BackgroundImage,
			Signature: usr.Signature,
			TotalFavorited: usr.TotalFavorited,
			WorkCount: usr.WorkCount,
		},
	}, nil
}
