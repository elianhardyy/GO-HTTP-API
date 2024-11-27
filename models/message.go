package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content			string
	User			[]User	`gorm:"many2many:user_messages"`
	Group			[]Group	`gorm:"many2many:message_groups"`
}

type UserMessage struct{
	GroupID			*uint
	UserID			*uint
	IsRead			bool
	ReadAt			*time.Time
	User			User
	Group			Group
}