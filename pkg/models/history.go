package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UserID     uint
	ReceiverID uint
	Message    string
}
