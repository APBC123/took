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

type GetVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoLogic) GetVideo(in *video.FeedRequest) (*video.FeedResponse, error) {
	resp := new(video.FeedResponse)
	uc := new(helper.UserClaim)
	var err error
	if in.Token != "" {
		uc, err = helper.AnalyzeToken(in.Token)
		if err != nil {
			return nil, err
		}
		//获取用户信息
		has, err1 := l.svcCtx.Engine.Where("id = ? AND enable  = ? AND deleted = ?", uc.Id, true, false).Get(new(models2.User))
		if err1 != nil {
			return nil, err
		}
		if !has {
			resp.StatusMsg = "The user is not enable"
			resp.StatusCode = -1
			resp.VideoList = make([]*video.Video, 0)
			resp.NextTime = -1
			return resp, nil
		}
	} else {
		uc.Id = 0
	}
	//获取视频列表
	//尝试从redis获取本次请求的数据
	list, err := l.svcCtx.RDB.Get(l.ctx, "LatestTime:"+strconv.FormatInt(in.LatestTime, 10)).Result() //需要将查询语句放至微服务中
	//redis中不存在数据
	if err != nil {
		vdList := make([]*models2.Video, 0)
		err = l.svcCtx.Engine.Limit(8, int(in.LatestTime)).Desc("id").Where("removed = ? AND deleted = ?", false, false).Find(&vdList) //获取投稿时间最近的8个视频
		if err != nil {
			return nil, err
		}
		listResp := make([]*video.Video, len(vdList))
		for i := range vdList {
			ur := new(models2.User)
			l.svcCtx.Engine.Where("id = ? AND enable = ? AND deleted = ?", vdList[i].AuthorId, true, false).Get(ur)
			listResp[i] = new(video.Video)
			listResp[i].Id = vdList[i].Id
			listResp[i].Title = vdList[i].Title
			listResp[i].CoverUrl = vdList[i].CoverUrl
			listResp[i].PlayUrl = vdList[i].PlayUrl
			listResp[i].CommentCount = vdList[i].CommentCount
			listResp[i].FavoriteCount = vdList[i].FavoriteCount
			listResp[i].Author = new(video.User)
			listResp[i].Author.Username = ur.Username
			listResp[i].Author.Id = ur.Id
			listResp[i].Author.FollowerCount = ur.FollowerCount
			listResp[i].Author.FollowCount = ur.FollowCount
			listResp[i].Author.Avatar = ur.Avatar
			listResp[i].Author.BackgroundImage = ur.BackgroundImage
			listResp[i].Author.Signature = ur.Signature
			has, _ := l.svcCtx.Engine.Where("user_id = ? AND fan_id = ?", ur.Id, uc.Id).Get(new(models2.Follow))
			if has {
				listResp[i].Author.IsFollow = true
			} else {
				listResp[i].Author.IsFollow = false
			}
			vd := new(models2.Video)
			listResp[i].Author.TotalFavorited, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).SumInt(vd, "favorite_count") //获赞总数
			listResp[i].Author.WorkCount, _ = l.svcCtx.Engine.Where("author_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Video))         //作品数
			listResp[i].Author.FavoriteCount, _ = l.svcCtx.Engine.Where("user_id = ? AND deleted = ? AND removed = ?", ur.Id, false, false).Count(new(models2.Favorite))    //喜欢数

			has, _ = l.svcCtx.Engine.Where("video_id = ? AND user_id = ? AND removed = ? AND deleted = ?", vdList[i].Id, uc.Id, false, false).Get(new(models2.Favorite))
			if has {
				listResp[i].IsFavorite = true
			} else {
				listResp[i].IsFavorite = false
			}
		}
		resp.VideoList = listResp
		//将本次请求到的数据写入redis,LatestTime作为key
		s, err := json.Marshal(resp.VideoList)
		if err != nil {
			return nil, err
		}
		l.svcCtx.RDB.Set(l.ctx, "LatestTime:"+strconv.FormatInt(in.LatestTime, 10), s, time.Second*time.Duration(define.CacheExpire+helper.Random()))
		l.svcCtx.RDB.Set(l.ctx, "LatestTimeCount:"+strconv.FormatInt(in.LatestTime, 10), 0, time.Second*time.Duration(define.CacheExpire+helper.Random())) //设置计数
		resp.StatusMsg = ""
		resp.StatusCode = 0
		resp.NextTime = in.LatestTime + int64(len(vdList))
	} else { //成功从redis中获取数据
		resp.VideoList = make([]*video.Video, 0)
		err = json.Unmarshal([]byte(list), &resp.VideoList)
		if err != nil {
			return nil, err
		}
		//计数加一并判断是否为热点数据
		if times, _ := l.svcCtx.RDB.Incr(l.ctx, "LatestTimeCount:"+strconv.FormatInt(in.LatestTime, 10)).Uint64(); times > define.TimesOffset { //当1~2分钟内的访问量超过设定的阈值将设置记录永不过期
			l.svcCtx.RDB.Expire(l.ctx, "LatestTime:"+strconv.FormatInt(in.LatestTime, 10), -1)
			l.svcCtx.RDB.Expire(l.ctx, "LatestTimeCount"+strconv.FormatInt(in.LatestTime, 10), -1)
		}
		resp.StatusMsg = ""
		resp.StatusCode = 0
		resp.NextTime = in.LatestTime + int64(len(resp.VideoList))
	}

	return resp, nil
}
