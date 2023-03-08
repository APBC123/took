package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	var usr model.User
	has, _ := l.svcCtx.Engine.Where("username=? AND password=?", in.Username, in.Password).Get(&usr)

	if !has {
		return &user.LoginResp{
			StatusCode: 1,
			StatusMsg: "用户名或密码错误",
		}, nil
	}
	
	// 更新最近登录时间
	usr.LoginTime = time.Now()
	l.svcCtx.Engine.ID(usr.Id).Cols("login_time").Update(usr)

	jwtToken, _ := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, usr.Id)

	return &user.LoginResp{
		StatusCode: 0,
		StatusMsg: "登录成功",
		UserId: usr.Id,
		Token: jwtToken,
	}, nil
}
