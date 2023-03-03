package logic

import (
	"context"
	"model_video/core/helper"
	"model_video/core/models"

	"model_video/core/internal/svc"
	"model_video/core/internal/types"

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
	uc := new(helper.UserClaim)
	if req.Token != "" {
		uc, err = helper.AnalyzeToken(req.Token)
		if err != nil {
			return nil, err
		}
		//获取用户信息
		has, err1 := l.svcCtx.Engine.Where("id = ? AND enable  = ? AND deleted = ?", uc.Id, true, false).Get(new(models.User))
		if err1 != nil {
			return nil, err
		}
		if !has {
			resp.StatusMsg = "The user doesn't enable"
			resp.StatusCode = -1
			resp.VideoList = make([]types.Video, 0)
			resp.NextTime = -1
			return
		}
	} else {
		uc.Id = 0
	}
	resp = new(types.FeedResponse)
	//获取视频列表
	vdList := make([]*models.Video, 0)
	err = l.svcCtx.Engine.Desc("id").Limit(10, int(req.LatestTime)).Find(&vdList) //获取投稿时间最近的10个视频
	if err != nil {
		return nil, err
	}
	listResp := make([]types.Video, 10)
	for i := range vdList {
		ur := new(models.User)
		has, _ := l.svcCtx.Engine.Where("id = ? AND enable = ? AND deleted = ?", vdList[i].AuthorId, true, false).Get(ur)
		listResp[i].Id = vdList[i].Id
		listResp[i].Title = vdList[i].Title
		listResp[i].CoverUrl = vdList[i].CoverUrl
		listResp[i].PlayUrl = vdList[i].PlayUrl
		listResp[i].CommentCount = vdList[i].CommentCount
		listResp[i].FavoriteCount = vdList[i].FavoriteCount
		listResp[i].Author.Username = ur.Username
		listResp[i].Author.Id = ur.Id
		listResp[i].Author.FollowerCount = ur.FollowerCount
		listResp[i].Author.FollowCount = ur.FollowCount
		has, _ = l.svcCtx.Engine.Where("user_id = ? AND fan_id = ?", ur.Id, uc.Id).Get(new(models.Follow))
		if has {
			listResp[i].Author.IsFollow = true
		} else {
			listResp[i].Author.IsFollow = false
		}
		has, _ = l.svcCtx.Engine.Where("video_id = ? AND user_id = ? AND removed = ? AND deleted = ?", vdList[i].Id, uc.Id, false, false).Get(new(models.Favorite))
		if has {
			listResp[i].IsFavorite = true
		} else {
			listResp[i].IsFavorite = false
		}
	}
	resp.VideoList = listResp
	resp.StatusMsg = ""
	resp.StatusCode = 0
	resp.NextTime = listResp[0].Id

	return
}
