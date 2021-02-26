package smtpconnect

import (
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func Smtptest(c *gin.Context) {
	var smtpinfo SmtpInfo
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
}
