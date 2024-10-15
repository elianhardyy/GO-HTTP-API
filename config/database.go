package config

import (
	"os"
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	db,err := gorm.Open(mysql.Open(os.Getenv("DATABASE_CONNECTION")),&gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	DB = db
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})	

	//seeders.SeedRoles(db)
}