package models

import (
	"github.com/jinzhu/gorm"
)

type ProductPict struct {
	gorm.Model
	ProductID					uint				`gorm:"type:bigint; not null"`
	Filename					string				`gorm:"type:varchar(255); null"`
}
