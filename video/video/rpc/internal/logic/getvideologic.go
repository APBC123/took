package logic

import (
	"context"
	"took/video/helper"
	models2 "took/video/models"
	"took/video/video/rpc/internal/svc"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoLogic) GetVideo(in *video.FeedRequest) (*video.FeedResponse, error) {
	resp := new(video.FeedResponse)
	uc := new(helper.UserClaim)
	var err error
	if in.Token != "" {
		uc, err = helper.AnalyzeToken(in.Token)
		if err != nil {
			return nil, err
		}
		//获取用户信息
		has, err1 := l.svcCtx.Engine.Where("id = ? AND enable  = ? AND deleted = ?", uc.Id, true, false).Get(new(models2.User))
		if err1 != nil {
			return nil, err
		}
		if !has {
			resp.StatusMsg = "The user is not enable"
			resp.StatusCode = -1
			resp.VideoList = make([]*video.Video, 0)
			resp.NextTime = -1
			return resp, nil
		}
	} else {
		uc.Id = 0
	}
	//获取视频列表
	vdList := make([]*models2.Video, 0)
	err = l.svcCtx.Engine.Limit(8, int(in.LatestTime)).Desc("id").Find(&vdList) //获取投稿时间最近的8个视频
	if err != nil {
		return nil, err
	}
	listResp := make([]*video.Video, len(vdList))
	for i := range vdList {
		ur := new(models2.User)
		l.svcCtx.Engine.Where("id = ? AND enable = ? AND deleted = ?", vdList[i].AuthorId, true, false).Get(ur)
		listResp[i] = new(video.Video)
		listResp[i].Id = vdList[i].Id
		listResp[i].Title = vdList[i].Title
		listResp[i].CoverUrl = vdList[i].CoverUrl
		listResp[i].PlayUrl = vdList[i].PlayUrl
		listResp[i].CommentCount = vdList[i].CommentCount
		listResp[i].FavoriteCount = vdList[i].FavoriteCount
		listResp[i].Author = new(video.User)
		listResp[i].Author.Username = ur.Username
		listResp[i].Author.Id = ur.Id
		listResp[i].Author.FollowerCount = ur.FollowerCount
		listResp[i].Author.FollowCount = ur.FollowCount
		listResp[i].Author.Avatar = ur.Avatar
		listResp[i].Author.BackgroundImage = ur.BackgroundImage
		listResp[i].Author.Signature = ur.Signature
		has, _ := l.svcCtx.Engine.Where("user_id = ? AND fan_id = ?", ur.Id, uc.Id).Get(new(models2.Follow))
		if has {
			listResp[i].Author.IsFollow = true
		} else {
			listResp[i].Author.IsFollow = false
		}
		vd := new(models2.Video)
		listResp[i].Author.TotalFavorited, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).SumInt(vd, "favorite_count") //获赞总数
		listResp[i].Author.WorkCount, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Video))         //作品数
		listResp[i].Author.FavoriteCount, _ = l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Favorite))    //喜欢数

		has, _ = l.svcCtx.Engine.Where("video_id = ? AND user_id = ? AND removed = ? AND deleted = ?", vdList[i].Id, uc.Id, false, false).Get(new(models2.Favorite))
		if has {
			listResp[i].IsFavorite = true
		} else {
			listResp[i].IsFavorite = false
		}
	}
	resp.VideoList = listResp
	resp.StatusMsg = ""
	resp.StatusCode = 0
	resp.NextTime = in.LatestTime + 8

	return resp, nil
}
