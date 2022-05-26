package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/physX23/go-cityapp/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterCityRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", r))
}
