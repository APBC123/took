package logic

import (
	"context"

	"took/user/api/internal/helper"
	"took/user/api/internal/svc"
	"took/user/api/internal/types"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowerListLogic) GetFollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
	_, err = helper.AnalyzeToken(req.Token, l.svcCtx.Config.JwtAuth.SecretKey)
	if err != nil {
		return &types.FollowerListResp{
			StatusCode: 3,
			StatusMsg: err.Error(),
		}, nil;
	}

	rpcResp, _ := l.svcCtx.UserRpc.GetFollowerList(l.ctx, &user.FollowerListReq{
		UserId: req.UserId,
	})

	return &types.FollowerListResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg: rpcResp.StatusMsg,
		UserList: types.NewUserList(rpcResp.UserList),
	}, nil
}
