package repositories

import (
	"fmt"
	"server/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveOrUpdate(user models.User)(models.User,error)
	FindAll()[]models.User
	FindByEmail(email string, password string) (string,error)
	FindById(id uint) models.User
	SingleEmail(email string) string
	Delete(id uint) error
}
type userRepository struct{
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository{
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) SaveOrUpdate(user models.User)(models.User,error){
	if err := u.DB.Create(&user).Error; err != nil{ //jika tidak ada error maka tambahi .Error
		return user,err
	}
	return user,nil
}
func (u *userRepository) FindAll()[]models.User{
	var users []models.User
	u.DB.Find(&users)
	return users
}
func (u *userRepository) FindByEmail(email string, password string) (string,error){
	var users models.User
	u.DB.Where("email = ? ",email).First(&users)
	err := bcrypt.CompareHashAndPassword([]byte(users.Password),[]byte(password))
	if err != nil{
		return "error",err
	}
	return users.Email,nil
}
func (u *userRepository) FindById(id uint) models.User{
	var users models.User
	//u.DB.First(&users,id)
	//u.DB.Model(&users).Association("Roles").Find(&users)
	err := u.DB.Where("id = ? ",id).Preload("Roles").Find(&users).Error
	if err != nil {
		fmt.Println("error")
	}
	return users
}

func (u *userRepository) SingleEmail(email string) string{
	var users models.User
	u.DB.Where("email = ?",email).Find(&users)
	return users.Email
}

func (u *userRepository) Delete(id uint) error {
	var users models.User
	if err := u.DB.Where("id =?",id).Delete(&users).Error; err != nil{
		return err
	}
	return nil
}