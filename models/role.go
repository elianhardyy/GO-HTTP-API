package models

import (
	"time"
)

type Role struct
{
	ID 			uint	`gorm:"primaryKey" json:"id"`
	Role_Name 	string	`gorm:"unique"`
	CreatedAt 	time.Time // Automatically managed by GORM for creation time
	UpdatedAt 	time.Time
}
