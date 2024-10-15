package main

import (
	"log"
	"net/http"
	"server/config"
	"server/routes"

	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error to load env")
	}
	config.DBConnection()
	routes.InitRoute()
	log.Println("http://localhost:7000")
	log.Fatal(http.ListenAndServe(":7000",nil))
}