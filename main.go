package main

import (
	"os"
	"log"
	"time"
	"github.com/gin-gonic/gin"
	hnd "api-dapurhafi/handlers"
	cfg "api-dapurhafi/configs"
	"github.com/subosito/gotenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-contrib/cors"
)

func main() {
	gotenv.Load()
	
	db := cfg.DBInit()
	
	dbconn := &hnd.DBConn{DB: db}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://" + os.Getenv("DOMAIN_NAME")},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://dapurhafi.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// custom endpoint
	router.POST("/register", dbconn.Register)
	router.POST("/verify", dbconn.Verify)
	
	// CRUD endpoint
	router.GET("/users", dbconn.GetUsers)
	router.GET("/user/:id", dbconn.GetUser)
	router.POST("/user", dbconn.CreateUser)
	router.PUT("/user/:id", dbconn.UpdateUser)
	router.DELETE("/user/:id", dbconn.DeleteUser)
	
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}