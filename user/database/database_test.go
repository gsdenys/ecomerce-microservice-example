package database

import (
	"fmt"
	"testing"
	"time"

	model "github.com/gsdenys/ecomerce-microservice-example/database/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestConnectDataBase(t *testing.T) {
	assert := assert.New(t)

	db := Connect("test")

	assert.NotNil(db)

	// ConnectDataBase()
	// DB.Create(&model.UserType{
	// 	Name: "seller",
	// })

	// var ut []model.UserType
	// DB.Find(&ut)

	// assert.GreaterOrEqual(1, len(ut))

	// DB.Create(&model.User{
	// 	Name:  "John Smith",
	// 	Email: "joao",
	// 	Phone: "+5511987652145",
	// 	Type:  ut[0],
	// })

	// var user []model.User
	// DB.Find(&user)

	// assert.GreaterOrEqual(1, len(ut))

}

func TestDataBase_1K(t *testing.T) {
	assert := assert.New(t)

	db := Connect("test")

	start := time.Now()
	for i := 0; i < 1000; i = i + 1 {
		db.Create(&model.UserType{
			Name: fmt.Sprintf("type_%d", i),
		})
	}
	elapsed := time.Since(start)

	assert.Equal(float64(1), elapsed.Seconds())

}
