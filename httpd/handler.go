package httpd

import (
	"fmt"
	"net/http"
	"strconv"
	"zeus/dblayer"
	"zeus/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type HandlerInterface interface {
	//env setting
	GetEnvData(c *gin.Context)
	UpdateEnvData(c *gin.Context)

	//smtp setting
	Smtptest(c *gin.Context)
	SmtpSave(c *gin.Context)
	SmtpGet(c *gin.Context)

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

// init zeus env
func (h *Handler) GetEnvData(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}

	env_setting_tbls, err := h.db.GetAllEnvData()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(env_setting_tbls))
	c.JSON(http.StatusOK, env_setting_tbls)
}
func (h *Handler) UpdateEnvData(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
	var env models.Envs
	err := c.BindJSON(&env)

	rst, err := h.db.UpdateEnvData(env)

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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
	var smtpinfo models.SmtpInfo
	password, _ := bcrypt.GenerateFromPassword([]byte(smtpinfo.Password), bcrypt.DefaultCost)

	err := c.BindJSON(&smtpinfo)
	err = SmtpConnectionCheck(&smtpinfo)
	smtpinfo.Password = string(password)

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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}

	smtpinfo, err := h.db.SmtpInfoGet()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, smtpinfo)
}

//menu setting
func (h *Handler) GetTopMenuData(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}

	top_menu, err := h.db.GetAllTopMenu()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(top_menu))
	c.JSON(http.StatusOK, top_menu)
}
func (h *Handler) SubTopMenuData(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}

	sub_menu, err := h.db.GetAllSubMenu()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(sub_menu))
	c.JSON(http.StatusOK, sub_menu)
}
func (h *Handler) SaveTopMenu(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
	icon, err := h.db.GetAllIcon()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, icon)
}

func (h *Handler) SaveUrlLink(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
	if h.db == nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Database error"})
		return
	}
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
