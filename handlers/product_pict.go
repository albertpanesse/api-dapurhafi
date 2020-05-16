package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) CreateProductPict(c *gin.Context) {
	var (
		productPict mdl.ProductPict
		result gin.H
	)
	
	c.BindJSON(&productPict)
	
	dbconn.DB.Create(&productPict)

	result = gin.H{
		"success": true,
		"data": productPict,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) UpdateProductPict(c *gin.Context) {
	var (
		oldProductPict, newProductPict mdl.ProductPict
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&oldProductPict, id)
	
	c.BindJSON(&newProductPict)	
	
	dbconn.DB.Model(&oldProductPict).Updates(&newProductPict)

	result = gin.H{
		"success": true,
		"data": oldProductPict,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) DeleteProductPict(c *gin.Context) {
	var (
		productPict mdl.ProductPict
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&productPict, id)
	dbconn.DB.Delete(&productPict)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetProductPict(c *gin.Context) {
	var (
		productPict mdl.ProductPict
		result gin.H
	)	
	
	id := c.Param("id")
	
	dbconn.DB.First(&productPict, id)

	result = gin.H{
		"success": true,
		"data": productPict,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetProductPicts(c *gin.Context) {
	var (
		productPicts []mdl.ProductPict
		result gin.H
	)
	
	dbconn.DB.Find(&productPicts)

	result = gin.H{
		"success": true,
		"data": productPicts,
	}
	
	c.JSON(http.StatusOK, result)
}
