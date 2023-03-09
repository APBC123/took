package logic

import (
	"context"

	"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	has, _ := l.svcCtx.Engine.Exist(&model.User{
		Username: in.Username,
	})
	if has {
		return &user.RegisterResp{
			StatusCode: 1,
			StatusMsg: "该用户名已存在",
		}, nil
	}
	
	usr := model.User {
		Username: in.Username,
		Password: in.Password,
		Enable: 1,
		Avatar: "https://pannta-picture.oss-cn-shenzhen.aliyuncs.com/20230308051013.jpg",
		BackgroundImage: "https://vip1.loli.io/2022/05/11/Nd4lgGtMrFpa8PW.jpg",
		Signature: "爱好是吃蛋挞",
	}

	l.svcCtx.Engine.Insert(&usr)

	return &user.RegisterResp{
		StatusCode: 0,
		StatusMsg: "注册成功",
		UserId: usr.Id,
	}, nil
}
