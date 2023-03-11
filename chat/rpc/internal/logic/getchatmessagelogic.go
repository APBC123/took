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
	err = l.svcCtx.Engine.Where("(from_user_id = ? AND to_user_id = ? AND create_time > ?) OR (from_user_id = ? AND to_user_id = ? AND create_time > ?)", in.ToUserId, uc.Id, in.PreMsgTime, uc.Id, in.ToUserId, in.PreMsgTime).Find(&messageList)
	if err != nil {
		return nil, err
	}
	resp.MessageList = messageList
	resp.StatusMsg = ""
	resp.StatusCode = 0
	return resp, nil
}
