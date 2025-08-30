package models

import "gorm.io/gorm"

type PostEntity struct {
	gorm.Model
	Title string
	Body  string
}
