package httpd

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"zeus/dblayer"
	"zeus/models"

	"github.com/gin-gonic/gin"
	_"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type HandlerInterface interface {
	DBConnectionCheck (c *gin.Context)

	//env setting
	GetEnvData(c *gin.Context)
	UpdateEnvData(c *gin.Context)

	//smtp setting
	Smtptest(c *gin.Context)
	SmtpSave(c *gin.Context)
	SmtpGet(c *gin.Context)
	SendMail(c *gin.Context)

	//menu setting
	GetTopMenuData(c *gin.Context)
	SubTopMenuData(c *gin.Context)
	SaveTopMenu(c *gin.Context)
	SaveSubMenu(c *gin.Context)
	DeleteTopMenu(c *gin.Context)
	DeleteSubMenu(c *gin.Context)
	GetIcon(c *gin.Context)
	SaveUrlLink(c *gin.Context)
	SaveUrlSubLink(c *gin.Context)
	DeleteTopMenuUrl(c *gin.Context)
	DeleteSubMenuUrl(c *gin.Context)
	GetMenuTargetUrl(c *gin.Context)
	GetTopMenuTargetUrl(c *gin.Context)
	UpdateTopMenuInfo(c *gin.Context)
	UpdateSubMenuInfo(c *gin.Context)

	//auth setting
	AuthInfoData(c *gin.Context)
	SaveAuthData(c *gin.Context)

	//Invitation User
	InvitationUser(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
	fmt.Println("Handler")
	return NewHandlerWithParams()
}
func NewHandlerWithParams() (HandlerInterface, error) {
	db, err := dblayer.NewDBInit()
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) DBConnectionCheck(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
}

// init zeus env
func (h *Handler) GetEnvData(c *gin.Context) {
	env_setting_tbls, err := h.db.GetAllEnvData()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(env_setting_tbls))
	c.JSON(http.StatusOK, env_setting_tbls)
}

func (h *Handler) UpdateEnvData(c *gin.Context) {
	var env models.Env_setting_Tbls
	err := c.ShouldBindJSON(&env)
	fmt.Println("env : ", env)

	rst, err := h.db.UpdateEnvData(env)
	fmt.Println("rst : ", rst)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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

//menu setting
func (h *Handler) GetTopMenuData(c *gin.Context) {

	top_menu, err := h.db.GetAllTopMenu()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(top_menu))
	c.JSON(http.StatusOK, top_menu)
}
func (h *Handler) SubTopMenuData(c *gin.Context) {

	sub_menu, err := h.db.GetAllSubMenu()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(sub_menu))
	c.JSON(http.StatusOK, sub_menu)
}
func (h *Handler) SaveTopMenu(c *gin.Context) {
	var topMenu models.TopMenuInfo
	err := c.ShouldBindJSON(&topMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.SaveTopMenuInfo(topMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) SaveSubMenu(c *gin.Context) {
	var subMenu models.SubMenuInfo

	err := c.ShouldBindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.SaveSubMenuInfo(subMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) DeleteTopMenu(c *gin.Context) {
	var topMenu models.TopMenuInfo

	err := c.BindJSON(&topMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.DeleteTopMenuInfo(topMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) DeleteSubMenu(c *gin.Context) {
	var subMenu models.SubMenuInfo

	err := c.BindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.DeleteSubMenuInfo(subMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) GetIcon(c *gin.Context) {

	icon, err := h.db.GetAllIcon()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, icon)
}

func (h *Handler) SaveUrlLink(c *gin.Context) {
	
	var topMenu models.TopMenuInfo
	err := c.ShouldBindJSON(&topMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.SaveUrlLink(topMenu)
	fmt.Println(rst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})

}
func (h *Handler) SaveUrlSubLink(c *gin.Context) {
	var subMenu models.SubMenuInfo
	err := c.ShouldBindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.SaveUrlSubLink(subMenu)
	fmt.Println(rst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})

}

func (h *Handler) DeleteTopMenuUrl(c *gin.Context) {
	var topMenu models.TopMenuInfo
	err := c.BindJSON(&topMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}

	rst, err := h.db.DeleteTopMenuUrl(topMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}

func (h *Handler) DeleteSubMenuUrl(c *gin.Context) {
	var subMenu models.SubMenuInfo
	err := c.BindJSON(&subMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.DeleteSubMenuUrl(subMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) GetMenuTargetUrl(c *gin.Context) {
	var menuCode models.SubMenuInfo
	err := c.ShouldBindJSON(&menuCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	fmt.Println("MenuCode :", menuCode)

	rst, err := h.db.GetMenuTargetUrl(menuCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println("targetUrl : ", rst)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst.Sub_Menu_Target_Url,
	})

}
func (h *Handler) GetTopMenuTargetUrl(c *gin.Context) {
	var menuCode models.TopMenuInfo
	err := c.ShouldBindJSON(&menuCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	fmt.Println("MenuCode :", menuCode)

	rst, err := h.db.GetTopMenuTargetUrl(menuCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println("targetUrl : ", rst)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst.Top_Menu_Target_Url,
	})

}

func (h *Handler) UpdateTopMenuInfo(c *gin.Context) {
	var topMenu models.TopMenuInfo
	err := c.ShouldBindJSON(&topMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.UpdateTopMenuInfo(topMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})

}

func (h *Handler) UpdateSubMenuInfo(c *gin.Context) {
	var subMenu models.SubMenuInfo
	err := c.ShouldBindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.UpdateSubMenuInfo(subMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
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
func (h *Handler) InvitationUser(c *gin.Context) {
	var inviteInfo models.Invitation
	var err error

	c.BindJSON(&inviteInfo)
	accessAuth := inviteInfo.AccessAuth
	invitationAddress := inviteInfo.InvitationAddress

	fmt.Println(accessAuth, invitationAddress)
	
	var smtpinfo []models.SmtpInfo
	smtpinfo, err = h.db.SmtpInfoGet()
	if(err == nil) {
		fmt.Println(smtpinfo)
	}
	if sendInvitataionEmail(accessAuth, invitationAddress, smtpinfo) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":http.StatusBadRequest,
		})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusOK,
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

	userRegisterLink := "http://192.168.0.102:4201/user/invitation/"+accessAuth
	
	fmt.Println("SMTP Info : ", smtpinfo)
	m := gomail.NewMessage()
	m.SetHeader("From", smtpinfo[0].AdminAddress)
	m.SetAddressHeader("To", invitationAddress, invitationAddress)
	m.SetHeader("Subject", "개발자 등록 요청")
	m.SetBody("text/html", fmt.Sprintf("아래의 링크를 복사해서 붙여넣으세요. %s", userRegisterLink))

	if err := gomail.Send(s, m); err != nil {
		return fmt.Errorf(
			"Could not send email to %q: %v", invitationAddress, err)
	}
	m.Reset()
	return nil
}