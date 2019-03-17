package main

import (
	"log"
	"github.com/gin-gonic/gin"
	hnd "api-dapurhafi/handlers"
	cfg "api-dapurhafi/configs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := cfg.DBInit()
	
	dbconn := &hnd.DBConn{DB: db}

	router := gin.Default()

	router.GET("/users", dbconn.GetUsers)
	router.GET("/user/:id", dbconn.GetUser)
	router.POST("/user", dbconn.CreateUser)
	router.PUT("/user/:id", dbconn.UpdateUser)
	router.DELETE("/user/:id", dbconn.DeleteUser)
	
	log.Fatal(router.Run(":3000"))
}