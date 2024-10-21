package migration

import (
	"server/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB){
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
}