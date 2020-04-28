package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) GetMenus(c *gin.Context) {
	var (
		menus []mdl.Menu
		result gin.H
	)

	dbconn.DB.Preload("Menupics").Preload("Menuprice").Find(&menus)

	result = gin.H{
		"success": true,
		"result": menus,
	}

	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetMenu(c *gin.Context) {
	var (
		menu mdl.Menu
		result gin.H
	)

	id := c.Param("id")

	dbconn.DB.Preload("Menupics").Preload("Menuprice").First(&menu, id)

	result = gin.H{
		"success": true,
		"result": menu,
	}

	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) CreateMenu(c *gin.Context) {
	var (
		menu mdl.Menu
		result gin.H
	)

	c.BindJSON(&menu)

	dbconn.DB.Create(&menu)

	result = gin.H{
		"success": true,
		"result": menu,
	}

	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetFavs(c *gin.Context) {
	var (
		menus []mdl.Menu
		result gin.H
	)

	dbconn.DB.Preload("Menupics").Limit(5).Order("order_count desc").Find(&menus)

	result = gin.H{
		"success": true,
		"result": menus,
	}

	c.JSON(http.StatusOK, result)
}
