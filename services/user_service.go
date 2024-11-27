package services

import (
	"server/dto"
	"server/mapper"
	"server/models"
	"server/repositories"
)

type UserService interface{
	SaveOrUpdate(dto dto.UserDto)(dto.UserDto,error)
	FindByEmail(email string, password string) (dto.UserResponse,error)
	EmailAuth(email string) (dto.UserResponse,error)
	FindById(id uint) models.User
	VerifyTokenS(token string)error
	UpdateProfile(user dto.ProfileDto, id uint)(dto.ProfileDto,error)
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
func (u *userService) FindByEmail(email string, password string) (dto.UserResponse,error){
	users,err := u.UserRepository.FindByEmail(email,password)
	if err != nil{
		return dto.UserResponse{},err
	}
	err = u.UserRepository.TokenIsUsed(email)
	if err != nil {
		return dto.UserResponse{},err
	}
	return dto.UserResponse{
		ID: users.ID,
		Name: users.Name,
		Email: users.Email,
	},nil
}

func (u *userService) EmailAuth(email string) (dto.UserResponse,error){
	user,err := u.UserRepository.SingleEmail(email)
	if err != nil {
		return dto.UserResponse{},err
	}
	return dto.UserResponse{
		Name: user.Name,
		Email: user.Email,
	},nil
}

func (u *userService) FindById(id uint) models.User{
	user := u.UserRepository.FindById(id)
	return user
}

func (u *userService) VerifyTokenS(token string)error{
	err := u.UserRepository.VerifyToken(token)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) UpdateProfile(user dto.ProfileDto, id uint)(dto.ProfileDto,error) {
	profileMapper := mapper.ToUserProfileModel(user)

	users,err := u.UserRepository.UpdateProfile(profileMapper,id)
	if err != nil {
		return dto.ProfileDto{},err
	}
	return mapper.ToUserProfileDto(users),nil
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