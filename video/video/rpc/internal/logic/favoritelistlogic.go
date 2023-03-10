package logic

import (
	"context"
	"took/video/helper"
	models2 "took/video/models"

	"took/video/video/rpc/internal/svc"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *video.FavoriteListRequest) (*video.FavoriteListResponse, error) {
	if in.Token != "" {
		_, err := helper.AnalyzeToken(in.Token)
		if err != nil {
			return nil, err
		}
	}
	resp := new(video.FavoriteListResponse)
	FavoriteList := make([]*models2.Favorite, 0)
	err := l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", in.UserId, false, false).Find(&FavoriteList)
	if err != nil {
		return nil, err
	}
	resp.VideoList = make([]*video.Video, len(FavoriteList))
	for i := range resp.VideoList {
		resp.VideoList[i] = new(video.Video)
		vd := new(models2.Video)
		_, err = l.svcCtx.Engine.Where("id = ? AND deleted = ? AND removed = ?", FavoriteList[i].VideoId, false, false).Get(vd)
		if err != nil {
			return nil, err
		}
		ur := new(models2.User)
		_, err = l.svcCtx.Engine.Where("id = ? AND enable = ? AND deleted = ?", vd.AuthorId, true, false).Get(ur)
		if err != nil {
			return nil, err
		}
		resp.VideoList[i].IsFavorite = true
		resp.VideoList[i].Id = vd.Id
		resp.VideoList[i].Title = vd.Title
		resp.VideoList[i].PlayUrl = vd.PlayUrl
		resp.VideoList[i].CoverUrl = vd.CoverUrl
		resp.VideoList[i].CommentCount = vd.CommentCount
		resp.VideoList[i].FavoriteCount = vd.FavoriteCount
		resp.VideoList[i].Author = new(video.User)
		resp.VideoList[i].Author.Id = ur.Id
		resp.VideoList[i].Author.Username = ur.Username
		resp.VideoList[i].Author.Signature = ur.Signature
		resp.VideoList[i].Author.Avatar = ur.Avatar
		resp.VideoList[i].Author.BackgroundImage = ur.BackgroundImage
		resp.VideoList[i].Author.FollowerCount = ur.FollowerCount
		resp.VideoList[i].Author.FollowCount = ur.FollowCount
		resp.VideoList[i].Author.TotalFavorited, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).SumInt(vd, "favorite_count") //获赞总数
		resp.VideoList[i].Author.WorkCount, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Video))         //作品数
		resp.VideoList[i].Author.FavoriteCount, _ = l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Favorite))    //喜欢数
		has, _ := l.svcCtx.Engine.Where("user_id = ? AND fan_id = ? AND deleted = ?", ur.Id, in.UserId).Get(new(models2.Follow))
		if has {
			resp.VideoList[i].Author.IsFollow = true
		} else {
			resp.VideoList[i].Author.IsFollow = false
		}
	}
	resp.StatusMsg = ""
	resp.StatusCode = 0

	return resp, err
}
