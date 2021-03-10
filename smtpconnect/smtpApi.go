package smtpconnect

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"

	model "zeus/initModel"
)

func (SmtpInfo) TableName() string {
	return "smtp_setting_tbl"
}

func GetSmtpInfo() *[]SmtpInfo {
	var smtpinfo []SmtpInfo

	db := model.DbInit()
	defer db.Close()
	db.Find(&smtpinfo)

	return &smtpinfo
}

func (smtpinfo *SmtpInfo) SmtpConnectionCheck() error {
	password := smtpinfo.Password
	port, _ := strconv.Atoi(smtpinfo.Port)
	d := gomail.NewDialer(smtpinfo.SmtpAddress, port, smtpinfo.AdminAddress, password)
	_, err := d.Dial()
	if err != nil {
		return err
	}
	return nil
}

func SmtpSave(c *gin.Context) {
	var smtpinfo SmtpInfo
	db := model.DbInit()
	password, _ := bcrypt.GenerateFromPassword([]byte(smtpinfo.Password), bcrypt.DefaultCost)
	defer db.Close()
	err := c.BindJSON(&smtpinfo)
	smtpinfo.Password = string(password)
	err = smtpinfo.SmtpConnectionCheck()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	db.Model(&smtpinfo).Update(&smtpinfo)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   smtpinfo,
	})
}

func Smtptest(c *gin.Context) {
	var smtpinfo SmtpInfo
	err := c.BindJSON(&smtpinfo)
	err = smtpinfo.SmtpConnectionCheck()
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

func GetSmtpInfoData(c *gin.Context) {
	smtpinfo := GetSmtpInfo()
	c.JSON(http.StatusOK, smtpinfo)
}
