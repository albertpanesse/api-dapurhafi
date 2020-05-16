package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) GetProducts(c *gin.Context) {
	var (
		products []mdl.Product
		result gin.H
	)

	dbconn.DB.Preload("ProductPicts").Preload("ProductPrice").Find(&products)

	result = gin.H{
		"success": true,
		"data": products,
	}

	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetProduct(c *gin.Context) {
	var (
		product mdl.Product
		result gin.H
	)

	id := c.Param("id")

	dbconn.DB.Preload("ProductPicts").Preload("ProductPrice").First(&product, id)

	result = gin.H{
		"success": true,
		"data": product,
	}

	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) CreateProduct(c *gin.Context) {
	var (
		product mdl.Product
		result gin.H
	)

	c.BindJSON(&product)

	dbconn.DB.Create(&product)

	result = gin.H{
		"success": true,
		"data": product,
	}

	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetFavorite(c *gin.Context) {
	var (
		products []mdl.Product
		result gin.H
	)

	dbconn.DB.Preload("ProductPicts").Limit(5).Order("order_count desc").Find(&products)

	result = gin.H{
		"success": true,
		"data": products,
	}

	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetLatest(c *gin.Context) {
	var (
		products []mdl.Product
		result gin.H
	)

	dbconn.DB.Preload("ProductPicts").Limit(5).Order("order_count desc").Find(&products)

	result = gin.H{
		"success": true,
		"data": products,
	}

	c.JSON(http.StatusOK, result)
}
