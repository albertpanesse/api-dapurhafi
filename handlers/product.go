package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
	"encoding/json"
)

type SearchObj struct {
	Keyword    string `json:"keyword"`
	CategoryID int 		`json:"categoryId"`
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

func (dbconn *DBConn) UpdateProduct(c *gin.Context) {
	var (
		oldProduct, newProduct mdl.Product
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&oldProduct, id)
	
	c.BindJSON(&newProduct)	
	
	dbconn.DB.Model(&oldProduct).Updates(&newProduct)

	result = gin.H{
		"success": true,
		"data": oldProduct,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) DeleteProduct(c *gin.Context) {
	var (
		product mdl.Product
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&product, id)
	dbconn.DB.Delete(&product)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetProduct(c *gin.Context) {
	var (
		product mdl.Product
		result gin.H
	)	
	
	id := c.Param("id")
	
	dbconn.DB.Preload("ProductPicts").Preload("ProductPrice").Preload("Category").First(&product, id)

	result = gin.H{
		"success": true,
		"data": product,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetProducts(c *gin.Context) {
	var (
		products []mdl.Product
		result gin.H
	)
	
	dbconn.DB.Preload("ProductPicts").Preload("ProductPrice").Preload("Category").Find(&products)

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
	
	dbconn.DB.Preload("ProductPicts").Preload("ProductPrice").Preload("Category").Limit(10).Find(&products).Order("updated_at desc")

	result = gin.H{
		"success": true,
		"data": products,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) ProductSearch(c *gin.Context) {
	var (
		products []mdl.Product
		searchObj SearchObj
		result gin.H
	)

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&searchObj)
	if err != nil {
    panic(err)
  }

	var chain = dbconn.DB

	if (searchObj.Keyword != "") {
		chain = chain.Where("name LIKE ?", "%" + searchObj.Keyword + "%").Find(&products)
	}

	if (searchObj.CategoryID != -1) {
		chain = chain.Where(mdl.Product{CategoryID: uint(searchObj.CategoryID)}).Find(&products)
	}
	
	chain.Preload("ProductPicts").Preload("ProductPrice").Preload("Category").Find(&products)

	result = gin.H{
		"success": true,
		"data": products,
	}
	
	c.JSON(http.StatusOK, result)
}