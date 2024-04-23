package handler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/anjush-bhargavan/go_trade_chat/pkg/models"
	pb "github.com/anjush-bhargavan/go_trade_chat/pkg/proto"
)

type MessageHandle struct {
	MQue []models.History
	mu   sync.Mutex
}

var messageHandleObject = MessageHandle{}

func (c *ChatServiceServer) Connect(csi pb.ChatService_ConnectServer) error {

	errch := make(chan error)

	//recieve msgs
	go c.receiveFromStream(csi, errch)

	//send messages
	go sendToStream(csi, errch)

	return <-errch
}

func (c *ChatServiceServer) receiveFromStream(csi pb.ChatService_ConnectServer, errch chan error) {

	for {
		msg, err := csi.Recv()
		if err != nil {
			log.Printf("error in receiving msg from client : %v", err)
			errch <- err
		} else {
			messageHandleObject.mu.Lock()
			c.svc.CreateChatService(msg)
			messageHandleObject.MQue = append(messageHandleObject.MQue, models.History{
				UserID:     uint(msg.User_ID),
				ReceiverID: uint(msg.Receiver_ID),
				Message:    msg.Content,
			})
			messageHandleObject.mu.Unlock()

			// log.Printf("%v",messageHandleObject.MQue[len(messageHandleObject.MQue)-1])

		}
	}
}

func sendToStream(csi pb.ChatService_ConnectServer, errch chan error) {

	for {

		for {
			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			userID := messageHandleObject.MQue[0].UserID
			recieverID := messageHandleObject.MQue[0].ReceiverID
			message := messageHandleObject.MQue[0].Message

			messageHandleObject.mu.Unlock()

			if userID != recieverID {
				// fmt.Println(message)
				err := csi.Send(&pb.Message{
					User_ID:     uint32(userID),
					Receiver_ID: uint32(recieverID),
					Content:     message,
				})

				if err != nil {
					errch <- err
				}

				messageHandleObject.mu.Lock()

				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:]
				} else {
					messageHandleObject.MQue = []models.History{}
				}

				messageHandleObject.mu.Unlock()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}


func (c *ChatServiceServer) FetchHistory(ctx context.Context,p *pb.ChatID) (*pb.ChatHistory, error) {
	response, err := c.svc.FetchChatService(p)
	if err != nil {
		return response, err
	}
	return response, nil

}