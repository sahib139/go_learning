package routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sahib139/go_Basic_API_using_mongo/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controller.Greater).Methods("GET")
	router.HandleFunc("/api/create", controller.CreateOne).Methods("POST")
	router.HandleFunc("/api/get", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/get/{id}", controller.GetById).Methods("GET")
	router.HandleFunc("/api/delete/{id}", controller.DeleteById).Methods("DELETE")
	router.HandleFunc("/api/deleteAll", controller.DeleteAll).Methods("DELETE")
	router.HandleFunc("/api/update/{id}", controller.UpdateById).Methods("PUT")
	router.HandleFunc("/api/isWatched/{id}", controller.IsWatched).Methods("PUT")
	return router
}
