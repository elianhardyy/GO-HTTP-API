package services

import (
	"server/dto"
	"server/mapper"
	"server/models"
	"server/repositories"
)

// type UserResponse struct {
// 	Name, Email string
// }
// type UserRepositoryContract interface{
// 	SaveOrUpdate(dto dto.UserDto) (dto.UserDto, error)
// 	FindAll()[]dto.UserDto
// }

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) UserService{
	return UserService{
		UserRepository:u,
	}
}
//implementation
func(u *UserService) SaveOrUpdate(dto dto.UserDto)(dto.UserDto,error){
	userMapper := mapper.ToUserModel(dto)
	user, err := u.UserRepository.SaveOrUpdate(userMapper)
	if err != nil{
		return mapper.ToUserDto(user),err
	}
	return mapper.ToUserDto(user), nil
}
func (u *UserService) FindByEmail(email string, password string) (dto.UserLoginDto,error){
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

func (u *UserService) EmailAuth(email string) string{
	emails := u.UserRepository.SingleEmail(email)
	return emails
}

func (u *UserService) FindById(id uint) models.User{
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