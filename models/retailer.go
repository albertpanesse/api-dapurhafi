package models

import (
	"github.com/jinzhu/gorm"
)

type Retailer struct {
	gorm.Model
	UserID					uint					`gorm:"type:bigint; not null"`
	Name						string				`gorm:"type:varchar(128); not null"`
	ProfileImage		string				`gorm:"type:varchar(128)"`
	IsActive				bool					`form:"type:bool; not null; default false`
}
