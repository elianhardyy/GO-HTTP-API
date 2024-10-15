package routes

import (
	"net/http"
	"server/controllers"
	"server/middlewares"

	"github.com/gorilla/mux"
)
var R = mux.NewRouter()
func UserRoutes(api controllers.UserController) {
	R.HandleFunc("/create",api.Register).Methods("POST")
	R.HandleFunc("/login",api.Login).Methods("POST")
	R.Handle("/userme",middlewares.ProtectedHandler(http.HandlerFunc(api.UserMe))).Methods("GET")
	R.Handle("/user",middlewares.ProtectedHandler(http.HandlerFunc(api.SingleUser))).Methods("GET")
}