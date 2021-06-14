package httpd

import (
	"bytes"
	_ "bytes"
	"fmt"
	_ "html/template"
	"io/ioutil"
	"net/http"
	"time"
	"zeus/dblayer"
	"zeus/druid"
	"zeus/k8s"
	"zeus/models"

	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/bcrypt"
)

type HandlerInterface interface {
	DBConnectionCheck(c *gin.Context)

	//env setting
	GetEnvData(c *gin.Context)
	GetEnvDataShare() ([]models.Env_setting_Tbls, error)
	UpdateEnvData(c *gin.Context)

	MenuHandler
	SmtpHandler
	AuthHandler

	DruidHandler

	K8SNamespaceInterface
	DevServerHandler
}

type Handler struct {
	db    dblayer.DBLayer
	druid druid.DruidInterface
	k8s   k8s.K8SInterface
}

func NewHandlerWithParams() (HandlerInterface, error) {
	db, err := dblayer.NewDBInit()
	druid := druid.NewClientInfo()
	k8s := k8s.SettingK8SInfo("ope")

	if err != nil {
		return nil, err
	}

	return &Handler{
		db:    db,
		druid: druid,
		k8s:   k8s,
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
	env_setting_tbls, err := h.GetEnvDataShare()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(env_setting_tbls))
	c.JSON(http.StatusOK, env_setting_tbls)
}

func (h *Handler) GetEnvDataShare() ([]models.Env_setting_Tbls, error) {
	env_setting_tbls, err := h.db.GetAllEnvData()
	if err != nil {
		return nil, err
	}
	return env_setting_tbls, nil
}

func (h *Handler) UpdateEnvData(c *gin.Context) {
	var env models.Env_setting_Tbls
	err := c.ShouldBindJSON(&env)
	fmt.Println("env : ", env)

	env_setting_tbls, err := h.GetEnvDataShare()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rst, err := h.db.UpdateEnvData(env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updateGrafanaPreference string
	if env.ThemeSettingVal == "LIGHT" {
		updateGrafanaPreference = `{"theme" : "light" }`
	} else {
		updateGrafanaPreference = `{"theme" : "dark" }`
	}

	var reqJson = bytes.NewBuffer([]byte(updateGrafanaPreference))
	respVal, _, err := HTTPGetGrafana("PUT", "http://112.217.226.91:32001/api/org/preferences", reqJson, env.GrafanaToken)
	if err != nil {
		_, err := h.db.UpdateEnvData(env_setting_tbls[0])

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("RespVal : ", string(respVal))

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}

func HTTPGetGrafana(method string, url string, reqJson *bytes.Buffer, token string) (respBody []byte, statusCode int, err error) {
	var req *http.Request
	req, err = http.NewRequest(method, url, reqJson)

	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	if token != "" {

		req.Header.Add("Authorization", "Bearer "+token)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, http.StatusRequestTimeout, err
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return respBody, resp.StatusCode, nil
}
