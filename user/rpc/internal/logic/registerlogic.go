package logic

import (
	"context"
	"crypto/md5"
	"fmt"

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
			StatusMsg: "用户名已存在",
		}, nil
	}

	usr := model.User {
		Username: in.Username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(in.Password))), // 哈希加密
		Avatar: "https://pannta-picture.oss-cn-shenzhen.aliyuncs.com/20230308051013.jpg",
		BackgroundImage: "https://vip1.loli.io/2022/05/11/Nd4lgGtMrFpa8PW.jpg",
		Signature: "爱好是吃蛋挞",
	}

	l.svcCtx.Engine.Cols("username", "password", "avatar", "background_image",
	 "signature").Insert(&usr)

	return &user.RegisterResp{
		StatusCode: 0,
		StatusMsg: "注册成功",
		UserId: usr.Id,
	}, nil
}
