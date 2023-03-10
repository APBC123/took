package logic

import (
	"context"
	"crypto/md5"
	"time"
	"fmt"

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
	var usr model.User
	has, _ := l.svcCtx.Engine.Where("username=? AND password=?", in.Username, hashPwd).Get(&usr)

	if !has {
		return &user.LoginResp{
			StatusCode: 2,
			StatusMsg: "用户名或密码错误",
		}, nil
	}
	
	if usr.Enable == 0 {
		return &user.LoginResp{
			StatusCode: 6,
			StatusMsg: "该帐号封禁中...",
		}, nil
	}

	if usr.Deleted == 1 {
		return &user.LoginResp{
			StatusCode: 6,
			StatusMsg: "该帐号已注销",
		}, nil
	}

	// 更新最近登录时间
	usr.LoginTime = time.Now()
	l.svcCtx.Engine.ID(usr.Id).Cols("login_time").Update(usr)

	return &user.LoginResp{
		StatusCode: 0,
		StatusMsg: "登录成功",
		UserId: usr.Id,
	}, nil
}
