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
		oldUser, newUser mdl.User
		result gin.H
	)
	
	c.BindJSON(&newUser)
	
	if err := dbconn.DB.Where("email = ?", newUser.Email).First(&oldUser); err != nil {
		verificationToken, _ := util.RandomString(32)
		verificationExpired := time.Now()

		newUser.VerificationToken = verificationToken
		newUser.VerificationExpired = verificationExpired

		dbconn.DB.Create(&newUser)

		if &newUser != nil {
			plainText := "Silakan kopi dan paste URL berikut ini di browser Anda: http://dapurhafi.com/verify/" + verificationToken
			
			htmlText := "<a href=\"https://dapurhafi.com/verify/" + verificationToken + "\" target=\"_blank\">Silakan klik disini untuk aktivasi akun Anda.</a>"
			
			subject := "You are one step away"

			status, code, body := util.SendEmail(newUser.Fullname, newUser.Email, os.Getenv("DH_MAIL_NAME"), os.Getenv("DH_MAIL_EMAIL"), subject, plainText, htmlText)
			if status == true {			
				result = gin.H{
					"success": true,
					"result": newUser.Email,
				}
			} else {
				result = gin.H{
					"success": false,
					"error_type": "SENDGRID_ERROR",
					"error": "Code: " + string(code) + " [" + body + "]",
				}			
			}
		} else {
			result = gin.H{
				"success": false,
				"error_type": "DATABASE_ERROR",
				"error": "Creating new user has failed",
			}
		}
	} else {
		result = gin.H{
			"success": false,
			"error_type": "USER_EXISTS",
			"error_msg": "User already exists",
		}
	}

	c.JSON(http.StatusOK, result)	
}

func (dbconn *DBConn) Verify(c *gin.Context) {
	var (
		oldUser, newUser mdl.User
		result gin.H
	)
	
	token := c.Param("token")

	dbconn.DB.Where("verificationToken = '?'", token).First(&oldUser)
		
	accessToken, _ := util.RandomString(32)
	accessExpired := time.Now()

	oldUser.AccessToken = accessToken
	oldUser.AccessExpired = accessExpired
	
	dbconn.DB.Model(&newUser).Updates(&oldUser)
	
	if &newUser != nil {
		plainText := "Selamat Datang!"
		
		htmlText := "Selamat Datang!"
		
		subject := "Welcome, to Dapur Hafi"

		status, code, body := util.SendEmail(newUser.Fullname, newUser.Email, os.Getenv("DH_MAIL_NAME"), os.Getenv("DH_MAIL_EMAIL"), subject, plainText, htmlText)
		if status == true {			
			result = gin.H{
				"success": true,
				"result": newUser.AccessToken,
			}
		} else {
			result = gin.H{
				"success": false,
				"error_type": "SENDGRID_ERROR",
				"error": "Code: " + string(code) + " [" + body + "]",
			}			
		}
	} else {
		result = gin.H{
			"success": false,
			"error_type": "INVALID_TOKEN",
			"error_msg": "Invalid token",
		}
	}

	c.JSON(http.StatusOK, result)	
}
