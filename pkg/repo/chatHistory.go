package repo

import (
	"github.com/anjush-bhargavan/go_trade_chat/pkg/models"
)

// Createchat method for creating a new chat in db
func (c *chatRepo) Createchat(chat *models.History) error {
	if err := c.DB.Create(&chat).Error; err != nil {
		return err
	}
	return nil
}

// FindchatByEmail method finds the chat from database using user and receiver id
func (u *chatRepo) Findchat(userID, receiverID uint) (*[]models.History, error) {
	var chat []models.History

	if err := u.DB.Where("user_id = ? AND receiver_id = ?", userID, receiverID).Find(&chat).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}
