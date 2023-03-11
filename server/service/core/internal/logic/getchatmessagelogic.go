package logic

import (
	"context"
	"took/chat/rpc/types/chat"
	"took/server/service/core/helper"

	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatMessageLogic {
	return &GetChatMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatMessageLogic) GetChatMessage(req *types.ChatMessageRequest) (resp *types.ChatMessageResponse, err error) {
	chatMessage, err := l.svcCtx.ChatRpc.GetChatMessage(l.ctx, &chat.ChatMessageRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.ChatMessageResponse)
	resp.MessageList = helper.NewChatMessageList(chatMessage.MessageList)
	resp.StatusMsg = chatMessage.StatusMsg
	resp.StatusCode = chatMessage.StatusCode
	return resp, nil

}
