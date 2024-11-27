package seeders

import "gorm.io/gorm"

func Seeder(db *gorm.DB){
	SeedUsers(db)
	//SeedCategories(db)
}