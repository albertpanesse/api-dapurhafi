package main

import (
	"log"
	"github.com/gin-gonic/gin"
	hnd "api-dapurhafi/handlers"
	cfg "api-dapurhafi/configs"
)

func main() {
	db := cfg.DBInit()
	
	DBConn := &hnd.DBConn{DB: db}

	router := gin.Default()

	router.POST("/person", DBConn.CreatePerson)
	
	log.Fatal(router.Run(":3000"))
}