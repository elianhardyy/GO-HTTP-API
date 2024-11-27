package routes

import (
	"net/http"
	"server/controllers"
	"server/middlewares"
)

func UserRoutes(api controllers.UserController) {
	R.HandleFunc("/create",api.Register).Methods("POST")
	R.HandleFunc("/login",api.Login).Methods("POST")
	R.HandleFunc("/verify",api.VerifyUser).Methods("POST")
	R.Handle("/userme",middlewares.ProtectedHandler(http.HandlerFunc(api.UserMe))).Methods("GET")
	R.Handle("/user",middlewares.ProtectedHandler(http.HandlerFunc(api.SingleUser))).Methods("GET")
	R.Handle("/profile",middlewares.ProtectedHandler(http.HandlerFunc(api.UploadProfile))).Methods("POST")
	R.Handle("/logout",middlewares.ProtectedHandler(http.HandlerFunc(api.Logout))).Methods("POST")
}