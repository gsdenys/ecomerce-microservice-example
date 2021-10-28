package main

import (
	model "github.com/gsdenys/ecomerce-microservice-example/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&model.UserType{})
	database.AutoMigrate(&model.User{})

	DB = database
}
