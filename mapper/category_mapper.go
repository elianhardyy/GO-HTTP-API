package mapper

import (
	"server/dto"
	"server/models"
)

func ToCategoryModel(dto dto.CategoryDto)models.Category{
	return models.Category{
		Name: dto.Name,
	}
}
func ToCategoryDto(category models.Category)dto.CategoryDto{
	return dto.CategoryDto{
		Name: category.Name,
	}
}

func ToCategoryModelList(dtos []dto.CategoryDto) []models.Category{
	category := make([]models.Category,len(dtos))
	for i, item := range dtos {
		category[i] = ToCategoryModel(item)
	}
	return category
}