package injection

import (
	"server/controllers"
	"server/repositories"
	"server/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserApi(db *gorm.DB) controllers.UserController {
	wire.Build(repositories.NewUserRepository,services.NewUserService,controllers.NewUserController)
	return controllers.UserController{}
}