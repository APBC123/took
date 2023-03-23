package logic

import (
	"context"
	"strconv"
	"time"
	helper2 "took/video/helper"
	models2 "took/video/models"

	"took/chat/rpc/internal/svc"
	"took/chat/rpc/types/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendChatMessageLogic {
	return &SendChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendChatMessageLogic) SendChatMessage(in *chat.SendChatMessageRequest) (*chat.SendChatMessageResponse, error) {
	uc, err := helper2.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(chat.SendChatMessageResponse)
	if in.ActionType == 1 {
		_, err = l.svcCtx.Engine.Insert(&models2.Message{
			FromUserId: uc.Id,
			ToUserId:   in.ToUserId,
			Content:    in.Content,
			CreateTime: time.Now().Unix(),
		})
		resp.StatusMsg = "发送成功"
		resp.StatusCode = 0
	} else {
		resp.StatusMsg = "发送失败"
		resp.StatusCode = 0
	}
	l.svcCtx.RDB.Unlink(l.ctx, "ChatMessage:"+strconv.FormatInt(uc.Id, 10)+"to"+strconv.FormatInt(in.ToUserId, 10))
	return resp, nil
}
