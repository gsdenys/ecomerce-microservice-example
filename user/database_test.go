package main

import (
	"testing"

	model "github.com/gsdenys/ecomerce-microservice-example/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestConnectDataBase(t *testing.T) {
	assert := assert.New(t)

	ConnectDataBase()
	DB.Create(&model.UserType{
		Name: "seller",
	})

	var ut []model.UserType
	DB.Find(&ut)

	assert.GreaterOrEqual(1, len(ut))

	DB.Create(&model.User{
		Name:  "John Smith",
		Email: "joao",
		Phone: "+5511987652145",
		Type:  ut[0],
	})

	var user []model.User
	DB.Find(&user)

	assert.GreaterOrEqual(1, len(ut))

}
