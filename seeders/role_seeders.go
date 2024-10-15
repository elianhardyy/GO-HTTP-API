package seeders

import (
	"fmt"
	"server/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB){
	roles := []models.Role{
		{Role_Name : "admin"},
		{Role_Name:"user"},
	}
	for _, role := range roles {
		db.FirstOrCreate(&roles, models.Role{Role_Name: role.Role_Name})
	}
	fmt.Println("seed roles to roles table")
}