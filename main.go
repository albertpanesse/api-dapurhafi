package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	hnd "api-dapurhafi/handlers"
	cfg "api-dapurhafi/configs"
	"github.com/subosito/gotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	gotenv.Load()
	
	db := cfg.DBInit()
	
	dbconn := &hnd.DBConn{DB: db}

	router := gin.Default()

	// custom endpoint
	router.POST("/register", dbconn.Register)
	
	// CRUD endpoint
	router.GET("/users", dbconn.GetUsers)
	router.GET("/user/:id", dbconn.GetUser)
	router.POST("/user", dbconn.CreateUser)
	router.PUT("/user/:id", dbconn.UpdateUser)
	router.DELETE("/user/:id", dbconn.DeleteUser)
	
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}