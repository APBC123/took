package logic

import (
	"context"
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
	messageList := make([]*chat.Message, 0)
	err = l.svcCtx.Engine.Where("to_user_id = ?", uc.Id).Find(&messageList)
	if err != nil {
		return nil, err
	}
	resp.MessageList = messageList
	return resp, nil
}
