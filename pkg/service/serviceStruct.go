package service

import (
	// pb "github.com/anjush-bhargavan/go_trade_chat/pkg/proto"
	// "github.com/anjush-bhargavan/go_trade_chat/pkg/repo"
	"github.com/anjush-bhargavan/go_trade_chat/pkg/repo/interfaces"
	inter "github.com/anjush-bhargavan/go_trade_chat/pkg/service/interfaces"
)
type chatService struct {
	repo  interfaces.ChatRepoInter
}

func NewChatService(repo interfaces.ChatRepoInter) inter.ChatServiceInter {
	return &chatService{
		repo: repo,
	}
}
