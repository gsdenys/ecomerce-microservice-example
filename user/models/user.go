package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Name  string   `gorm:"size:255;not null;default:null"`
	Email string   `gorm:"size:255;unique;not null;default:null"`
	Phone string   `gorm:"size:14;not null;default:null"`
	Type  UserType `gorm:"not null"`
}
