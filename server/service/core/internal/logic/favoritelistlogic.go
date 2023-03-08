package logic

import (
	"context"
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
		resp.VideoList[i].Author.Id = favoriteList.VideoList[i].Author.Id
		resp.VideoList[i].Author.Username = favoriteList.VideoList[i].Author.Username
		resp.VideoList[i].Author.Signature = favoriteList.VideoList[i].Author.Signature
		resp.VideoList[i].Author.Avatar = favoriteList.VideoList[i].Author.Avatar
		resp.VideoList[i].Author.BackgroundImage = favoriteList.VideoList[i].Author.BackgroundImage
		resp.VideoList[i].Author.IsFollow = favoriteList.VideoList[i].Author.IsFollow
		resp.VideoList[i].Author.FollowCount = favoriteList.VideoList[i].Author.FollowCount
		resp.VideoList[i].Author.FollowerCount = favoriteList.VideoList[i].Author.FollowerCount
		resp.VideoList[i].Author.TotalFavorited = favoriteList.VideoList[i].Author.TotalFavorited
		resp.VideoList[i].Author.FavoriteCount = favoriteList.VideoList[i].Author.FavoriteCount
		resp.VideoList[i].Author.WorkCount = favoriteList.VideoList[i].Author.WorkCount

	}
	resp.StatusCode = 0
	resp.StatusMsg = ""

	return resp, nil
}
