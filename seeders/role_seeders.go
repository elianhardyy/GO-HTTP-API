package seeders

import (
	"fmt"
	"server/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB){
	roles := []models.Role{
		{Role_Name : "admin"},
		{Role_Name : "user"},
	}
	// for _, role := range roles {
	// }
	db.Create(&roles)
	fmt.Println("seed roles to roles table")
}