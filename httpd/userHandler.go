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

func (h *RequestHandler) RegisterUser(c *gin.Context) {
	var userInfo models.UserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*
		ctx := h.ctx
		var client = h.requestH.GetClient(ctx)
		var err error

			if v := h.ctx.Value(oauth2.HTTPClient); v == nil {
				fmt.Println("ㄴㄴ")
				client, err = h.requestH.GetClient(ctx)
				if err != nil {
					panic(err)
				}
			}


		client, err = h.requestH.GetClient(ctx)
		client2, err2 := h.requestH.GetClient(ctx)

		if err2 != nil {
			fmt.Println(client2)
		}
	*/
	err := h.requestH.RequestRegisterUserApi(h.ctx, userInfo, h.client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}
