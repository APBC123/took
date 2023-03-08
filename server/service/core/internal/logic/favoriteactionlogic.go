package logic

import (
	"context"
	"took/video/video/rpc/types/video"

	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionRequest) (resp *types.FavoriteActionResponse, err error) {
	favoriteAction, err := l.svcCtx.VideoRpc.FavoriteAction(l.ctx, &video.FavoriteActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	resp = new(types.FavoriteActionResponse)
	resp.StatusCode = favoriteAction.StatusCode
	resp.StatusMsg = favoriteAction.StatusMsg
	return resp, nil
}
