package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/physX23/go-cityapp/pkg/models"
	"github.com/physX23/go-cityapp/pkg/utils"
)

var NewCity models.City

func GetCity(w http.ResponseWriter, r *http.Request) {
	newCities := models.GetAllCity()
	res, _ := json.Marshal(newCities)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityId := vars["id"]
	ID, err := strconv.ParseInt(cityId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	cityDetails, _ := models.GetCityById(ID)
	res, _ := json.Marshal(cityDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCity(w http.ResponseWriter, r *http.Request) {
	CreateCity := &models.City{}
	utils.ParseBody(r, CreateCity)
	b := CreateCity.CreateCity()
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cityId := vars["id"]
	ID, err := strconv.ParseInt(cityId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	city := models.DeleteCity(ID)

	res, _ := json.Marshal(city)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCity(w http.ResponseWriter, r *http.Request) {
	var updateCity = &models.City{}
	utils.ParseBody(r, updateCity)
	vars := mux.Vars(r)
	cityId := vars["id"]
	ID, err := strconv.ParseInt(cityId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	cityDetails, db := models.GetCityById(ID)
	if updateCity.City != "" {
		cityDetails.City = updateCity.City
	}
	db.Save(&cityDetails)
	res, _ := json.Marshal(cityDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
