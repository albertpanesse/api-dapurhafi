package models

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model
	Name						string				`gorm:"type:varchar(255); not null"`
	Price						int					`gorm:"type:bigint; not null; default: 0"`
	Tags						string				`gorm:"type:varchar(255)"`
	Menupics					[]Menupic
}
