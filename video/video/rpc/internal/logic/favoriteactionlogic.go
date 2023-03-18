package logic

import (
	"context"
	"strconv"
	"took/video/helper"
	models2 "took/video/models"
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
	uc, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(video.FavoriteActionResponse)
	//点赞取消
	if in.ActionType == 2 {
		session := l.svcCtx.Engine.NewSession()
		defer session.Close()
		if err = session.Begin(); err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update favorite set removed = true where video_id = ? and user_id = ? and removed = ? and deleted = ?", in.VideoId, uc.Id, false, false)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update video set favorite_count = favorite_count-1 where id = ? and removed = ? and deleted = ?", in.VideoId, false, false)
		if err != nil {
			return nil, err
		}
		vd := new(models2.Video)
		_, err = l.svcCtx.Engine.Where("id = ? AND removed = ? AND deleted = ?", in.VideoId, false, false).Get(vd)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update user set favorite_count = favorite_count-1 where id = ? and enable = ? and deleted = ?", uc.Id, true, false)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update user set total_favorited = total_favorited-1 where id = ? and enable = ? and deleted = ?", vd.AuthorId, true, false)
		if err != nil {
			return nil, err
		}
		if err = session.Commit(); err != nil {
			return nil, err
		}

		resp.StatusCode = 0
		resp.StatusMsg = "点赞取消"
	}
	//点赞
	if in.ActionType == 1 {
		favoriteRecord := new(models2.Favorite)
		favoriteRecord.VideoId = in.VideoId
		favoriteRecord.UserId = uc.Id

		session := l.svcCtx.Engine.NewSession()
		defer session.Close()
		if err = session.Begin(); err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Insert(favoriteRecord)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update video set favorite_count = favorite_count+1 where id = ? and removed = ? and deleted = ?", in.VideoId, false, false)
		if err != nil {
			return nil, err
		}
		vd := new(models2.Video)
		_, err = l.svcCtx.Engine.Where("id = ? AND removed = ? AND deleted = ?", in.VideoId, false, false).Get(vd)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update user set favorite_count = favorite_count+1 where id = ? and enable = ? and deleted = ?", uc.Id, true, false)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Engine.Exec("update user set total_favorited = total_favorited+1 where id = ? and enable = ? and deleted = ?", vd.AuthorId, true, false)
		if err != nil {
			return nil, err
		}
		if err = session.Commit(); err != nil {
			return nil, err
		}
		resp.StatusMsg = "点赞成功"
		resp.StatusCode = 0
	}
	//删除Redis中对应的缓存
	l.svcCtx.RDB.Del(l.ctx, "FavoriteList_UserId:"+strconv.FormatInt(uc.Id, 10))
	return resp, nil
}
