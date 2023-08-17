package logic

import (
	"context"
	"took/video/helper"
	models2 "took/video/models"
	"took/video/video/rpc/internal/svc"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *video.CommentListRequest) (*video.CommentListResponse, error) {
	resp := new(video.CommentListResponse)
	//验证Token
	if in.Token != "" {
		_, err := helper.AnalyzeToken(in.Token)
		if err != nil {
			return nil, err
		}
	}

	//查询评论信息
	ct := make([]*models2.Comment, 0)
	err := l.svcCtx.Engine.Where("video_id  = ? AND deleted = ? AND removed = ?", in.VideoId, false, false).Desc("create_time").Find(&ct)
	if err != nil {
		return nil, err
	}

	//查询视频作者id
	vd := new(models2.Video)
	has, err := l.svcCtx.Engine.Where("id = ? AND deleted = ? AND removed = ?", in.VideoId, false, false).Get(vd)
	if err != nil {
		return nil, err
	}
	if !has {
		resp.StatusCode = -1
		resp.StatusMsg = "The author doesn't exist"
		return resp, nil
	}
	authorId := vd.AuthorId

	//查询评论对应的用户信息
	ur := make([]video.User, len(ct))
	for i := range ur {
		user := new(models2.User)
		_, err = l.svcCtx.Engine.Where("id = ? AND deleted = ? AND enable = ?", ct[i].UserId, false, true).Get(user)
		if err != nil {
			return nil, err
		}
		ur[i].Username = user.Username
		ur[i].Id = user.Id
		ur[i].FollowerCount = user.FollowerCount
		ur[i].FollowCount = user.FollowCount
		ur[i].Avatar = user.Avatar
		ur[i].BackgroundImage = user.BackgroundImage
		ur[i].Signature = user.Signature
		has, _ = l.svcCtx.Engine.Where("user_id = ? AND fan_id = ?", authorId, user.Id).Get(new(models2.Follow))
		if has {
			ur[i].IsFollow = true
		} else {
			ur[i].IsFollow = false
		}
		ur[i].TotalFavorited, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", user.Id, false, false).SumInt(vd, "favorite_count")
		ur[i].WorkCount, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", user.Id, false, false).Count(new(models2.Video))
		ur[i].FavoriteCount, _ = l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", user.Id, false, false).Count(new(models2.Favorite))
	}
	//生成CommentList
	ctList := make([]*video.Comment, len(ct))
	for i := range ctList {
		ctList[i] = new(video.Comment)
		ctList[i].Id = ct[i].Id
		ctList[i].Content = ct[i].Content
		ctList[i].CreateDate = ct[i].CreateTime.String()
		ctList[i].User = &ur[i]
	}
	//
	resp.CommentList = ctList
	resp.StatusCode = 0
	resp.StatusMsg = ""

	return resp, nil
}
