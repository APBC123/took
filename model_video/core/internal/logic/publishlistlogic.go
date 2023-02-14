package logic

import (
	"context"
	"model_video/core/helper"
	"model_video/core/internal/svc"
	"model_video/core/internal/types"
	"model_video/core/models"

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
	//验证Token
	//
	resp = new(types.PublishListResponse)
	uc, err := helper.AnalyzeToken(req.Token)
	if err != nil {
		return nil, err
	}
	if uc.Id != req.UserId {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = "User doesn't match Token"
		return
	}
	ur := new(models.User)
	//获取用户信息
	has, err := l.svcCtx.Engine.Where("id = ? AND deleted = ? AND enable = ?", req.UserId, false, true).Get(ur)
	if err != nil {
		return nil, err
	}
	//验证Token
	if ur.Password != uc.Password || ur.Username != uc.Username {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = "User doesn't match Token"
		return
	}
	if !has {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = "the user doesn't exist"
		return
	}
	user := new(types.User)
	user.Username = ur.Username
	user.Id = ur.Id
	user.FollowerCount = ur.FollowerCount
	user.FollowCount = ur.FollowCount
	user.IsFollow = false
	ur = nil
	//查询用户投稿的视频目录
	vd := make([]*models.Video, 0)
	err = l.svcCtx.Engine.Where("author_id = ? AND removed = ? AND deleted = ?", user.Id, false, false).Find(&vd)
	if err != nil {
		return nil, err
	}
	vdList := make([]types.Video, len(vd))
	for i := range vdList {
		vdList[i].Author = *user
		vdList[i].Id = vd[i].Id
		vdList[i].CommentCount = vd[i].CommentCount
		vdList[i].FavoriteCount = vd[i].FavoriteCount
		vdList[i].CoverUrl = vd[i].CoverUrl
		vdList[i].PlayUrl = vd[i].PlayUrl
		has, _ = l.svcCtx.Engine.Where("video_id = ? AND user_id = ?", vdList[i].Id, user.Id).Get(new(models.Favorite))
		if has {
			vdList[i].IsFavorite = true
		} else {
			vdList[i].IsFavorite = false
		}
		vdList[i].Title = vd[i].Title
	}
	resp.VideoList = vdList
	resp.Response.StatusCode = 0
	resp.Response.StatusMsg = ""

	return
}
