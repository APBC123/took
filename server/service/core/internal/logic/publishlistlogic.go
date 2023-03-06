package logic

import (
	"context"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListRequest) (resp *types.PublishListResponse, err error) {
	publishList, err := l.svcCtx.VideoRpc.PublishList(l.ctx, &video.PublishListRequest{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	list := publishList.VideoList
	resp = new(types.PublishListResponse)
	resp.StatusMsg = publishList.StatusMsg
	resp.StatusCode = publishList.StatusCode
	resp.VideoList = make([]types.Video, len(list))
	for i := range list {
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
		resp.VideoList[i].IsFavorite = list[i].IsFavorite
		resp.VideoList[i].Id = list[i].Id
		resp.VideoList[i].Title = list[i].Title
		resp.VideoList[i].CoverUrl = list[i].CoverUrl
		resp.VideoList[i].PlayUrl = list[i].PlayUrl
		resp.VideoList[i].FavoriteCount = list[i].FavoriteCount
		resp.VideoList[i].CommentCount = list[i].CommentCount
	}
	return resp, nil
}
