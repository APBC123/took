package logic

import (
	"context"
	"model_video/core/helper"
	"model_video/core/models"

	"model_video/core/internal/svc"
	"model_video/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	//验证Token
	resp = new(types.CommentListResponse)
	uc, err := helper.AnalyzeToken(req.Token)
	if err != nil {
		return nil, err
	}
	has, err := l.svcCtx.Engine.Where("id = ? AND username = ? AND password = ?", uc.Id, uc.Username, uc.Password).Get(new(models.User))
	if err != nil {
		return nil, err
	}
	if !has {
		resp.StatusCode = -1
		resp.StatusMsg = "User doesn't match Token"
		return
	}
	uc = nil
	//查询评论信息
	ct := make([]*models.Comment, 0)
	err = l.svcCtx.Engine.Where("video_id  = ? AND deleted = ? AND removed = ?", req.VideoId, false, false).Desc("create_time").Find(&ct)
	if err != nil {
		return nil, err
	}

	//查询视频作者id
	vd := new(models.Video)
	has, err = l.svcCtx.Engine.Where("id = ? AND deleted = ? AND removed = ?", req.VideoId, false, false).Get(vd)
	if err != nil {
		return nil, err
	}
	if !has {
		resp.StatusCode = -1
		resp.StatusMsg = "The author doesn't exist"
		return
	}
	authorId := vd.AuthorId
	vd = nil

	//查询评论对应的用户信息
	ur := make([]types.User, len(ct))
	for i := range ur {
		user := new(models.User)
		_, err = l.svcCtx.Engine.Where("id = ? AND deleted = ? AND enable = ?", ct[i].UserId, false, true).Get(user)
		if err != nil {
			return nil, err
		}
		ur[i].Username = user.Username
		ur[i].Id = user.Id
		ur[i].FollowerCount = user.FollowerCount
		ur[i].FollowCount = user.FollowCount
		has, _ = l.svcCtx.Engine.Where("user_id = ? AND fan_id = ?", authorId, user.Id).Get(new(models.Follow))
		if has {
			ur[i].IsFollow = true
		} else {
			ur[i].IsFollow = false
		}
	}

	//生成commentList
	ctList := make([]types.Comment, len(ur))
	for i := range ctList {
		ctList[i].Id = ct[i].Id
		ctList[i].Content = ct[i].Content
		ctList[i].CreateDate = ct[i].CreateTime.String()
		ctList[i].User = ur[i]
	}
	resp.CommentList = ctList
	resp.StatusCode = 0
	resp.StatusMsg = ""

	return
}
