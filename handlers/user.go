package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) GetUsers(c *gin.Context) {
	var (
		users []mdl.User
		result gin.H
	)
	
	dbconn.DB.Find(&users)

	result = gin.H{
		"success": true,
		"result": users,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetUser(c *gin.Context) {
	var (
		user mdl.User
		result gin.H
	)	
	
	id := c.Param("id")
	
	dbconn.DB.First(&user, id)

	result = gin.H{
		"success": true,
		"result": user,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) CreateUser(c *gin.Context) {
	var (
		user mdl.User
		result gin.H
	)
	
	c.BindJSON(&user)
	
	dbconn.DB.Create(&user)

	result = gin.H{
		"success": true,
		"result": user,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) UpdateUser(c *gin.Context) {
	var (
		oldUser, newUser mdl.User
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&oldUser, id)
	
	c.BindJSON(&newUser)	
	
	dbconn.DB.Model(&oldUser).Updates(&newUser)

	result = gin.H{
		"success": true,
		"result": oldUser,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) DeleteUser(c *gin.Context) {
	var (
		user mdl.User
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&user, id)
	dbconn.DB.Delete(&user)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetRetailer(c *gin.Context) {
	var (
		user mdl.User
		result gin.H
	)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}