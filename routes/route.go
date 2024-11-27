package routes

import (
	"server/config"
	"server/injection"

	"github.com/gorilla/mux"
)
var R = mux.NewRouter()
func InitRoute() {
	userApi := injection.InitUserApiGen(config.DB)
	UserRoutes(userApi)
}