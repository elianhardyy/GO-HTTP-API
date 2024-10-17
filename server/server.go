package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"server/routes"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func Server(){
    err := godotenv.Load()
	if err != nil{
		log.Fatal("error to load env")
	}
	var store *sessions.CookieStore

	key := "mysecret"
	maxAge := 86400 * 30
	isProduction := false

	store = sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.HttpOnly = true
	store.Options.Secure = isProduction
	gothic.Store = store
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleCallback := "http://localhost:7000/auth/google/callback"
	
	goth.UseProviders(
		google.New(clientID,clientSecret,googleCallback,"email","profile"),

	)
	routes.R.HandleFunc("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		_, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
		  fmt.Fprintln(res, err)
		  return
		}
		
	  }).Methods("GET")
	
	routes.R.HandleFunc("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	}).Methods("GET")
	
	
	routes.InitRoute()
	log.Println("http://localhost:7000")
	log.Fatal(http.ListenAndServe(":7000",routes.R))
}