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

	db.Exec("DROP TABLE users").Exec("DROP TABLE menupics").Exec("DROP TABLE menus")

	db.CreateTable(&mdl.User{})
	db.CreateTable(&mdl.Menu{})
	db.CreateTable(&mdl.Menupic{}).AddForeignKey("menu_id", "menus(id)", "CASCADE", "CASCADE")

	h := md5.New()

	io.WriteString(h, "123123")
	db.Create(&mdl.User{Fullname: "Albert Panesse", Email: "albert.panesse@gmail.com", Paswd: fmt.Sprintf("%x", h.Sum(nil))})

	var menuId uint

	db.Callback().Create().After("get_new_id").Register("get_new_id", func(scope *gorm.Scope) {
		menuId = scope.PrimaryKeyValue().(uint)
	})

	db.Create(&mdl.Menu{Name: "Nasi Goreng", Price: 10000, Tags: "nasi, nasi goreng"})
	db.Create(&mdl.Menupic{MenuID: menuId, Filename: "c37ab8d0072937659975b874dafe04cb.jpg"})

	db.Create(&mdl.Menu{Name: "Rendang Daging Sapi", Price: 25000, Tags: "daging, rendang, sapi"})
	db.Create(&mdl.Menupic{MenuID: menuId, Filename: "1b70abfe370b98393d3e08fcf97891cc.jpg"})

	db.Create(&mdl.Menu{Name: "Ayam Balado", Price: 20000, Tags: "daging, balado, ayam"})
	db.Create(&mdl.Menupic{MenuID: menuId, Filename: "6533e11e314b63821e1f569899d8f331.jpg"})

	db.Callback().Create().Remove("get_new_id")

	return db
}