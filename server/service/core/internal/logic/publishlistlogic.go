package logic

import (
	"context"
	"took/server/service/core/helper"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
	"took/video/video/rpc/types/video"

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
	publishList, err := l.svcCtx.VideoRpc.PublishList(l.ctx, &video.PublishListRequest{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.PublishListResponse)
	resp.StatusMsg = publishList.StatusMsg
	resp.StatusCode = publishList.StatusCode
	resp.VideoList = helper.NewVideoList(publishList.VideoList)
	return resp, nil
}
