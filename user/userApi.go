package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register_user(c *gin.Context) {
	//keycloak.SetHeader(c)
	bindUser := InitUserInfo()
	fmt.Println("bindUser : ", bindUser)
}
