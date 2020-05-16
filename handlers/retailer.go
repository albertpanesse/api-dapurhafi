package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) GetRetailers(c *gin.Context) {
	var (
		retailers []mdl.Retailer
		result gin.H
	)
	
	dbconn.DB.Find(&retailers)

	result = gin.H{
		"success": true,
		"data": retailers,
	}
	
	c.JSON(http.StatusOK, result)
}
