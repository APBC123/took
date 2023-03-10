package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"
	"took/video/helper"
	"took/video/models"
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
	_, err = l.svcCtx.Engine.Exec("update user set work_count=work_count+1 where id=?", uc.Id)
	if err != nil {
		return nil, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "上传成功"

	return
}
