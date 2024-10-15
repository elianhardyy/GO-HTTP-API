package dto

type UserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginDto struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password" validate:"required"`
}