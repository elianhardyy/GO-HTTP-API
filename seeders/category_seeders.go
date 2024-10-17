package seeders

import (
	"server/models"

	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB){
	categories := []models.Category{
		{Name: "Food"},
		{Name: "Drink"},
	}
	for _,category := range categories{
		db.Create(&category)
	}
}