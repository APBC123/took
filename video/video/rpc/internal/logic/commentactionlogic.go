package logic

import (
	"context"
	"time"
	"took/video/helper"
	models2 "took/video/models"
	"took/video/video/rpc/internal/svc"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *video.CommentActionRequest) (*video.CommentActionResponse, error) {

	uc, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}

	resp := new(video.CommentActionResponse)
	if in.ActionType == 2 {
		ct := new(models2.Comment)
		ct.Deleted = true
		_, err = l.svcCtx.Engine.Where("id = ?", in.CommentId).Cols("deleted").Update(ct)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update video set comment_count = comment_count-1 where id = ?", in.VideoId)
		if err != nil {
			return nil, err
		}
		resp.Comment = nil
		resp.StatusMsg = "删除成功"
		resp.StatusCode = 0
	}
	if in.ActionType == 1 {
		ct := new(models2.Comment)
		ct.UserId = uc.Id
		ct.Content = in.CommentText
		ct.VideoId = in.VideoId
		ct.CreateTime = time.Now()
		_, err = l.svcCtx.Engine.Insert(ct)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update video set comment_count = comment_count+1 where id = ?", in.VideoId)
		if err != nil {
			return nil, err
		}
		user := new(models2.User)
		_, err = l.svcCtx.Engine.Where("id = ?", uc.Id).Get(user)
		if err != nil {
			return nil, err
		}
		vd := new(models2.Video)
		_, err = l.svcCtx.Engine.Where("id = ?", in.VideoId).Get(vd)
		if err != nil {
			return nil, err
		}
		resp.Comment = new(video.Comment)
		resp.Comment.Id = ct.Id
		resp.Comment.Content = ct.Content
		resp.Comment.CreateDate = ct.CreateTime.String()
		resp.Comment.User = new(video.User)
		resp.Comment.User.Username = user.Username
		resp.Comment.User.Id = user.Id
		resp.Comment.User.FollowerCount = user.FollowerCount
		resp.Comment.User.FollowCount = user.FollowCount
		resp.Comment.User.Avatar = user.Avatar
		resp.Comment.User.BackgroundImage = user.BackgroundImage
		resp.Comment.User.Signature = user.Signature
		resp.Comment.User.TotalFavorited, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", user.Id, false, false).SumInt(vd, "favorite_count") //获赞总数
		resp.Comment.User.WorkCount, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", user.Id, false, false).Count(new(models2.Video))         //作品数
		resp.Comment.User.FavoriteCount, _ = l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", user.Id, false, false).Count(new(models2.Favorite))    //喜欢数
		has, _ := l.svcCtx.Engine.Where("user_id = ? AND fan_id = ? AND deleted = ? AND removed = ?", vd.AuthorId, user.Id, false, false).Get(new(models2.Follow))
		if has {
			resp.Comment.User.IsFollow = true
		} else {
			resp.Comment.User.IsFollow = false
		}
		resp.StatusCode = 0
		resp.StatusMsg = "评论成功"
	}
	return resp, nil
}
