package models

import (
	"github.com/jinzhu/gorm"
)

type Campaign struct {
	gorm.Model
	RetailerID			uint				`gorm:"type:bigint; not null"`
	Name						string				`gorm:"type:varchar(128); not null"`
	Description			string				`gorm:"type:varchar(512); not null"`
	IsActive				bool					`form:"type:bool; not null; default false`
}
