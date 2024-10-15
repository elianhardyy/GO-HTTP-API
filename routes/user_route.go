package routes

import (
	"net/http"
	"server/controllers"
	"server/middlewares"
)

func UserRoutes(api controllers.UserController) {
	http.HandleFunc("/create",api.Register)
	http.HandleFunc("/login",api.Login)
	http.Handle("/userme",middlewares.ProtectedHandler(http.HandlerFunc(api.UserMe)))
	http.Handle("/user",middlewares.ProtectedHandler(http.HandlerFunc(api.SingleUser)))
}