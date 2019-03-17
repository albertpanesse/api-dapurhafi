package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Fullname		string
	Email				string
	Paswd				string
}
