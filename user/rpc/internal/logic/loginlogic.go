package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	hashPwd := fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))
	usr := model.User{
		Username: in.Username,
	}

	has, _ := l.svcCtx.UserModel.GetByName(l.ctx, &usr)
	if !has {
		return &user.LoginResp{
			StatusCode: 2,
			StatusMsg:  "用户名或密码错误",
		}, nil
	}

	if usr.Enable == false {
		return &user.LoginResp{
			StatusCode: 6,
			StatusMsg:  "该帐号封禁中...",
		}, nil
	}

	if usr.Deleted == true {
		return &user.LoginResp{
			StatusCode: 6,
			StatusMsg:  "该帐号已注销",
		}, nil
	}

	if usr.Password != hashPwd {
		return &user.LoginResp{
			StatusCode: 2,
			StatusMsg:  "用户名或密码错误", // 不要明确告知是用户名错误还是密码错误
		}, nil
	}

	// 更新最近登录时间
	usr.LoginTime = time.Now()
	l.svcCtx.UserModel.Update(l.ctx, &usr)

	return &user.LoginResp{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		UserId:     usr.Id,
	}, nil
}
