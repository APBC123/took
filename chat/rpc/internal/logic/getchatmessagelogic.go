package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	"took/chat/define"
	"took/chat/rpc/internal/svc"
	"took/chat/rpc/types/chat"
	"took/video/helper"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatMessageLogic {
	return &GetChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChatMessageLogic) GetChatMessage(in *chat.ChatMessageRequest) (*chat.ChatMessageResponse, error) {
	resp := new(chat.ChatMessageResponse)
	uc, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp.MessageList = make([]*chat.Message, 0)
	list, err := l.svcCtx.RDB.Get(l.ctx, "ChatMessage:"+strconv.FormatInt(uc.Id, 10)+"to"+strconv.FormatInt(in.ToUserId, 10)).Result()
	if err != nil {
		err = l.svcCtx.Engine.Where("(from_user_id = ? AND to_user_id = ? AND create_time > ?) OR (from_user_id = ? AND to_user_id = ? AND create_time > ?)", in.ToUserId, uc.Id, in.PreMsgTime, uc.Id, in.ToUserId, in.PreMsgTime).Find(&resp.MessageList)
		if err != nil {
			return nil, err
		}
		s, err := json.Marshal(resp.MessageList)
		if err != nil {
			return nil, err
		}
		l.svcCtx.RDB.Set(l.ctx, "ChatMessage:"+strconv.FormatInt(uc.Id, 10)+"to"+strconv.FormatInt(in.ToUserId, 10), s, time.Second*time.Duration(define.ChatMessageExpire))
	} else {
		err = json.Unmarshal([]byte(list), &resp.MessageList)
		if err != nil {
			return nil, err
		}
		if resp.MessageList[len(resp.MessageList)-1].CreateTime <= in.PreMsgTime {
			resp.MessageList = make([]*chat.Message, 0)
		}
	}
	resp.StatusMsg = ""
	resp.StatusCode = 0
	return resp, nil
}
