package mapper

import (
	"server/dto"
	"server/models"
)

func ToUserModel(dto dto.UserDto)models.User{
	return models.User{
		Name: dto.Name,
		Email: dto.Email,
		Password: dto.Password,
	}
}

func ToUserProfileModel(dto dto.ProfileDto)models.User{
	return models.User{
		Name: dto.Name,
		Email: dto.Email,
		Profile: dto.Profile,
	}
}

func ToUserModelList(dtos []dto.UserDto) []models.User{
	users := make([]models.User,len(dtos))
	for i,item := range dtos{
		users[i] = ToUserModel(item)
	}
	return users
}

func ToUserDto(user models.User) dto.UserDto{
	return dto.UserDto{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
}



func ToUserProfileDto(user models.User) dto.ProfileDto {
	return dto.ProfileDto{
		Name: user.Name,
		Email: user.Email,
		Profile : user.Profile,
	}
}

func ToUserDtoList(users []models.User) []dto.UserDto{
	dtos := make([]dto.UserDto,len(users))
	for i,item := range users{
		dtos[i] = ToUserDto(item)
	}
	return dtos
}
