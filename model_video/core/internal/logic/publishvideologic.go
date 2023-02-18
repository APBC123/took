package logic

import (
	"context"
	"model_video/core/helper"
	"model_video/core/models"

	"model_video/core/internal/svc"
	"model_video/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	resp = new(types.PublishResponse)
	uc, err := helper.AnalyzeToken(req.Token)
	if err != nil {
		return nil, err
	}
	//验证Token
	has, err := l.svcCtx.Engine.Where("id = ? AND username = ? AND password = ?", uc.Id, uc.Username, uc.Password).Get(new(models.User))
	if !has {
		resp.StatusCode = -1
		resp.StatusMsg = "User doesn't match Token"
		return
	}
	//插入上传记录，上传已在handler层中完成
	vd := &models.Video{
		AuthorId:      uc.Id,
		PlayUrl:       req.PlayUrl,
		CoverUrl:      req.CoverUrl,
		Title:         req.Title,
		CommentCount:  0,
		FavoriteCount: 0,
	}
	_, err = l.svcCtx.Engine.Insert(vd)
	if err != nil {
		return nil, err
	}

	resp.StatusCode = 0
	resp.StatusMsg = ""

	return
}
