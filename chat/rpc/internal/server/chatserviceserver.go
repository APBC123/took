// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package server

import (
	"context"

	"took/chat/rpc/internal/logic"
	"took/chat/rpc/internal/svc"
	"took/chat/rpc/types/chat"
)

type ChatServiceServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedChatServiceServer
}

func NewChatServiceServer(svcCtx *svc.ServiceContext) *ChatServiceServer {
	return &ChatServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServiceServer) GetChatMessage(ctx context.Context, in *chat.ChatMessageRequest) (*chat.ChatMessageResponse, error) {
	l := logic.NewGetChatMessageLogic(ctx, s.svcCtx)
	return l.GetChatMessage(in)
}

func (s *ChatServiceServer) SendChatMessage(ctx context.Context, in *chat.SendChatMessageRequest) (*chat.SendChatMessageResponse, error) {
	l := logic.NewSendChatMessageLogic(ctx, s.svcCtx)
	return l.SendChatMessage(in)
}