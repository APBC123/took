package logic

import (
	"context"
	"took/server/service/core/helper"
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
		resp.Comment.User = helper.NewUser(commentAction.Comment.User)
	}

	resp.StatusMsg = commentAction.StatusMsg
	resp.StatusCode = commentAction.StatusCode

	return resp, nil
}
