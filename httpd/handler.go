package httpd

import (
	_ "bytes"
	"fmt"
	_ "html/template"
	"net/http"
	"zeus/dblayer"
	"zeus/druid"
	"zeus/models"

	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/bcrypt"
)

type HandlerInterface interface {
	DBConnectionCheck(c *gin.Context)

	//env setting
	GetEnvData(c *gin.Context)
	UpdateEnvData(c *gin.Context)

	MenuHandler
	SmtpHandler
	AuthHandler

	DruidHandler

	NamespaceHandler
}

type Handler struct {
	db    dblayer.DBLayer
	druid druid.DruidInterface
}

func NewHandlerWithParams() (HandlerInterface, error) {
	db, err := dblayer.NewDBInit()
	druid := druid.NewClientInfo()

	if err != nil {
		return nil, err
	}
	return &Handler{
		db:    db,
		druid: druid,
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
