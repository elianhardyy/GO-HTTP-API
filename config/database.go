package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection()(*gorm.DB,error) {
	db,err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/gochatapp?parseTime=true"),&gorm.Config{})
	if err != nil {
		//panic("could not connect to the database")
		return DB,err
	}
	DB = db
	// db.AutoMigrate(&models.Role{})
	// db.AutoMigrate(&models.User{})	

	// seeders.SeedRoles(db)
	return DB,nil
}