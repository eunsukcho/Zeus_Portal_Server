package httpd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"zeus/models"
	"zeus/requestLayer"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var errConnFail = errors.New("Connection Failed")

type RequestHandlerInterface interface {
	//user setting
	errHandler(err_str string, err error) (bool, string)
	UserList(c *gin.Context)
	UserListByGroup(c *gin.Context)
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
	token    *oauth2.Token
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
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Binding Error"})
		return
	}
	var rst bool
	if h.authInfo != nil {
		rst, _ = CompareInfo(authBinding.Admin, h.authInfo)
	}
	auth, _ := requestLayer.NewAuthInfo(authBinding.Admin)
	h.authInfo = auth

	if !rst || h.token == nil { // Auth 정보가 변경되었거나 Token이 null 일 때
		h.token, _ = auth.GetApiClientTokenSource(h.ctx)
	}
	var clientErr error
	h.client, clientErr = requestLayer.GetClient(h.ctx, h.token)
	if clientErr != nil {
		return
	}

	c.Set("User", authBinding.User)
	c.Next()
}
func (h *RequestHandler) GroupClientInit(c *gin.Context) {
	var authBinding models.GroupAdminAPIInfo

	if err := c.ShouldBindJSON(&authBinding); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": false, "message": "Binding Error"})
		return
	}
	fmt.Println("Group init : ", authBinding)
	var rst bool
	if h.authInfo != nil {
		rst, _ = CompareInfo(authBinding.Admin, h.authInfo)
	}
	auth, _ := requestLayer.NewAuthInfo(authBinding.Admin)
	h.authInfo = auth

	if !rst || h.token == nil { // Auth 정보가 변경되었거나 Token이 null 일 때
		h.token, _ = auth.GetApiClientTokenSource(h.ctx)
	}
	var clientErr error
	h.client, clientErr = requestLayer.GetClient(h.ctx, h.token)
	if clientErr != nil {
		return
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

func (h *RequestHandler) errHandler(err_str string, err error) (bool, string) {
	var clientErr error
	var status bool = false
	var message string
	message = err.Error()

	auth := h.authInfo
	fmt.Println("Err Error : ", err.Error())
	if err.Error() == "Connection Failed" {

		h.token, _ = auth.GetApiClientTokenSource(h.ctx)
		h.client, clientErr = requestLayer.GetClient(h.ctx, h.token)
		if clientErr != nil {
			return status, ""
		}
		message = "Client Connection False"
	} else {
		fmt.Println("err != errConnFail")
	}
	return status, message
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
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Binding Error",
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	// End point에 따라 분기
	h.RequestUserInfo(uri.Id, c)
}

func (h *RequestHandler) UserListByGroup(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		fmt.Println("error bind uri")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Binding Error",
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	var userinfo []models.ResponseUserInfo
	var err error

	userinfo, err = h.requestH.RequestUserListByGroupApi(h.ctx, uri.Id, h.client)
	if err != nil {
		status, err_str := h.errHandler("RequestUserListByGroupApi error 발생", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  status,
			"message": err_str,
		})
		return
	}
	for _, value := range userinfo {
		fmt.Println("userList :", value)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   userinfo,
	})
}

func (h *RequestHandler) RequestUserInfo(id string, c *gin.Context) {
	if id == "all" {
		var userinfo []models.ResponseUserInfo
		var err error

		userinfo, err = h.requestH.RequestUserListApi(h.ctx, h.client)
		if err != nil {
			status, err_str := h.errHandler("RequestUserListApi error 발생", err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status":  status,
				"message": err_str,
			})
			return
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
			status, err_str := h.errHandler("RequestUserListApi error 발생", err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status":  status,
				"message": err_str,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   userinfo,
		})
	}
}

func (h *RequestHandler) RegisterUser(c *gin.Context) {

	regi := c.MustGet("User").(models.RegisterUserInfo)
	fmt.Println("Regi : ", regi)
	rst, statusCode, err := h.requestH.RequestRegisterUserApi(h.ctx, regi, h.client)

	if err != nil {
		status, err_str := h.errHandler("Request register error 발생", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  status,
			"message": err_str,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": statusCode,
		"data":   rst,
	})
}

func (h *RequestHandler) DeleteUser(c *gin.Context) {

	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	rst, err := h.requestH.DeleteUserApi(h.ctx, uri.Id, h.client)
	if err != nil {
		status, err_str := h.errHandler("Request user delete error 발생", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  status,
			"message": err_str,
		})
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
		status, err_str := h.errHandler("Request user update error 발생", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  status,
			"message": err_str,
		})
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
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	rst, err := h.requestH.UpdateUserCredentialsApi(h.ctx, uri.Id, h.client)

	if err != nil {
		status, err_str := h.errHandler("Request user update credentials error 발생", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  status,
			"message": err_str,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}

func (h *RequestHandler) GroupsList(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.Id)

	var groups []models.ResGroupInfo
	var err error

	groups, err = h.requestH.RequestGroupListApi(h.ctx, uri.Id, h.client)
	if err != nil {
		status, err_str := h.errHandler("Request group list error 발생", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  status,
			"message": err_str,
		})
		return
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
		status, err_str := h.errHandler("Request group token error 발생", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  status,
			"message": err_str,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   rst,
	})
}
