package logic

import (
	"context"

	"took/user/api/internal/svc"
	"took/user/api/internal/types"
	"took/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	var user model.User;
	l.svcCtx.Engine.Where("id=?", req.UserId).Get(&user)

	return &types.UserInfoResp{
		StatusCode: 0,
		StatusMsg: "success",
		User: types.User{
			Id: user.Id,
			Username: user.Username,
			FollowCount: user.FollowCount,
			FollowerCount: user.FollowerCount,
			FavoriteCount: user.FavoriteCount,
			IsFollow: false, // TODO...
			Avatar: user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature: user.Signature,
			TotalFavorited: user.TotalFavorited,
			WorkCount: user.WorkCount,
		},
	}, nil
}
