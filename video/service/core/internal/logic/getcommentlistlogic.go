package logic

import (
	"context"
	"took/video/service/core/internal/svc"
	"took/video/service/core/internal/types"
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
		resp.CommentList[i].User.Id = list[i].User.Id
		resp.CommentList[i].User.Username = list[i].User.Username
		resp.CommentList[i].User.FollowCount = list[i].User.FollowCount
		resp.CommentList[i].User.FollowerCount = list[i].User.FollowerCount
		resp.CommentList[i].User.IsFollow = list[i].User.IsFollow
		resp.CommentList[i].User.WorkCount = list[i].User.WorkCount
		resp.CommentList[i].User.TotalFavorited = list[i].User.TotalFavorited
		resp.CommentList[i].User.FavoriteCount = list[i].User.FavoriteCount
		resp.CommentList[i].User.Avatar = list[i].User.Avatar
		resp.CommentList[i].User.BackgroundImage = list[i].User.BackgroundImage
		resp.CommentList[i].User.Signature = list[i].User.Signature
	}
	resp.StatusMsg = commentList.StatusMsg
	resp.StatusCode = commentList.StatusCode
	return resp, nil
}
