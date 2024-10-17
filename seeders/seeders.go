package seeders

import "gorm.io/gorm"

func Seeder(db *gorm.DB){
	SeedRoles(db)
	SeedCategories(db)
}