package logic

import (
	"context"
	"took/video/helper"
	models2 "took/video/models"
	"took/video/video/rpc/internal/svc"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *video.PublishListRequest) (*video.PublishListResponse, error) {
	//验证Token
	//
	resp := new(video.PublishListResponse)

	if in.Token != "" {
		uc, err := helper.AnalyzeToken(in.Token)
		if err != nil {
			return nil, err
		}
		if uc.Id != in.UserId {
			resp.StatusCode = -1
			resp.StatusMsg = "User doesn't match Token"
			return resp, nil
		}
	}

	ur := new(models2.User)
	//获取用户信息
	has, err := l.svcCtx.Engine.Where("id = ? AND deleted = ? AND enable = ?", in.UserId, false, true).Get(ur)
	if err != nil {
		return nil, err
	}
	//查询用户投稿的视频目录
	vd := make([]*models2.Video, 0)
	err = l.svcCtx.Engine.Where("author_id = ? AND removed = ? AND deleted = ?", ur.Id, false, false).Find(&vd)
	if err != nil {
		return nil, err
	}
	vdList := make([]*video.Video, len(vd))
	for i := range vdList {
		vdList[i] = new(video.Video)
		vdList[i].Author = new(video.User)
		vdList[i].Author.Id = ur.Id
		vdList[i].Author.FollowCount = ur.FollowCount
		vdList[i].Author.FollowerCount = ur.FollowerCount
		vdList[i].Author.Username = ur.Username
		vdList[i].Author.IsFollow = false
		vdList[i].Author.Avatar = ur.Avatar
		vdList[i].Author.BackgroundImage = ur.BackgroundImage
		vdList[i].Author.Signature = ur.Signature
		vdList[i].Author.TotalFavorited, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).SumInt(vd, "favorite_count") //获赞总数
		vdList[i].Author.WorkCount, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Video))         //作品数
		vdList[i].Author.FavoriteCount, _ = l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Favorite))    //喜欢数

		vdList[i].Id = vd[i].Id
		vdList[i].CommentCount = vd[i].CommentCount
		vdList[i].FavoriteCount = vd[i].FavoriteCount
		vdList[i].CoverUrl = vd[i].CoverUrl
		vdList[i].PlayUrl = vd[i].PlayUrl
		has, _ = l.svcCtx.Engine.Where("video_id = ? AND user_id = ?", vdList[i].Id, ur.Id).Get(new(models2.Favorite))
		if has {
			vdList[i].IsFavorite = true
		} else {
			vdList[i].IsFavorite = false
		}
		vdList[i].Title = vd[i].Title
	}
	resp.VideoList = vdList
	resp.StatusCode = 0
	resp.StatusMsg = ""

	return resp, nil
}
