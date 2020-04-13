package models

import (
	"github.com/jinzhu/gorm"
)

type Menuprice struct {
	gorm.Model
	CampaignID					uint				`gorm:"type:int; not null"`
	MenuID						uint				`gorm:"type:int; not null"`
	Price						uint				`gorm:"type:bigint; not null; default 0"`
}
