package logic

import (
	"context"
	"errors"
	"strings"

	"took/user/rpc/types/user"
	"took/server/service/core/helper"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
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

	rpcResp, _ := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: req.Password,
		Password: req.Password,
	})
	if rpcResp.StatusCode != 0 {
		return &types.RegisterResp{
			StatusCode: rpcResp.StatusCode,
			StatusMsg: rpcResp.StatusMsg,
		}, nil
	}

	jwtToken, _ := helper.GenerateToken(rpcResp.UserId, req.Username, req.Password, 
	l.svcCtx.Config.JwtAuth.SecretKey, l.svcCtx.Config.JwtAuth.Duration)

	return &types.RegisterResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg: rpcResp.StatusMsg,
		UserId: rpcResp.UserId,
		Token: jwtToken,
	}, nil
}
