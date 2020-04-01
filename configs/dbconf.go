package configs

import (
	mdl "api-dapurhafi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dapurhafi dbname=dapurhafi password=dapurhafi")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&mdl.User{})
	
	return db
}