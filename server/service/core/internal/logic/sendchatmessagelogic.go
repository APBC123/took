package logic

import (
	"context"
	"took/chat/rpc/types/chat"

	"took/server/service/core/internal/svc"
	"took/server/service/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendChatMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendChatMessageLogic {
	return &SendChatMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendChatMessageLogic) SendChatMessage(req *types.SendChatMessageRequest) (resp *types.SendChatMessageResponse, err error) {
	sendChatMessage, err := l.svcCtx.ChatRpc.SendChatMessage(l.ctx, &chat.SendChatMessageRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.SendChatMessageResponse)
	resp.StatusMsg = sendChatMessage.StatusMsg
	resp.StatusCode = sendChatMessage.StatusCode
	return resp, nil

}
