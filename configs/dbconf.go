package configs

import (
	"fmt"
	"io"
	mdl "api-dapurhafi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"crypto/md5"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dapurhafi dbname=dapurhafi password=dapurhafi")
	if err != nil {
		panic("failed to connect to database")
	}

	db.
		Exec("DROP TABLE users").
		Exec("DROP TABLE menupics").
		Exec("DROP TABLE menuprices").
		Exec("DROP TABLE campaigns").
		Exec("DROP TABLE menus")

	db.CreateTable(&mdl.User{})
	db.CreateTable(&mdl.Campaign{})
	db.CreateTable(&mdl.Menu{})
	db.CreateTable(&mdl.Menuprice{}).
			AddForeignKey("campaign_id", "campaigns(id)", "CASCADE", "CASCADE").
			AddForeignKey("menu_id", "menus(id)", "CASCADE", "CASCADE")
	db.CreateTable(&mdl.Menupic{}).
			AddForeignKey("menu_id", "menus(id)", "CASCADE", "CASCADE")

	h := md5.New()

	io.WriteString(h, "123123")
	db.Create(&mdl.User{Fullname: "Albert Panesse", Email: "albert.panesse@gmail.com", Username: "albert.panesse@gmail.com", Password: fmt.Sprintf("%x", h.Sum(nil))})

	var campaignId uint

	db.Callback().Create().After("get_new_id").Register("get_new_id", func(scope *gorm.Scope) {
		campaignId = scope.PrimaryKeyValue().(uint)
	})

	db.Create(&mdl.Campaign{Name: "Regular", Description: "Regular campaign"})

	db.Callback().Create().Remove("get_new_id")

	var menuId uint

	db.Callback().Create().After("get_new_id").Register("get_new_id", func(scope *gorm.Scope) {
		menuId = scope.PrimaryKeyValue().(uint)
	})

	db.Create(&mdl.Menu{Name: "Nasi Goreng", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. \n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.", Tags: "nasi, nasi goreng"})
	db.Create(&mdl.Menupic{MenuID: menuId, Filename: "c37ab8d0072937659975b874dafe04cb.jpg"})
	db.Create(&mdl.Menuprice{CampaignID: campaignId, MenuID: menuId, Price: 10000})

	db.Create(&mdl.Menu{Name: "Rendang Daging Sapi", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. \n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.", Tags: "daging, rendang, sapi"})
	db.Create(&mdl.Menupic{MenuID: menuId, Filename: "1b70abfe370b98393d3e08fcf97891cc.jpg"})
	db.Create(&mdl.Menuprice{CampaignID: campaignId, MenuID: menuId, Price: 25000})

	db.Create(&mdl.Menu{Name: "Ayam Balado", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. \n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.", Tags: "daging, balado, ayam"})
	db.Create(&mdl.Menupic{MenuID: menuId, Filename: "6533e11e314b63821e1f569899d8f331.jpg"})
	db.Create(&mdl.Menuprice{CampaignID: campaignId, MenuID: menuId, Price: 20000})

	db.Callback().Create().Remove("get_new_id")

	return db
}