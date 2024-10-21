package seeders

import (
	"server/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB){
	roles := []models.Role{
		{Role_Name : "admin"},
		{Role_Name : "user"},
	}
	db.Create(&roles)
}