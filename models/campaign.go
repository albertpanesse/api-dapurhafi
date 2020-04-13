package models

import (
	"github.com/jinzhu/gorm"
)

type Campaign struct {
	gorm.Model
	Name						string				`gorm:"type:varchar(128); not null"`
	Description					string				`gorm:"type:varchar(512); not null"`
}
