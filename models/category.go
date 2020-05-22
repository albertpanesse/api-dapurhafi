package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name						string				`gorm:"type:varchar(128); not null"`
	Description			string				`gorm:"type:varchar(512); not null"`
	IsActive				bool					`form:"type:bool; not null; default true`
}
