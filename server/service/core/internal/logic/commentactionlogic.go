package logic

import (
	"context"
	"took/video/video/rpc/types/video"

	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionRequest) (resp *types.CommentActionResponse, err error) {
	commentAction, err := l.svcCtx.VideoRpc.CommentAction(l.ctx, &video.CommentActionRequest{
		Token:       req.Token,
		VideoId:     req.VideoId,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentId,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.CommentActionResponse)
	if commentAction.Comment != nil {
		resp.Comment = new(types.Comment)
		resp.Comment.Id = commentAction.Comment.Id
		resp.Comment.Content = commentAction.Comment.Content
		resp.Comment.CreateDate = commentAction.Comment.CreateDate
		resp.Comment.User.Id = commentAction.Comment.User.Id
		resp.Comment.User.Username = commentAction.Comment.User.Username
		resp.Comment.User.FollowCount = commentAction.Comment.User.FollowCount
		resp.Comment.User.FollowerCount = commentAction.Comment.User.FollowerCount
		resp.Comment.User.IsFollow = commentAction.Comment.User.IsFollow
		resp.Comment.User.WorkCount = commentAction.Comment.User.WorkCount
		resp.Comment.User.TotalFavorited = commentAction.Comment.User.TotalFavorited
		resp.Comment.User.FavoriteCount = commentAction.Comment.User.FavoriteCount
		resp.Comment.User.Avatar = commentAction.Comment.User.Avatar
		resp.Comment.User.BackgroundImage = commentAction.Comment.User.BackgroundImage
		resp.Comment.User.Signature = commentAction.Comment.User.Signature
	}

	resp.StatusMsg = commentAction.StatusMsg
	resp.StatusCode = commentAction.StatusCode

	return resp, nil
}
