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
	usr := model.User{
		Id: in.UserId,
	};
	l.svcCtx.UserModel.Get(&usr)

	// 该接口用于用户本人获取人个信息，用户无法关注自己
	respUser := user.NewUser(&usr)
	respUser.IsFollow = false

	return &user.UserInfoResp{
		StatusCode: 0,
		StatusMsg: "success",
		User: respUser,
	}, nil
}
