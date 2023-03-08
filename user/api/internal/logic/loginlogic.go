package logic

import (
	"context"
	"errors"
	"strings"

	"took/user/api/internal/svc"
	"took/user/api/internal/types"
	"took/user/api/internal/helper"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
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

	rpcResp, _ := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if rpcResp.StatusCode != 0 {
		return &types.LoginResp{
			StatusCode: rpcResp.StatusCode,
			StatusMsg: rpcResp.StatusMsg,
		}, nil
	}

	jwtToken, _ := helper.GenerateToken(rpcResp.UserId, req.Username, req.Password,
		l.svcCtx.Config.JwtAuth.SecretKey, l.svcCtx.Config.JwtAuth.Duration)

	return &types.LoginResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg: rpcResp.StatusMsg,
		UserId: rpcResp.UserId,
		Token: jwtToken,
	}, nil
}
