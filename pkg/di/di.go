package di

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_chat/config"
	"github.com/anjush-bhargavan/go_trade_chat/pkg/db"
	"github.com/anjush-bhargavan/go_trade_chat/pkg/handler"
	"github.com/anjush-bhargavan/go_trade_chat/pkg/repo"
	"github.com/anjush-bhargavan/go_trade_chat/pkg/server"
	"github.com/anjush-bhargavan/go_trade_chat/pkg/service"
)

// Init initializes the application by setting up its dependencies.
func Init() {
	cnfg := config.LoadConfig()

	db := db.ConnectDB(cnfg)


	chatRepo := repo.NewChatRepo(db)

	chatService := service.NewChatService(chatRepo)

	chatServer := handler.NewChatServiceServer(chatService)

	err := server.NewGrpcUserServer(cnfg.GrpcPort, chatServer)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}

}
