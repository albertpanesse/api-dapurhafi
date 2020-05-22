package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
)

func (dbconn *DBConn) CreateCampaign(c *gin.Context) {
	var (
		campaign mdl.Campaign
		result gin.H
	)
	
	c.BindJSON(&campaign)
	
	dbconn.DB.Create(&campaign)

	result = gin.H{
		"success": true,
		"result": campaign,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) UpdateCampaign(c *gin.Context) {
	var (
		oldCampaign, newCampaign mdl.Campaign
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&oldCampaign, id)
	
	c.BindJSON(&newCampaign)	
	
	dbconn.DB.Model(&oldCampaign).Updates(&newCampaign)

	result = gin.H{
		"success": true,
		"result": oldCampaign,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) DeleteCampaign(c *gin.Context) {
	var (
		campaign mdl.Campaign
		result gin.H
	)
	
	id := c.Param("id")

	dbconn.DB.First(&campaign, id)
	dbconn.DB.Delete(&campaign)

	result = gin.H{
		"success": true,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetCampaign(c *gin.Context) {
	var (
		campaign mdl.Campaign
		result gin.H
	)	
	
	id := c.Param("id")
	
	dbconn.DB.First(&campaign, id)

	result = gin.H{
		"success": true,
		"result": campaign,
	}
	
	c.JSON(http.StatusOK, result)
}

func (dbconn *DBConn) GetCampaigns(c *gin.Context) {
	var (
		campaigns []mdl.Campaign
		result gin.H
	)
	
	dbconn.DB.Find(&campaigns)

	result = gin.H{
		"success": true,
		"result": campaigns,
	}
	
	c.JSON(http.StatusOK, result)
}
