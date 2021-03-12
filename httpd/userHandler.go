package httpd

import (
	"context"
	"fmt"
	"net/http"
	"zeus/models"
	"zeus/requestLayer"

	"github.com/gin-gonic/gin"
)

type RequestHandlerInterface interface {
	//user setting
	UserList(c *gin.Context)
	RegisterUser(c *gin.Context)
}

type RequestHandler struct {
	requestH requestLayer.RequestLayer
	ctx      context.Context
	client   *http.Client
}

func NewRequestHandler() (RequestHandlerInterface, error) {
	ctx := context.Background()
	auth, err := requestLayer.NewAuthInfo()
	apiClient, err := requestLayer.GetClient(ctx, auth)
	if err != nil {
		return nil, err
	}
	fmt.Printf("new request handler")
	return &RequestHandler{
		requestH: auth,
		ctx:      ctx,
		client:   apiClient,
	}, nil
}

func (h *RequestHandler) UserList(c *gin.Context) {

	userinfo, err := h.requestH.RequestUserListApi(h.ctx, h.client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, value := range userinfo {
		fmt.Println(value)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   userinfo,
	})
}

func (h *RequestHandler) RegisterUser(c *gin.Context) {

	var userInfo models.RegisterUserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("userinfo : ", userInfo)

	rst, err := h.requestH.RequestRegisterUserApi(h.ctx, userInfo, h.client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}
