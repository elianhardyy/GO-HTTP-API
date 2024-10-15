package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleReposotiry(DB *gorm.DB) RoleRepository{
	return RoleRepository{
		DB: DB,
	}
}

func (r *RoleRepository) FindByUserId(id uint)(models.Role, error){
	var role models.Role
	err := r.DB.Model(&role).Preload("Users").Where("user_id =?",id).Find(&role).Error
	if err != nil{
		return role,err
	}
	return role,nil
}