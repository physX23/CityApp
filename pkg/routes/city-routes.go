package routes

import (
	"github.com/gorilla/mux"
	"github.com/physX23/go-cityapp/pkg/controllers"
)

var RegisterCityRoutes = func(router *mux.Router) {
	router.HandleFunc("/add", controllers.CreateCity).Methods("POST")
	router.HandleFunc("/add", controllers.GetCity).Methods("GET")
	router.HandleFunc("/add/{id}", controllers.GetCityById).Methods("GET")
	router.HandleFunc("/add/{id}", controllers.UpdateCity).Methods("PUT")
	router.HandleFunc("/add/{id}", controllers.DeleteCity).Methods("DELETE")
}
