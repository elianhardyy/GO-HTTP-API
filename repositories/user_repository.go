package repositories

import (
	"fmt"
	"math/rand"
	"server/models"
	"server/utils"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveOrUpdate(user models.User)(models.User,error)
	FindAll()[]models.User
	FindByEmail(email string, password string) (models.User,error)
	FindById(id uint) models.User
	SingleEmail(email string) (models.User,error)
	Delete(id uint) error
	VerifyToken(token string) error
	TokenIsUsed(email string) error
	UpdateProfile(user models.User, id uint) (models.User,error)
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
	rand.Seed(time.Now().UnixNano())
	randomNumber := 100000 + rand.Intn(900000)
	randomString := strconv.Itoa(randomNumber)
	err := utils.SendVerificationEmail(user.Email,randomString)
	if err != nil{
		fmt.Print("error")
		return user,err
	}
	token := models.Token{
		UserID: user.ID,
		TokenString: randomString,
	}
	u.DB.Create(&token)
	return user,nil
}
func (u *userRepository) FindAll()[]models.User{
	var users []models.User
	u.DB.Find(&users)
	return users
}
func (u *userRepository) FindByEmail(email string, password string) (models.User,error){
	var users models.User
	u.DB.Where("email = ? ",email).First(&users)
	err := bcrypt.CompareHashAndPassword([]byte(users.Password),[]byte(password))
	if err != nil{
		return users,err
	}
	return users,nil
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

func (u *userRepository) SingleEmail(email string) (models.User,error){
	var users models.User
	err := u.DB.Where("email = ?",email).Find(&users).Error
	if err != nil {
		return users,err
	}
	return users,nil
}

func (u *userRepository) Delete(id uint) error {
	var users models.User
	if err := u.DB.Where("id =?",id).Delete(&users).Error; err != nil{
		return err
	}
	return nil
}

func (u *userRepository) VerifyToken(token string)error{
	var tokenString models.Token
	// Cari token berdasarkan TokenString
	err := u.DB.Where("token_string = ?", token).First(&tokenString).Error
	if err != nil {
		// Jika tidak ditemukan, atau error lain
		return fmt.Errorf("errorr")
	}

	// Cek apakah IsUsed bernilai true
	if tokenString.IsUsed {
		return fmt.Errorf("token has already been used")
	}
	if time.Now().After(tokenString.ExpiredAt){
		return fmt.Errorf("token has expired")
	}
	err = u.DB.Where("token_string = ?",token).Updates(
		&models.Token{
			IsUsed: true,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) TokenIsUsed(email string) error {
	var user models.User
	err := u.DB.Where("email = ?",email).First(&user).Error
	if err != nil{
		return err
	}
	var token models.Token
	UserID := uint(user.ID)
	err = u.DB.Where("user_id = ?",UserID).First(&token).Error
	if err != nil{
		return err
	}
	if !token.IsUsed {
		return fmt.Errorf("token not verified")
	}
	return nil
}

func (u *userRepository) UpdateProfile(user models.User, id uint)(models.User,error){
	var users models.User
	err := u.DB.Where("id = ?",id).Updates(&models.User{
		Name: user.Name,
		Email: user.Email,
		Profile: user.Profile,
	}).Error
	if err != nil {
		return models.User{},err
	}
	return users,nil
}