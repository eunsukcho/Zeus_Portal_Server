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
	UserClientInit(c *gin.Context)

	GroupsList(c *gin.Context)
	RegisterToken(c *gin.Context)
}

type RequestHandler struct {
	requestH requestLayer.RequestLayer
	ctx      context.Context
	client   *http.Client
}

func NewRequestHandler() (RequestHandlerInterface, error) {
	var authBinding models.Authdetails
	auth, err := requestLayer.NewAuthInfo(authBinding)
	if err != nil {
		return nil, err
	}
	return &RequestHandler{
		requestH: auth,
		ctx : context.Background(),
	}, nil
}

func (h *RequestHandler) UserClientInit(c *gin.Context) {
	var authBinding models.Authdetails
	if err := c.ShouldBindJSON(&authBinding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("authBinding List : " , authBinding)
	auth, _ := requestLayer.NewAuthInfo(authBinding)
	apiClient, _ := requestLayer.GetClient(h.ctx, auth)
	
	h.client = apiClient
}

func (h *RequestHandler) UserList(c *gin.Context) {
	if h.client == nil {
		fmt.Println("h.client is nill")
		h.UserClientInit(c)
	}

	var userinfo []models.ResponseUserInfo
	var err error

	userinfo, err = h.requestH.RequestUserListApi(h.ctx, h.client)
	if err != nil {
		fmt.Println("requestList error 발생")
		h.UserClientInit(c)
		userinfo, err = h.requestH.RequestUserListApi(h.ctx, h.client)
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
	if h.client == nil {
		fmt.Println("h.client is nill")
		h.UserClientInit(c)
	}

	var regi models.AdminAPIInfo
	if err := c.ShouldBindJSON(&regi); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("userinfo : ", regi)

	rst, err := h.requestH.RequestRegisterUserApi(h.ctx, regi.User, h.client)
	if err != nil {
		fmt.Println("request register error 발생")
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}

func (h *RequestHandler) GroupsList(c *gin.Context) {
	if h.client == nil {
		fmt.Println("h.client is nill")
		h.UserClientInit(c)
	}

	var groups []models.ResGroupInfo
	var err error

	groups, err = h.requestH.RequestGroupListApi(h.ctx, h.client)
	if err != nil {
		fmt.Println("requestGroups error 발생")
		h.UserClientInit(c)
		groups, err = h.requestH.RequestGroupListApi(h.ctx, h.client)
	}
	for _, value := range groups {
		fmt.Println(value)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   groups,
	})
}

func (h *RequestHandler) RegisterToken(c *gin.Context) {
	if h.client == nil {
		fmt.Println("h.client is nill")
		h.UserClientInit(c)
	}

	var regi models.ReqToken
	if err := c.ShouldBindJSON(&regi); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("ResGroupInfo : ", regi)

	rst, err := h.requestH.RequestRegisterGroupsApi(h.ctx, regi, h.client)
	if err != nil {
		fmt.Println("request register group add  error 발생")
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}