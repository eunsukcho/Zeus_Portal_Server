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

type AuthHandler interface {
	//auth setting
	AuthInfoData(c *gin.Context)
	SaveAuthData(c *gin.Context)

	//Invitation User
	InvitationUser(c *gin.Context)
	CreateDevUser(c *gin.Context)

	GetDevUser(c *gin.Context)
	AcceptUser(c *gin.Context)
	DeleteTmpUser(c *gin.Context)
	CkDuplicateTmpDev(c *gin.Context)
}

func (h *Handler) AuthInfoData(c *gin.Context) {
	auth_detail_tbls, err := h.db.GetAllAuthData()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(auth_detail_tbls))
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   auth_detail_tbls,
		"len":    len(auth_detail_tbls),
	})
}
func (h *Handler) SaveAuthData(c *gin.Context) {
	var authDetails models.Authdetails
	err := c.ShouldBindJSON(&authDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	auth_detail_tbls, err := h.db.SaveAuthData(authDetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(auth_detail_tbls))
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   auth_detail_tbls,
		"len":    len(auth_detail_tbls),
	})
}

func (h *Handler) CreateDevUser(c *gin.Context) {
	var dev models.Dev_Info
	err := c.ShouldBindJSON(&dev)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	devInfo, err := h.db.SaveDevUserInfo(dev)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(devInfo)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   devInfo,
	})
}
func (h *Handler) GetDevUser(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("Group Name : ", uri.Id)

	devuser_info_tbls, err := h.db.GetDevUserInfo(uri.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(devuser_info_tbls))
	c.JSON(http.StatusOK, devuser_info_tbls)
}
func (h *Handler) AcceptUser(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.ReqId)
	accept_dev, err := h.db.AcceptUpdateUser(uri.ReqId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Found %d products", accept_dev)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
	})
}

func (h *Handler) DeleteTmpUser(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)
	err := h.db.DeleteUser(uri.Id, uri.ReqId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
	})
}

func (h *Handler) CkDuplicateTmpDev(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)
	accept_dev, err := h.db.CkDuplicateTmpDev(uri.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"body":   accept_dev,
	})
}

func (h *Handler) InvitationUser(c *gin.Context) {
	var inviteInfo models.Invitation
	var err error

	c.BindJSON(&inviteInfo)
	accessAuth := inviteInfo.AccessAuth
	invitationAddress := inviteInfo.InvitationAddress

	fmt.Println(accessAuth, invitationAddress)

	var smtpinfo []models.SmtpInfo
	smtpinfo, err = h.db.SmtpInfoGet()
	if err == nil {
		fmt.Println(smtpinfo)
	}
	if sendInvitataionEmail(accessAuth, invitationAddress, smtpinfo) != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func sendInvitataionEmail(accessAuth string, invitationAddress string, smtpinfo []models.SmtpInfo) error {
	port, _ := strconv.Atoi(smtpinfo[0].Port)
	d := gomail.NewDialer(smtpinfo[0].SmtpAddress, port, smtpinfo[0].AdminAddress, smtpinfo[0].Password)
	s, err := d.Dial()
	if err != nil {
		log.Println(err.Error(), smtpinfo)
		return err
	}

	userRegisterLink := "http://112.217.226.91:4207/user/invitation?accessAuth=" + accessAuth + "&email=" + invitationAddress

	fmt.Println("SMTP Info : ", smtpinfo)
	m := gomail.NewMessage()
	m.SetHeader("From", smtpinfo[0].AdminAddress)
	m.SetAddressHeader("To", invitationAddress, invitationAddress)
	m.SetHeader("Subject", "개발자 등록 요청")
	m.SetBody("text/html", fmt.Sprintf("\n아래의 링크를 복사해서 붙여넣으세요. %s", userRegisterLink))

	if err := gomail.Send(s, m); err != nil {
		return fmt.Errorf(
			"Could not send email to %q: %v", invitationAddress, err)
	}
	m.Reset()
	return nil
}
