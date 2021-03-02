package user

import (
	"fmt"
	keycloak "zeus/keycloak"

	"github.com/gin-gonic/gin"
)

func Register_user(c *gin.Context) {
	access_token := keycloak.GetAccessToken()
	if access_token == "" {
		fmt.Println("error")
	}
	bindUser := InitUserInfo()
	c.BindJSON(bindUser)

}
