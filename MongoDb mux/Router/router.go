package router

import (
	controller "go-rest-api-db/Controller"
	service "go-rest-api-db/Service"

	"github.com/gorilla/mux"
)

var (
	s service.NetflixService       = service.NewNetflixService()
	c controller.NetflixController = controller.NewController(s)
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", c.GetAll).Methods("GET")
	router.HandleFunc("/api/movie/{id}", c.Get).Methods("GET")
	router.HandleFunc("/api/create", c.Create).Methods("POST")
	router.HandleFunc("/api/update/{id}", c.Update).Methods("POST")
	router.HandleFunc("/api/delete/{id}", c.Delete).Methods("POST")
	router.HandleFunc("/api/deleteall", c.DeleteAll).Methods("POST")

	return router
}
