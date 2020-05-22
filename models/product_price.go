package models

import (
	"github.com/jinzhu/gorm"
)

type ProductPrice struct {
	gorm.Model
	CampaignID			uint				`gorm:"type:bigint; not null"`
	ProductID				uint				`gorm:"type:bigint; not null"`
	Price						uint				`gorm:"type:bigint; not null; default 0"`
}
