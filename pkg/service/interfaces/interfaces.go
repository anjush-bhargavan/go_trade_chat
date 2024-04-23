package interfaces

import (
	pb "github.com/anjush-bhargavan/go_trade_chat/pkg/proto"
)

type ChatServiceInter interface {
	CreateChatService(p *pb.Message) error
	FetchChatService(p *pb.ChatID) (*pb.ChatHistory,error)
}
