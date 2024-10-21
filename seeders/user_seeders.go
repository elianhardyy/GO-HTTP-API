package seeders

import (
	"server/models"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	roleAdmin := models.Role{Role_Name: "Admin"}
	roleUser := models.Role{Role_Name: "User"}
	db.Create(&roleAdmin)
	db.Create(&roleUser)

	user1 := models.User{Name: "Admin", Email: "admin@mail.com", Password: "12345678"}
	user2 := models.User{Name: "User", Email: "user@mail.com", Password: "12345678"}

	db.Create(&user1)
	db.Create(&user2)

	db.Model(&user1).Association("Roles").Append(&roleAdmin)
	db.Model(&user2).Association("Roles").Append(&roleUser)
}