package logic

import (
	"context"

	"took/server/service/core/helper"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
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
	if rpcResp.StatusCode != 0 {
		return &types.FollowResp{
			StatusCode: rpcResp.StatusCode,
			StatusMsg: rpcResp.StatusMsg,
		}, nil
	}

	return &types.FollowResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg: rpcResp.StatusMsg,
	}, nil
}
