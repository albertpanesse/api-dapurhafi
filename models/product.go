package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	CategoryID				uint				`gorm:"type:bigint; not null"`
	Name							string				`gorm:"type:varchar(128); not null"`
	Description				string				`gorm:"type:varchar(512); not null"`
	Tags							string				`gorm:"type:varchar(255)"`
	OrderCount				uint				`gorm:"type:bigint; not null; default: 0"`
	ProductPicts			[]ProductPict
	ProductPrice			ProductPrice
	Category  				Category
}
