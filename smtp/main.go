package main

import (
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func CORSMiddelware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Cache-Control, Pragma, jsonType, ")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://localhost:4400")
		c.Header("Access-Control-Allow-Methods", "GET,POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type smtpInfo struct {
	AdminAddress   string `json:"AdminAddress" binding:"required"`
	SmtpAddress    string `json:"SmtpAddress" binding:"required"`
	Port           string `json:"Port" binding:"required"`
	Password       string `json:"Password" binding:"required"`
	Authentication string `json:"Authentication" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddelware())

	r.POST("/get/smtp", func(c *gin.Context) {
		var smtpinfo smtpInfo
		c.BindJSON(&smtpinfo)
		auth := smtp.PlainAuth("", smtpinfo.AdminAddress, smtpinfo.Password, smtpinfo.SmtpAddress)
		from := smtpinfo.AdminAddress
		to := []string{"dudco0355@naver.com"}
		msg := []byte("test")
		err := smtp.SendMail(smtpinfo.SmtpAddress+":"+smtpinfo.Port, auth, from, to, msg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"isOK":   0,
				"error":  err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"isOK":   1,
			"data":   smtpinfo,
		})
	})

	r.Run("127.0.0.1:3001")
}
