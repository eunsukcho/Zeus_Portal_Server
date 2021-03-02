package env

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEnvData(c *gin.Context) {
	env_setting_tbls := GetEnvAll()

	c.JSON(http.StatusOK, env_setting_tbls)
}

func UpdateEnvData(c *gin.Context) {
	fmt.Println("UpdateEnvData")

	var env Env
	c.BindJSON(&env)
	fmt.Println(env)
	UpdateEnvVal(&env)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"rst":    "success",
	})
}
