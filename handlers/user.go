package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) CreateUser(c *gin.Context) {
	var (
		user mdl.User
		result gin.H
	)
	
	c.BindJSON(&user)
	
	dbconn.DB.Create(&user)
	result = gin.H{
		"result": user,
	}
	
	c.JSON(http.StatusOK, result)
}
