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
	GetEnvData(c *gin.Context)
	UpdateEnvData(c *gin.Context)

	Smtptest(c *gin.Context)
	SmtpSave(c *gin.Context)

	GetTopMenuData(c *gin.Context)
	SubTopMenuData(c *gin.Context)
	SaveTopMenu(c *gin.Context)
	SaveSubMenu(c *gin.Context)
	DeleteTopMenu(c *gin.Context)
	DeleteSubMenu(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
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

func NewHandlerWithDB(db dblayer.DBLayer) HandlerInterface {
	return &Handler{db: db}
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	smtp, err := h.db.SmtpInfoSave(smtpinfo, password)

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
