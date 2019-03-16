package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id					int				`json:"id,omitempty"`
	Fullname		string		`json:"fullname,omitempty"`
	Email				string		`json:"email,omitempty"`
	Paswd				string		`json:"paswd,omitempty"`
}
