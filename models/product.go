package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	UserID						uint					`gorm:"type:int; not null"`
	Name							string				`gorm:"type:varchar(128); not null"`
	Description				string				`gorm:"type:varchar(512); not null"`
	Tags							string				`gorm:"type:varchar(255)"`
	OrderCount				uint 					`gorm:"type:int; not null; default: 0"`
}
