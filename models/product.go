package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	CategoryID				uint				`gorm:"type:bigint; not null"`
	Name							string			`gorm:"type:varchar(128); not null"`
  Price							uint 				`gorm:"type:bigint; not null; default: 0"`	
	Desc							string			`gorm:"type:varchar(512); not null"`
	Tags							string			`gorm:"type:varchar(255)"`
  MinOrder					uint 				`gorm:"type:bigint; not null; default: 0"`	
  Weight						uint 				`gorm:"type:bigint; not null; default: 0"`	
  WeightUnit				string 			`gorm:"type:varchar(16); not null; default: 'kg'"`	
	OrderCount				uint				`gorm:"type:bigint; not null; default: 0"`
	ProductPicts			[]ProductPict
	Category  				Category
}
