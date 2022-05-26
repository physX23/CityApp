package models

import (
	"github.com/jinzhu/gorm"
	"github.com/physX23/go-cityapp/pkg/config"
)

var db *gorm.DB

type City struct {
	gorm.Model
	City string `json:"city"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&City{})
}

func (c *City) CreateCity() *City {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllCity() []City {
	var Cities []City
	db.Find(&Cities)
	return Cities
}

func GetCityById(Id int64) (*City, *gorm.DB) {
	var getCity City
	db := db.Where("ID=?", Id).Find(&getCity)
	return &getCity, db
}

func DeleteCity(Id int64) City {
	var city City
	db.Where("ID=?", Id).Delete(city)
	return city
}
