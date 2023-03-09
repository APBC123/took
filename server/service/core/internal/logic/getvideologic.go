package logic

import (
	"context"
	"took/server/service/core/helper"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
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
		resp.VideoList[i].Author = helper.NewUser(list[i].Author)
	}

	return resp, nil
}
