package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) CreateProductPrice(c *gin.Context) {
	var (
		productPrice mdl.ProductPrice
		result gin.H
	)
	
	c.BindJSON(&productPrice)
	
	dbconn.DB.Create(&productPrice)

	result = gin.H{
		"success": true,
		"data": productPrice,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) UpdateProductPrice(c *gin.Context) {
	var (
		oldProductPrice, newProductPrice mdl.ProductPrice
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&oldProductPrice, id)
	
	c.BindJSON(&newProductPrice)	
	
	dbconn.DB.Model(&oldProductPrice).Updates(&newProductPrice)

	result = gin.H{
		"success": true,
		"data": oldProductPrice,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) DeleteProductPrice(c *gin.Context) {
	var (
		productPrice mdl.ProductPrice
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&productPrice, id)
	dbconn.DB.Delete(&productPrice)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetProductPrice(c *gin.Context) {
	var (
		productPrice mdl.ProductPrice
		result gin.H
	)	
	
	id := c.Param("id")
	
	dbconn.DB.First(&productPrice, id)

	result = gin.H{
		"success": true,
		"data": productPrice,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetProductPrices(c *gin.Context) {
	var (
		productPrices []mdl.ProductPrice
		result gin.H
	)
	
	dbconn.DB.Find(&productPrices)

	result = gin.H{
		"success": true,
		"data": productPrices,
	}
	
	c.JSON(http.StatusOK, result)
}
