package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) CreateCategory(c *gin.Context) {
	var (
		category mdl.Category
		result gin.H
	)
	
	c.BindJSON(&category)
	
	dbconn.DB.Create(&category)

	result = gin.H{
		"success": true,
		"data": category,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) UpdateCategory(c *gin.Context) {
	var (
		oldCategory, newCategory mdl.Category
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&oldCategory, id)
	
	c.BindJSON(&newCategory)	
	
	dbconn.DB.Model(&oldCategory).Updates(&newCategory)

	result = gin.H{
		"success": true,
		"data": oldCategory,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) DeleteCategory(c *gin.Context) {
	var (
		category mdl.Category
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&category, id)
	dbconn.DB.Delete(&category)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetCategory(c *gin.Context) {
	var (
		category mdl.Category
		result gin.H
	)	
	
	id := c.Param("id")
	
	dbconn.DB.First(&category, id)

	result = gin.H{
		"success": true,
		"data": category,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetCategories(c *gin.Context) {
	var (
		categories []mdl.Category
		result gin.H
	)
	
	dbconn.DB.Find(&categories)

	result = gin.H{
		"success": true,
		"data": categories,
	}
	
	c.JSON(http.StatusOK, result)
}
