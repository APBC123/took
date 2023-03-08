package logic

import (
	"context"

	"took/user/api/internal/helper"
	"took/user/api/internal/svc"
	"took/user/api/internal/types"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	_, err = helper.AnalyzeToken(req.Token, l.svcCtx.Config.JwtAuth.SecretKey)
	if err != nil {
		return &types.UserInfoResp{
			StatusCode: 3,
			StatusMsg: err.Error(),
		}, nil;
	}

	rpcResp, _ := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserInfoReq{
		UserId: req.UserId,
	})
	
	return &types.UserInfoResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg: rpcResp.StatusMsg,
		User: types.User{
			Id: rpcResp.User.Id,
			Username: rpcResp.User.Username,
			FollowCount: rpcResp.User.FollowCount,
			FollowerCount: rpcResp.User.FollowerCount,
			FavoriteCount: rpcResp.User.FavoriteCount,
			IsFollow: rpcResp.User.IsFollow,
			Avatar: rpcResp.User.Avatar,
			BackgroundImage: rpcResp.User.BackgroundImage,
			Signature: rpcResp.User.Signature,
			TotalFavorited: rpcResp.User.TotalFavorited,
			WorkCount: rpcResp.User.WorkCount,
		},
	}, nil
}
