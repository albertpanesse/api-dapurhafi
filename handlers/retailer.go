package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) CreateRetailer(c *gin.Context) {
	var (
		retailer mdl.Retailer
		result gin.H
	)
	
	c.BindJSON(&retailer)
	
	dbconn.DB.Create(&retailer)

	result = gin.H{
		"success": true,
		"data": retailer,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) UpdateRetailer(c *gin.Context) {
	var (
		oldRetailer, newRetailer mdl.Retailer
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&oldRetailer, id)
	
	c.BindJSON(&newRetailer)	
	
	dbconn.DB.Model(&oldRetailer).Updates(&newRetailer)

	result = gin.H{
		"success": true,
		"data": oldRetailer,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) DeleteRetailer(c *gin.Context) {
	var (
		retailer mdl.Retailer
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&retailer, id)
	dbconn.DB.Delete(&retailer)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetRetailer(c *gin.Context) {
	var (
		retailer mdl.Retailer
		result gin.H
	)	
	
	id := c.Param("id")
	
	dbconn.DB.First(&retailer, id)

	result = gin.H{
		"success": true,
		"data": retailer,
	}
	
	c.JSON(http.StatusOK, result)
}

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
