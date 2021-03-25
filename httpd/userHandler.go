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
	RequestUserInfo(id string, c *gin.Context)
	RegisterUser(c *gin.Context)
	UserClientInit(c *gin.Context)
	GroupClientInit(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateUserCredentials(c *gin.Context)

	GroupsList(c *gin.Context)
	RegisterToken(c *gin.Context)
}

type RequestHandler struct {
	requestH requestLayer.RequestLayer
	ctx      context.Context
	client   *http.Client
	authInfo *requestLayer.AuthInfo
}

func NewRequestHandler() (RequestHandlerInterface, error) {
	var authBinding models.Authdetails
	auth, err := requestLayer.NewAuthInfo(authBinding)

	if err != nil {
		return nil, err
	}
	return &RequestHandler{
		requestH: auth,
		ctx:      context.Background(),
	}, nil
}

func (h *RequestHandler) UserClientInit(c *gin.Context) {
	fmt.Println("UserClientInit")
	var authBinding models.AdminAPIInfo

	if err := c.ShouldBindJSON(&authBinding); err != nil {
		fmt.Println("binding error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var rst bool
	if h.authInfo != nil {
		rst, _ = CompareInfo(authBinding.Admin, h.authInfo)
	}
	auth, _ := requestLayer.NewAuthInfo(authBinding.Admin)
	h.authInfo = auth
	fmt.Println("userclient init : ", h.authInfo)
	if rst == false {
		apiClient, err := requestLayer.GetClient(h.ctx, h.authInfo)
		if err != nil {
			c.Abort()
		}
		h.client = apiClient
	}
	fmt.Println("auth User : ", authBinding.User)
	c.Set("User", authBinding.User)
	c.Next()
}
func (h *RequestHandler) GroupClientInit(c *gin.Context) {
	fmt.Println("GroupClientInit")
	var authBinding models.GroupAdminAPIInfo

	if err := c.ShouldBindJSON(&authBinding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Group init : ", authBinding)
	var rst bool
	if h.authInfo != nil {
		rst, _ = CompareInfo(authBinding.Admin, h.authInfo)
	}
	auth, _ := requestLayer.NewAuthInfo(authBinding.Admin)
	h.authInfo = auth

	if rst == false {
		apiClient, err := requestLayer.GetClient(h.ctx, h.authInfo)
		if err != nil {
			c.Abort()
		}
		h.client = apiClient
	}
	c.Set("Groups", authBinding.Groups)
	c.Next()
}

func CompareInfo(inputAuth models.Authdetails, initInfo *requestLayer.AuthInfo) (bool, error) {
	switch {
	case inputAuth.ClientId != initInfo.ClientId:
		return false, nil
	case inputAuth.ClientSecret != initInfo.ClientSecret:
		return false, nil
	case inputAuth.AdminId != initInfo.AdminId:
		return false, nil
	case inputAuth.AdminPw != initInfo.AdminPw:
		return false, nil
	case inputAuth.TokenUrl != initInfo.TokenUrl:
		return false, nil
	}
	return true, nil
}

func (h *RequestHandler) UserList(c *gin.Context) {

	/**
		User List End point
			1) all - all list
			2) user id - user info
	**/
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		fmt.Println("error bind uri")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	// End point에 따라 분기
	h.RequestUserInfo(uri.Id, c)
}

func (h *RequestHandler) RequestUserInfo(id string, c *gin.Context) {
	if id == "all" {
		var userinfo []models.ResponseUserInfo
		var err error

		userinfo, err = h.requestH.RequestUserListApi(h.ctx, h.client)
		if err != nil {
			fmt.Println("RequestUserListApi error 발생")
			h.UserClientInit(c)
			userinfo, err = h.requestH.RequestUserListApi(h.ctx, h.client)
		}
		for _, value := range userinfo {
			fmt.Println("userList :", value)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   userinfo,
		})
	}
	if id != "all" {
		var userinfo models.ResponseUserInfo
		var err error

		userinfo, err = h.requestH.RequestOneUserApi(h.ctx, id, h.client)
		if err != nil {
			fmt.Println("RequestUserListApi error 발생")
			h.UserClientInit(c)
			userinfo, err = h.requestH.RequestOneUserApi(h.ctx, id, h.client)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   userinfo,
		})
	}
}

func (h *RequestHandler) RegisterUser(c *gin.Context) {

	regi := c.MustGet("User").(models.RegisterUserInfo)
	rst, err := h.requestH.RequestRegisterUserApi(h.ctx, regi, h.client)
	if err != nil {
		fmt.Println("request register error 발생")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}

func (h *RequestHandler) DeleteUser(c *gin.Context) {

	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	rst, err := h.requestH.DeleteUserApi(h.ctx, uri.Id, h.client)

	if err != nil {
		fmt.Println("request user delete error 발생")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}
func (h *RequestHandler) UpdateUser(c *gin.Context) {
	regi := c.MustGet("User").(models.RegisterUserInfo)
	rst, err := h.requestH.UpdateUserApi(h.ctx, regi, h.client)
	if err != nil {
		fmt.Println("request register error 발생")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}

func (h *RequestHandler) UpdateUserCredentials(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	rst, err := h.requestH.UpdateUserCredentialsApi(h.ctx, uri.Id, h.client)

	if err != nil {
		fmt.Println("request user delete error 발생")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}
func (h *RequestHandler) GroupsList(c *gin.Context) {

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

	regi := c.MustGet("Groups").(models.ReqToken)
	fmt.Println("test : ", regi)

	rst, err := h.requestH.RequestRegisterGroupsApi(h.ctx, regi, h.client)
	if err != nil {
		fmt.Println("request register group add error 발생")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}
