package main

import (
	"os"
	"log"
	"time"
	"github.com/gin-gonic/gin"
	hnd "api-dapurhafi/handlers"
	cfg "api-dapurhafi/configs"
	"github.com/subosito/gotenv"
	"github.com/gin-contrib/cors"
)

func main() {
	gotenv.Load()

	db := cfg.DBInit()
	cfg.DBSeed(os.Getenv("SEEDING_ACTIVE") == "true", db)

	dbconn := &hnd.DBConn{DB: db}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://" + os.Getenv("DOMAIN_NAME")},
		AllowMethods: []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://" + os.Getenv("DOMAIN_NAME")
		},
		MaxAge: 12 * time.Hour,
	}))

	// serve image
	router.GET("/image/:filename", hnd.GetImage)
	router.POST("/image", hnd.UploadImage)

	// custom endpoints
	router.GET("/latest", dbconn.GetLatest)
	router.POST("/product-search", dbconn.ProductSearch)

	// category endpoints
	router.POST("/category", dbconn.CreateCategory)
	router.PUT("/category/:id", dbconn.UpdateCategory)
	router.DELETE("/category/:id", dbconn.DeleteCategory)
	router.GET("/category/:id", dbconn.GetCategory)
	router.GET("/categories", dbconn.GetCategories)

	// product endpoints
	router.POST("/product", dbconn.CreateProduct)
	router.PUT("/product/:id", dbconn.UpdateProduct)
	router.DELETE("/product/:id", dbconn.DeleteProduct)
	router.GET("/product/:id", dbconn.GetProduct)
	router.GET("/products", dbconn.GetProducts)
	
	// product_pict endpoints
	router.POST("/product-pict", dbconn.CreateProductPict)
	router.PUT("/product-pict/:id", dbconn.UpdateProductPict)
	router.DELETE("/product-pict/:id", dbconn.DeleteProductPict)
	router.GET("/product-pict/:id", dbconn.GetProductPict)
	router.GET("/product-picts", dbconn.GetProductPicts)
	
	log.Fatal(router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
