package handler

import (
	pb "github.com/anjush-bhargavan/go_trade_chat/pkg/proto"
	"github.com/anjush-bhargavan/go_trade_chat/pkg/service/interfaces"
)
type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	svc interfaces.ChatServiceInter
}

func NewChatServiceServer(svc interfaces.ChatServiceInter) *ChatServiceServer {
	return &ChatServiceServer{
		svc: svc,
	}
}
