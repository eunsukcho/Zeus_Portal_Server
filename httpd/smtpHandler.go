package httpd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"zeus/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type SmtpHandler interface {
	//smtp setting
	Smtptest(c *gin.Context)
	SmtpSave(c *gin.Context)
	SmtpGet(c *gin.Context)
	SendMail(c *gin.Context)
}

//smtp setting
func SmtpConnectionCheck(smtpinfo *models.SmtpInfo) error {
	password := smtpinfo.Password
	port, _ := strconv.Atoi(smtpinfo.Port)
	d := gomail.NewDialer(smtpinfo.SmtpAddress, port, smtpinfo.AdminAddress, password)
	_, err := d.Dial()
	if err != nil {
		return err
	}
	return nil
}
func (h *Handler) Smtptest(c *gin.Context) {
	var smtpinfo models.SmtpInfo
	err := c.BindJSON(&smtpinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = SmtpConnectionCheck(&smtpinfo)
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

func (h *Handler) SmtpSave(c *gin.Context) {
	var smtpinfo models.SmtpInfo
	//password, _ := bcrypt.GenerateFromPassword([]byte(smtpinfo.Password), bcrypt.DefaultCost)

	err := c.BindJSON(&smtpinfo)
	err = SmtpConnectionCheck(&smtpinfo)

	//smtpinfo.Password = string(password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	smtp, err := h.db.SmtpInfoSave(smtpinfo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   smtp,
	})
}

func (h *Handler) SmtpGet(c *gin.Context) {

	smtpinfo, err := h.db.SmtpInfoGet()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, smtpinfo)
}

func (h *Handler) SendMail(c *gin.Context) {

	var smtpinfo models.SmtpInfo
	c.BindJSON(&smtpinfo)
	password := smtpinfo.Password

	port, _ := strconv.Atoi(smtpinfo.Port)
	d := gomail.NewDialer(smtpinfo.SmtpAddress, port, smtpinfo.AdminAddress, password)
	s, err := d.Dial()
	if err != nil {
		log.Println(err.Error(), smtpinfo)
		return
	}

	fmt.Println("SMTP Info : ", smtpinfo)
	m := gomail.NewMessage()
	m.SetHeader("From", smtpinfo.AdminAddress)
	m.SetAddressHeader("To", smtpinfo.AdminAddress, "test")
	m.SetHeader("Subject", "testtest")
	m.SetBody("text/html", fmt.Sprintf("Hello %s!", "test"))

	if err := gomail.Send(s, m); err != nil {
		fmt.Println("fail")
	}
	fmt.Println(smtpinfo.AdminAddress)
	m.Reset()
}
