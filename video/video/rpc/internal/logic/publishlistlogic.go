package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	"took/server/service/core/define"
	"took/video/helper"
	models2 "took/video/models"
	"took/video/video/rpc/internal/svc"
	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *video.PublishListRequest) (*video.PublishListResponse, error) {
	//验证Token
	//
	resp := new(video.PublishListResponse)
	uc := new(helper.UserClaim)
	if in.Token != "" {
		var err error
		uc, err = helper.AnalyzeToken(in.Token)
		if err != nil {
			return nil, err
		}
	}
	//尝试从redis中获取对应数据
	list, err := l.svcCtx.RDB.Get(l.ctx, "PublishList:"+strconv.FormatInt(in.UserId, 10)).Result()
	//redis中不存在数据
	if err != nil {
		ur := new(models2.User)
		//获取用户信息
		has, err := l.svcCtx.Engine.Where("id = ? AND deleted = ? AND enable = ?", in.UserId, false, true).Get(ur)
		if err != nil {
			return nil, err
		}
		//查询用户投稿的视频目录
		vd := make([]*models2.Video, 0)
		err = l.svcCtx.Engine.Where("author_id = ? AND removed = ? AND deleted = ?", ur.Id, false, false).Find(&vd)
		if err != nil {
			return nil, err
		}
		vdList := make([]*video.Video, len(vd))
		for i := range vdList {
			vdList[i] = new(video.Video)
			vdList[i].Author = new(video.User)
			vdList[i].Author.Id = ur.Id
			vdList[i].Author.FollowCount = ur.FollowCount
			vdList[i].Author.FollowerCount = ur.FollowerCount
			vdList[i].Author.Username = ur.Username
			if in.Token == "" {
				vdList[i].Author.IsFollow = false
			} else {
				has, err = l.svcCtx.Engine.Where("user_id = ? AND fan_id = ?", ur.Id, uc.Id).Get(new(models2.Follow))
				if err != nil {
					return nil, err
				}
				if has {
					vdList[i].Author.IsFollow = true
				} else {
					vdList[i].Author.IsFollow = false
				}
			}
			vdList[i].Author.Avatar = ur.Avatar
			vdList[i].Author.BackgroundImage = ur.BackgroundImage
			vdList[i].Author.Signature = ur.Signature
			vdList[i].Author.TotalFavorited, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).SumInt(vd, "favorite_count") //获赞总数
			vdList[i].Author.WorkCount, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Video))         //作品数
			vdList[i].Author.FavoriteCount, _ = l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Favorite))    //喜欢数

			vdList[i].Id = vd[i].Id
			vdList[i].CommentCount = vd[i].CommentCount
			vdList[i].FavoriteCount = vd[i].FavoriteCount
			vdList[i].CoverUrl = vd[i].CoverUrl
			vdList[i].PlayUrl = vd[i].PlayUrl
			has, _ = l.svcCtx.Engine.Where("video_id = ? AND user_id = ? AND removed = ? AND deleted = ?", vdList[i].Id, uc.Id, false, false).Get(new(models2.Favorite))
			if has {
				vdList[i].IsFavorite = true
			} else {
				vdList[i].IsFavorite = false
			}
			vdList[i].Title = vd[i].Title
		}
		resp.VideoList = vdList
		//将查询到的数据写入redis
		s, err := json.Marshal(resp.VideoList)
		if err != nil {
			return nil, err
		}
		l.svcCtx.RDB.Set(l.ctx, "PublishList:"+strconv.FormatInt(in.UserId, 10), s, time.Second*time.Duration(define.CacheExpire+helper.Random()))

		resp.StatusCode = 0
		resp.StatusMsg = ""
	} else { //成功从redis中获取数据
		resp.VideoList = make([]*video.Video, 0)
		err = json.Unmarshal([]byte(list), &resp.VideoList)
		if err != nil {
			return nil, err
		}
		resp.StatusMsg = ""
		resp.StatusCode = 0
	}

	return resp, nil
}
