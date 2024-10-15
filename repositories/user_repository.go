package repositories

import (
	"fmt"
	"server/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryContract interface {
	SaveOrUpdate(user models.User)(models.User,error)
	FindAll()[]models.User
	FindById(id uint)models.User
	FindByEmail(email string)models.User
	Delete(id uint)error
}
type UserRepository struct{
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository{
	return UserRepository{
		DB: DB,
	}
}

func (u *UserRepository) SaveOrUpdate(user models.User)(models.User,error){
	if err := u.DB.Create(&user).Error; err != nil{ //jika tidak ada error maka tambahi .Error
		return user,err
	}
	return user,nil
}
func (u *UserRepository) FindAll()[]models.User{
	var users []models.User
	u.DB.Find(&users)
	return users
}
func (u *UserRepository) FindByEmail(email string, password string) (string,error){
	var users models.User
	u.DB.Where("email = ? ",email).First(&users)
	err := bcrypt.CompareHashAndPassword([]byte(users.Password),[]byte(password))
	if err != nil{
		return "error",err
	}
	return users.Email,nil
}
func (u *UserRepository) FindById(id uint) models.User{
	var users models.User
	//u.DB.First(&users,id)
	//u.DB.Model(&users).Association("Roles").Find(&users)
	err := u.DB.Where("id = ? ",id).Preload("Roles").Find(&users).Error
	if err != nil {
		fmt.Println("error")
	}
	return users
}

func (u *UserRepository) SingleEmail(email string) string{
	var users models.User
	u.DB.Where("email = ?",email).Find(&users)
	return users.Email
}

func (u *UserRepository) Delete(id uint) error {
	var users models.User
	if err := u.DB.Where("id =?",id).Delete(&users).Error; err != nil{
		return err
	}
	return nil
}