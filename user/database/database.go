package database

import (
	model "github.com/gsdenys/ecomerce-microservice-example/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect(databaseName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&model.UserType{})
	db.AutoMigrate(&model.User{})

	return db
}
