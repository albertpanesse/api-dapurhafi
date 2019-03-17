package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Fullname						string				`gorm:"type:varchar(128); not null"`
	Email								string				`gorm:"type:varchar(255); not null"`
	Paswd								string				`gorm:"type:varchar(64); not null"`
	VerificationToken		string				`gorm:"type:varchar(255); not null"`
	VerificationExpired	time.Time			`sql:"DEFAULT:current_timestamp"`
	ForgotToken					string				`gorm:"type:varchar(255)"`
	ForgotExpired				time.Time			`sql:"DEFAULT:current_timestamp"`	
	AccessToken					string				`gorm:"type:varchar(255)"`
	AccessExpired				time.Time			`sql:"DEFAULT:current_timestamp"`
	ProfileImage				string				`gorm:"type:varchar(128)"`
	FacebookId					string				`gorm:"type:varchar(16)"`
	GoogleId						string				`gorm:"type:varchar(16)"`
}
