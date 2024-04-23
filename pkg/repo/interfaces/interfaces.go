package interfaces

import "github.com/anjush-bhargavan/go_trade_chat/pkg/models"

type ChatRepoInter interface {
	Createchat(chat *models.History) error
	Findchat(userID, receiverID uint) (*[]models.History, error)
}
