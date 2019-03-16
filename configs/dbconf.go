package configs

import (
	mdl "api-dapurhafi/models"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "dapurhafi:@tcp(127.0.0.1:3306)/dapurhafi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(mdl.User{})
	
	return db
}