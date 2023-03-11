package logic

import (
	"context"

	"took/user/api/internal/helper"
	"took/user/api/internal/svc"
	"took/user/api/internal/types"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	uc, err := helper.AnalyzeToken(req.Token, l.svcCtx.Config.JwtAuth.SecretKey)
	if err != nil {
		return &types.FriendListResp{
			StatusCode: 3,
			StatusMsg: err.Error(),
		}, nil;
	}

	rpcResp, _ := l.svcCtx.UserRpc.GetFriendList(l.ctx, &user.FriendListReq{
		UserId: uc.Id,
		ToUserId: req.UserId,
	})

	return &types.FriendListResp{
		StatusCode: 0,
		StatusMsg: "success",
		UserList: types.NewFriendUserList(rpcResp.UserList),
	}, nil
}
