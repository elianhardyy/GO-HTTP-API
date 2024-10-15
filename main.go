package main

import (
	"log"
	"net/http"
	"server/config"
	"server/routes"
)



func main() {
	config.DBConnection()
	routes.InitRoute()
	log.Println("http://localhost:7000")
	log.Fatal(http.ListenAndServe(":7000",nil))
}