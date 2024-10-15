package routes

import (
	"server/config"
	"server/injection"
)

func InitRoute() {
	userApi := injection.InitUserApiGen(config.DB)
	//apis := http.NewServeMux()
	UserRoutes(userApi)
	//controllers.GoogleCallback()
	//controllers.GoogleAuth()
	// apis.Handle("/api",http.StripPrefix("/v1",UserRoutes(userApi)))
	//http.Handle("/api",http.StripPrefix("/v1",UserRoutes(userApi)))
	//http.Handle("/",UserRoutes(userApi))
}