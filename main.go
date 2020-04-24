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

	// auth endpoint
	router.POST("/register", dbconn.Register)
	router.POST("/verify", dbconn.Verify)

	// user endpoints
	router.GET("/users", dbconn.GetUsers)
	router.GET("/user/:id", dbconn.GetUser)
	router.POST("/user", dbconn.CreateUser)
	router.PUT("/user/:id", dbconn.UpdateUser)
	router.DELETE("/user/:id", dbconn.DeleteUser)

	// menu endpoints
	router.GET("/menus", dbconn.GetMenus)
	router.GET("/menu/:id", dbconn.GetMenu)
	router.POST("/menu", dbconn.CreateMenu)
	router.GET("/favs", dbconn.GetFavs)

	log.Fatal(router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
