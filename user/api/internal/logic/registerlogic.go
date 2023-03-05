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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *RegisterLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	has, _ := l.svcCtx.Engine.Exist(&model.User{
		Username: req.Username,
	})
	if has {
		return &types.RegisterResp{
			StatusCode: 1,
			StatusMsg: "用户已存在",
		}, nil
	}

	user := model.User {
		Username: req.Username,
		Password: req.Password,
		Enable: 1,
	}

	l.svcCtx.Engine.Insert(&user)

	jwtToken, _ := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, user.Id)

	return &types.RegisterResp{
		StatusCode: 0,
		StatusMsg: "注册成功",
		UserId: user.Id,
		Token: jwtToken,	
	}, nil
}
