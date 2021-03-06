package main

import (
	"fmt"
	router "go-rest-api-db/Router"
	service "go-rest-api-db/Service"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	service.LoadAppConfig()
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", service.AppConfig.Port), r))
	fmt.Println("Listening at port 8000 ...")
}
