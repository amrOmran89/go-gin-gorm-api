package models

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
