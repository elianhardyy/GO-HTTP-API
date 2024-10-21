package routes

import (
	"server/config"
	"server/injection"
)

func InitRoute() {
	userApi := injection.InitUserApiGen(config.DB)
	UserRoutes(userApi)
}