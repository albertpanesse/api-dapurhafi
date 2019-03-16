package handlers

import (
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
	cfg "api-dapurhafi/configs"
)

func (dbconn *cfg.DBConn) CreateUser(c *gin.Context) {
	var (
	person structs.Person
	result gin.H
	)
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name
	dbconn.DB.Create(&person)
	result = gin.H{
	"result": person,
	}
	c.JSON(http.StatusOK, result)
}
