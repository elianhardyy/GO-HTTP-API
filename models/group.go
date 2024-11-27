package models

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name	string
	RoomGroup *string
	Description *string
	Users	*[]User `gorm:"many2many:user_groups"`
	
}

type UserGroup struct{
	UserID   uint      `gorm:"primaryKey"`
	User	 User
	JoinedAt time.Time
}