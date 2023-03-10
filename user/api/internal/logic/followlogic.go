package logic

import (
	"context"

	"took/user/api/internal/helper"
	"took/user/api/internal/svc"
	"took/user/api/internal/types"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	uc, err := helper.AnalyzeToken(req.Token, l.svcCtx.Config.JwtAuth.SecretKey)
	if err != nil {
		return &types.FollowResp{
			StatusCode: 3,
			StatusMsg: err.Error(),
		}, nil;
	}

	rpcResp, _ := l.svcCtx.UserRpc.Follow(l.ctx, &user.FollowReq{
		UserId: uc.Id,
		ToUserId: req.ToUserId,
		ActionType: req.ActionType,
	})

	return &types.FollowResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg: rpcResp.StatusMsg,
	}, nil
}
