package injection

import (
	"server/controllers"
	"server/repositories"
	"server/services"

	"gorm.io/gorm"
)

func InitUserApiGen(db *gorm.DB) controllers.UserController{
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	return userController
}