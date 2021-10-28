package model

import "github.com/jinzhu/gorm"

type UserType struct {
	gorm.Model

	Name string `gorm:"size:16;unique;not null;default:null"`
}

func (UserType) TableName() string {
	return "user_type"
}
