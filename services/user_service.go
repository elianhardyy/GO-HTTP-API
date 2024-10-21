package services

import (
	"server/dto"
	"server/mapper"
	"server/models"
	"server/repositories"
)

type UserService interface{
	SaveOrUpdate(dto dto.UserDto)(dto.UserDto,error)
	FindByEmail(email string, password string) (dto.UserLoginDto,error)
	EmailAuth(email string) string
	FindById(id uint) models.User
}

type userService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) UserService{
	return &userService{
		UserRepository:u,
	}
}
//implementation
func(u *userService) SaveOrUpdate(dto dto.UserDto)(dto.UserDto,error){
	userMapper := mapper.ToUserModel(dto)
	user, err := u.UserRepository.SaveOrUpdate(userMapper)
	if err != nil{
		return mapper.ToUserDto(user),err
	}
	return mapper.ToUserDto(user), nil
}
func (u *userService) FindByEmail(email string, password string) (dto.UserLoginDto,error){
	emails,err := u.UserRepository.FindByEmail(email,password)
	if err != nil{
		return dto.UserLoginDto{
			Email: emails,
			Password: "",
		},err
	}
	return dto.UserLoginDto{
		Email: emails,
		Password: "",
	},nil
}

func (u *userService) EmailAuth(email string) string{
	emails := u.UserRepository.SingleEmail(email)
	return emails
}

func (u *userService) FindById(id uint) models.User{
	user := u.UserRepository.FindById(id)
	return user
}
// func (u *UserService) FindAll()[]dto.UserDto{

// }
// func (u *UserResponse)Register(user *models.User) interface{} {
// 	config.DB.Create(&user)
// 	usr := &UserResponse{
// 		Name: user.Name,
// 		Email: user.Email,
// 	}
// 	return usr
// }