package logic

import (
	"context"

	"took/user/api/internal/helper"
	"took/user/api/internal/svc"
	"took/user/api/internal/types"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowListLogic) GetFollowList(req *types.FollowListReq) (resp *types.FollowListResp, err error) {
	uc, err := helper.AnalyzeToken(req.Token, l.svcCtx.Config.JwtAuth.SecretKey)
	if err != nil {
		return &types.FollowListResp{
			StatusCode: 3,
			StatusMsg: err.Error(),
		}, nil;
	}

	rpcResp, _ := l.svcCtx.UserRpc.GetFollowList(l.ctx, &user.FollowListReq{
		UserId: uc.Id,
		ToUserId: req.UserId,
	})

	return &types.FollowListResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg: rpcResp.StatusMsg,
		UserList: types.NewUserList(rpcResp.UserList),
	}, nil
}
