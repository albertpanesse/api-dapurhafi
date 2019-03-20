package handlers

import (
	"os"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	mdl "api-dapurhafi/models"
	util "api-dapurhafi/utils"
)

func (dbconn *DBConn) Register(c *gin.Context) {
	var (
		user mdl.User
		result gin.H
	)
	
	c.BindJSON(&user)
	
	verificationToken, _ := util.RandomString(32)
	verificationExpired := time.Now()

	user.VerificationToken = verificationToken
	user.VerificationExpired = verificationExpired
	
	dbconn.DB.Create(&user)
	
	if &user != nil {
		plainText := "Silakan kopi dan paste URL berikut ini di browser Anda: http://dapurhafi.com/verify/" + verificationToken
		
		htmlText := "<a href=\"https://dapurhafi.com/verify/" + verificationToken + "\" target=\"_blank\">Silakan klik disini untuk aktivasi akun Anda.</a>"
		
		subject := "Welcome, to Dapur Hafi"

		status, code, body := util.SendEmail(user.Fullname, user.Email, os.Getenv("DH_MAIL_NAME"), os.Getenv("DH_MAIL_EMAIL"), subject, plainText, htmlText)
		if status == true {			
			result = gin.H{
				"success": true,
				"result": user,
			}
		} else {
			result = gin.H{
				"success": false,
				"error": "Code: " + string(code) + " [" + body + "]",
			}			
		}
	} else {
		result = gin.H{
			"success": false,
			"error": "Creating new user has failed",
		}
	}

	c.JSON(http.StatusOK, result)	
}
