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

	router.POST("/user", dbconn.CreateUser)
	
	log.Fatal(router.Run(":3000"))
}