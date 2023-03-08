package logic

import (
	"context"
	"took/video/helper"

	"took/video/video/rpc/internal/svc"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *video.FavoriteActionRequest) (*video.FavoriteActionResponse, error) {
	_, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	if in.ActionType == 1 {

	}
	if in.ActionType == 2 {

	}

	resp := new(video.FavoriteActionResponse)
	resp.StatusMsg = ""
	resp.StatusCode = 0
	return resp, nil
}
