package server

import (
	"fmt"
	"log"
	"net"

	pb "github.com/anjush-bhargavan/go_trade_chat/pkg/proto"
	"google.golang.org/grpc"
)

// NewGrpcUserServer creates a new gRPC server for user-related operations.
func NewGrpcUserServer(port string, handler pb.ChatServiceServer) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s",port)
	lis, err := net.Listen("tcp",addr)
	if err != nil {
		log.Printf("error creating listener on %v",addr)
		return err
	}
	grpc := grpc.NewServer()
	pb.RegisterChatServiceServer(grpc,handler)

	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil 
}