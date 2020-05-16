package configs

import (
	"fmt"
	"io"
	mdl "api-dapurhafi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"crypto/md5"
)

func DBSeed(isActive bool, db *gorm.DB ) {
	if isActive {
		db.
			Exec("DROP TABLE product_picts").
			Exec("DROP TABLE product_prices").
			Exec("DROP TABLE products").
			Exec("DROP TABLE campaigns").
			Exec("DROP TABLE retailers").
			Exec("DROP TABLE users")

		db.CreateTable(&mdl.User{})
		db.CreateTable(&mdl.Retailer{}).
				AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
		db.CreateTable(&mdl.Campaign{}).
				AddForeignKey("retailer_id", "retailers(id)", "CASCADE", "CASCADE")
		db.CreateTable(&mdl.Product{}).
				AddForeignKey("retailer_id", "retailers(id)", "CASCADE", "CASCADE")
		db.CreateTable(&mdl.ProductPrice{}).
				AddForeignKey("campaign_id", "campaigns(id)", "CASCADE", "CASCADE").
				AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
		db.CreateTable(&mdl.ProductPict{}).
				AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")

		h := md5.New()

		// --- create new user, and get the ID
		var userId uint
		db.Callback().Create().After("get_new_id").Register("get_new_id", func(scope *gorm.Scope) {
			userId = scope.PrimaryKeyValue().(uint)
		})
		
		io.WriteString(h, "123123")
		db.Create(&mdl.User{Fullname: "Albert Panesse", Email: "albert.panesse@gmail.com", Mobile: "081226919868", Password: fmt.Sprintf("%x", h.Sum(nil))})
		
		io.WriteString(h, "123123")
		db.Create(&mdl.User{Fullname: "Johny Jontor", Email: "situs.ok@gmail.com", Mobile: "081393717674", Password: fmt.Sprintf("%x", h.Sum(nil)), UserType: "retailer"})
		
		db.Callback().Create().Remove("get_new_id")
		// ---

		// create new retailer, and get the ID
		var retailerId uint
		db.Callback().Create().After("get_new_id").Register("get_new_id", func(scope *gorm.Scope) {
			retailerId = scope.PrimaryKeyValue().(uint)
		})

		db.Create(&mdl.Retailer{UserID: userId, Name: "Jontor Online Shop", ProfileImage: "logo-jontor.jpg", IsActive: true})

		db.Callback().Create().Remove("get_new_id")
		// ---

		// create new campaign, and get the ID
		var campaignId uint
		db.Callback().Create().After("get_new_id").Register("get_new_id", func(scope *gorm.Scope) {
			campaignId = scope.PrimaryKeyValue().(uint)
		})

		db.Create(&mdl.Campaign{RetailerID: retailerId, Name: "Regular", Description: "Regular campaign", IsActive: true})

		db.Callback().Create().Remove("get_new_id")
		// ---

		// create new product, and get the ID
		var productId uint
		db.Callback().Create().After("get_new_id").Register("get_new_id", func(scope *gorm.Scope) {
			productId = scope.PrimaryKeyValue().(uint)
		})

		db.Create(&mdl.Product{RetailerID: retailerId, Name: "Nasi Goreng", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. \n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.", Tags: "nasi, nasi goreng"})
		db.Create(&mdl.ProductPict{ProductID: productId, Filename: "c37ab8d0072937659975b874dafe04cb.jpg"})
		db.Create(&mdl.ProductPrice{CampaignID: campaignId, ProductID: productId, Price: 10000})
		db.Create(&mdl.Product{RetailerID: retailerId, Name: "Rendang Daging Sapi", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. \n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.", Tags: "daging, rendang, sapi"})
		db.Create(&mdl.ProductPict{ProductID: productId, Filename: "1b70abfe370b98393d3e08fcf97891cc.jpg"})
		db.Create(&mdl.ProductPrice{CampaignID: campaignId, ProductID: productId, Price: 25000})
		db.Create(&mdl.Product{RetailerID: retailerId, Name: "Ayam Balado", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. \n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.", Tags: "daging, balado, ayam"})
		db.Create(&mdl.ProductPict{ProductID: productId, Filename: "6533e11e314b63821e1f569899d8f331.jpg"})
		db.Create(&mdl.ProductPrice{CampaignID: campaignId, ProductID: productId, Price: 20000})

		db.Callback().Create().Remove("get_new_id")
	}
}

func DBInit() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dapurhafi dbname=dapurhafi password=dapurhafi")
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}