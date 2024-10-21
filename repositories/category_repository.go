package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type CategoryRepository interface{
	SaveCategory(category models.Category)(models.Category,error)
	FetchAllCategory()([]models.Category,error)
	FindByIdCategory(id uint)(models.Category,error)
	UpdateCategory(categoryDto struct{
		Name string
	},id uint) (models.Category, error) 
	DeleteCategory(id uint)(string, error)
}
type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) CategoryRepository {
	return &categoryRepository{
		DB: DB,
	}
}

func (c *categoryRepository) SaveCategory(category models.Category)(models.Category,error){
	if err := c.DB.Create(&category).Error; err != nil{ //jika tidak ada error maka tambahi .Error
		return category,err
	}
	return category,nil
}

func (c *categoryRepository) FetchAllCategory()([]models.Category,error){
	var category []models.Category
	if err := c.DB.Preload("Products").Find(&category).Error; err != nil{
		return category, err
	}
	return category,nil
}

func (c *categoryRepository) FindByIdCategory(id uint)(models.Category,error){
	var category models.Category
	err := c.DB.Preload("Products").Where("id = ?",id).Find(&category).Error
	if err != nil{
		return category,err
	}
	return category,nil
}
func (c *categoryRepository) UpdateCategory(categoryDto struct{
	Name string
},id uint) (models.Category, error) {
	var category models.Category
	err := c.DB.Where("id = ?",id).Updates(models.Category{
		Name: categoryDto.Name,
	}).Error
	if err != nil{
		return category,err
	}
	return category,nil
}
func (c *categoryRepository) DeleteCategory(id uint)(string, error){
	var category models.Category
	err := c.DB.Where("id = ?",id).Delete(&category).Error
	if err != nil{
		return "failed",err
	}
	return "success",nil
}