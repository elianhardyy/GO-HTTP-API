package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) CategoryRepository {
	return CategoryRepository{
		DB: DB,
	}
}

func (c *CategoryRepository) SaveCategory(category models.Category)(models.Category,error){
	if err := c.DB.Create(&category).Error; err != nil{ //jika tidak ada error maka tambahi .Error
		return category,err
	}
	return category,nil
}

func (c *CategoryRepository) FetchAll(category []models.Category)([]models.Category,error){
	if err := c.DB.Preload("Products").Find(&category).Error; err != nil{
		return category, err
	}
	return category,nil
}