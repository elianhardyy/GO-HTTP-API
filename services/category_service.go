package services

import (
	"server/dto"
	"server/mapper"
	"server/models"
	"server/repositories"
)

type CategoryService interface {
	SaveCategory(dto dto.CategoryDto)(dto.CategoryDto,error)
	FetchAllCategory()([]models.Category,error)
	FindByIdCategory(id uint)(models.Category,error)
	UpdateCategory(categoryDto dto.CategoryDto,id uint) (models.Category, error) 
	DeleteCategory(id uint)(string, error)
}
type categoryService struct {
	CategoryRepository repositories.CategoryRepository
}

func NewCategoryService(c repositories.CategoryRepository) CategoryService{
	return &categoryService{
		CategoryRepository: c,
	}
}
func (s *categoryService) SaveCategory(dto dto.CategoryDto)(dto.CategoryDto,error){
	categoryMapper := mapper.ToCategoryModel(dto)
	category, err := s.CategoryRepository.SaveCategory(categoryMapper)
	if err != nil {
		return mapper.ToCategoryDto(category),err
	}
	return mapper.ToCategoryDto(category),nil
	
}

func (s *categoryService) FetchAllCategory()([]models.Category,error){
	category, err := s.CategoryRepository.FetchAllCategory()
	if err != nil {
		return category,err
	}
	return category,nil
}

func (s *categoryService) FindByIdCategory(id uint)(models.Category,error){
	category, err := s.CategoryRepository.FindByIdCategory(id)
	if err != nil{
		return category,err
	}
	return category,nil
}

func (s *categoryService) UpdateCategory(categoryDto dto.CategoryDto,id uint) (models.Category, error){
	category,err := s.CategoryRepository.UpdateCategory(struct{Name string}{Name: categoryDto.Name},id)
	if err != nil {
		return category,err
	}
	return category,nil
}

func (s *categoryService) DeleteCategory(id uint)(string, error){
	_,err := s.CategoryRepository.DeleteCategory(id)
	if err != nil {
		return "failed delete",err
	}
	return "success delete",nil
}