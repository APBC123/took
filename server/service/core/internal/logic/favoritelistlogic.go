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
	resp.VideoList = make([]types.Video, len(favoriteList.VideoList))
	for i := range resp.VideoList {
		resp.VideoList[i].IsFavorite = favoriteList.VideoList[i].IsFavorite
		resp.VideoList[i].Id = favoriteList.VideoList[i].Id
		resp.VideoList[i].Title = favoriteList.VideoList[i].Title
		resp.VideoList[i].PlayUrl = favoriteList.VideoList[i].PlayUrl
		resp.VideoList[i].CoverUrl = favoriteList.VideoList[i].CoverUrl
		resp.VideoList[i].CommentCount = favoriteList.VideoList[i].CommentCount
		resp.VideoList[i].FavoriteCount = favoriteList.VideoList[i].FavoriteCount
		resp.VideoList[i].Author = helper.NewUser(favoriteList.VideoList[i].Author)
	}
	resp.StatusCode = 0
	resp.StatusMsg = ""

	return resp, nil
}
