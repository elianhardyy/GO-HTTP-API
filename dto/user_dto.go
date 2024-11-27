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

type TokenDto struct {
	Token string `json:"token"  validate:"required"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" `
}

type ProfileDto struct {
	Name    string `form:"name" binding:"required"`
	Email   string `form:"email" binding:"required,email"`
	Profile string `form:"profile" binding:"required"`
}