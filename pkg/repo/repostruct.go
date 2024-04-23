package repo

import (
	"github.com/anjush-bhargavan/go_trade_chat/pkg/repo/interfaces"
	"gorm.io/gorm"
)


type chatRepo struct {
	DB *gorm.DB
}


func NewChatRepo (db *gorm.DB) interfaces.ChatRepoInter {
	return &chatRepo{
		DB: db,
	}
}