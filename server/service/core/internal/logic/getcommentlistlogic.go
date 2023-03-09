package logic

import (
	"context"
	"took/server/service/core/helper"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
	"took/video/video/rpc/types/video"

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
	commentList, err := l.svcCtx.VideoRpc.GetCommentList(l.ctx, &video.CommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.CommentListResponse)

	list := commentList.CommentList
	resp.CommentList = make([]types.Comment, len(list))
	for i := range resp.CommentList {
		resp.CommentList[i].Id = list[i].Id
		resp.CommentList[i].Content = list[i].Content
		resp.CommentList[i].CreateDate = list[i].CreateDate
		resp.CommentList[i].User = helper.NewUser(list[i].User)

	}
	resp.StatusMsg = commentList.StatusMsg
	resp.StatusCode = commentList.StatusCode
	return resp, nil
}
