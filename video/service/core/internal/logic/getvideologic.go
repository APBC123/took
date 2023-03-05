package logic

import (
	"context"
	"took/video/service/core/internal/svc"
	"took/video/service/core/internal/types"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoLogic) GetVideo(req *types.FeedRequest) (resp *types.FeedResponse, err error) {
	videoFeed, err := l.svcCtx.VideoRpc.GetVideo(l.ctx, &video.FeedRequest{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	})
	if err != nil {
		return nil, err
	}
	list := videoFeed.VideoList
	resp = new(types.FeedResponse)
	resp.StatusMsg = videoFeed.StatusMsg
	resp.StatusCode = videoFeed.StatusCode
	resp.NextTime = videoFeed.NextTime
	resp.VideoList = make([]types.Video, len(list))
	for i := range resp.VideoList {
		resp.VideoList[i].Id = list[i].Id
		resp.VideoList[i].PlayUrl = list[i].PlayUrl
		resp.VideoList[i].CoverUrl = list[i].CoverUrl
		resp.VideoList[i].Title = list[i].Title
		resp.VideoList[i].FavoriteCount = list[i].FavoriteCount
		resp.VideoList[i].CommentCount = list[i].CommentCount
		resp.VideoList[i].IsFavorite = list[i].IsFavorite
		resp.VideoList[i].Author.Id = list[i].Author.Id
		resp.VideoList[i].Author.Username = list[i].Author.Username
		resp.VideoList[i].Author.FollowCount = list[i].Author.FollowCount
		resp.VideoList[i].Author.IsFollow = list[i].Author.IsFollow
		resp.VideoList[i].Author.FollowerCount = list[i].Author.FollowerCount
		resp.VideoList[i].Author.Avatar = list[i].Author.Avatar
		resp.VideoList[i].Author.BackgroundImage = list[i].Author.BackgroundImage
		resp.VideoList[i].Author.Signature = list[i].Author.Signature
		resp.VideoList[i].Author.TotalFavorited = list[i].Author.TotalFavorited
		resp.VideoList[i].Author.FavoriteCount = list[i].Author.FavoriteCount
		resp.VideoList[i].Author.WorkCount = list[i].Author.WorkCount
	}

	return resp, nil
}
