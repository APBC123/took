package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"model_user/api/internal/svc"
	"model_user/api/internal/types"
	"model_user/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
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
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	var user model.User
	has, _ := l.svcCtx.Engine.Where("username=? AND password=?", req.Username, req.Password).Get(&user)

	if !has {
		return &types.LoginResp{
			StatusCode: 1,
			StatusMsg: "用户名或密码错误",
		}, nil
	}

	jwtToken, _ := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, user.Id)

	return &types.LoginResp{
		StatusCode: 0,
		StatusMsg: "登录成功",
		UserId: user.Id,
		Token: jwtToken,
	}, nil
}
