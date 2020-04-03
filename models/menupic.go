package models

import (
	"github.com/jinzhu/gorm"
)

type Menupic struct {
	gorm.Model
	MenuID						uint				`gorm:"type:int; not null"`
	Filename					string				`gorm:"type:varchar(255); null"`
}
