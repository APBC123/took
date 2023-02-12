package logic

import (
	"context"
	"errors"
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
	uc, err := helper.AnalyzeToken(req.Token)
	if err != nil {
		return nil, err
	}
	if uc.Id != req.UserId {
		return nil, errors.New("UserId doesn't match Token")
	}
	vd := make([]types.Video, 0)
	resp = new(types.PublishListResponse)
	user := new(models.User)
	//获取用户基本信息
	has, err := l.svcCtx.Engine.Where("id = ? AND deleted = ? AND enable = ?", req.UserId, false, true).Get(user)
	if err != nil {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = err.Error()
		return
	}
	if !has {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = "the user doesn't exist"
		return
	}
	//查询用户关注数
	fcnt, err := l.svcCtx.Engine.Where("fan_id = ?", req.UserId).Count(new(models.Follow))
	if err != nil {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = err.Error()
		return
	}
	//查询用户粉丝数
	frcnt, err := l.svcCtx.Engine.Where("user_id = ?", req.UserId).Count(new(models.Follow))
	if err != nil {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = err.Error()
		return
	}
	//查询用户投稿的视频目录
	err = l.svcCtx.Engine.Table("video").Where("author_id = ? AND removed = ? AND deleted = ?", user.Id, false, false).Find(&vd)
	if err != nil {
		resp.Response.StatusCode = -1
		resp.Response.StatusMsg = err.Error()
		return
	}

	for i, _ := range vd {
		vd[i].Author.Id = user.Id
		vd[i].Author.Name = user.Username
		vd[i].Author.IsFollow = true
		vd[i].Author.FollowerCount = frcnt
		vd[i].Author.FollowCount = fcnt
		vd[i].FavoriteCount, err = l.svcCtx.Engine.Where("video_id = ?", vd[i].Id).Count(new(models.Favorite))
		if err != nil {
			resp.Response.StatusCode = -1
			resp.Response.StatusMsg = err.Error()
			return
		}
		vd[i].CommentCount, err = l.svcCtx.Engine.Where("video_id = ?", vd[i].Id).Count(new(models.Comment))
		if err != nil {
			resp.Response.StatusCode = -1
			resp.Response.StatusMsg = err.Error()
			return
		}
	}
	resp.VideoList = vd
	resp.Response.StatusCode = 0
	resp.Response.StatusMsg = ""

	return
}
