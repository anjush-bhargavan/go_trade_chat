package service

import (
	"sort"

	"github.com/anjush-bhargavan/go_trade_chat/pkg/models"
	pb "github.com/anjush-bhargavan/go_trade_chat/pkg/proto"
)

func (c *chatService) CreateChatService(p *pb.Message) error {
	chat := &models.History{
		UserID:     uint(p.User_ID),
		ReceiverID: uint(p.Receiver_ID),
		Message:    p.Content,
	}
	err := c.repo.Createchat(chat)
	if err != nil {
		return err
	}

	return nil
}

func (c *chatService) FetchChatService(p *pb.ChatID) (*pb.ChatHistory, error) {

	userHistory, err := c.repo.Findchat(uint(p.User_ID), uint(p.Receiver_ID))
	if err != nil {
		return nil, err
	}
	receiverHistory, err := c.repo.Findchat(uint(p.Receiver_ID), uint(p.User_ID))
	if err != nil {
		return nil, err
	}
	var chats []*pb.Message
	for _, msg := range *userHistory {
		chats = append(chats, &pb.Message{
			Chat_ID:     uint32(msg.ID),
			User_ID:     uint32(msg.UserID),
			Receiver_ID: uint32(msg.ReceiverID),
			Content:     msg.Message,
		})
	}
	for _, msg := range *receiverHistory {
		chats = append(chats, &pb.Message{
			Chat_ID:     uint32(msg.ID),
			User_ID:     uint32(msg.UserID),
			Receiver_ID: uint32(msg.ReceiverID),
			Content:     msg.Message,
		})
	}
	sortByChatID(chats)
	return &pb.ChatHistory{
		Chats: chats,
	}, nil
}

func sortByChatID(chats []*pb.Message) {
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].Chat_ID < chats[j].Chat_ID
	})
}
