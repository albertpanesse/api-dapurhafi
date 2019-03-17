package configs

import (
	mdl "api-dapurhafi/models"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "dapurhafi:dapurhafi@/dapurhafi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mdl.User{})
	
	return db
}