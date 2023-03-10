package logic

import (
	"context"
	"took/server/service/core/helper"
	"took/video/video/rpc/types/video"

	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListRequest) (resp *types.FavoriteListResponse, err error) {
	favoriteList, err := l.svcCtx.VideoRpc.FavoriteList(l.ctx, &video.FavoriteListRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.FavoriteListResponse)
	resp.VideoList = helper.NewVideoList(favoriteList.VideoList)
	resp.StatusCode = 0
	resp.StatusMsg = ""
	return resp, nil
}
