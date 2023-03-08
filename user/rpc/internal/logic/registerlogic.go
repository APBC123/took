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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
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
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	has, _ := l.svcCtx.Engine.Exist(&model.User{
		Username: in.Username,
	})
	if has {
		return &user.RegisterResp{
			StatusCode: 1,
			StatusMsg: "用户已存在",
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

	jwtToken, _ := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, usr.Id)

	return &user.RegisterResp{
		StatusCode: 0,
		StatusMsg: "注册成功",
		UserId: usr.Id,
		Token: jwtToken,	
	}, nil
}
