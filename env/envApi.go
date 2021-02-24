package env

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEnvData(c *gin.Context) {
	env_setting_tbls := GetEnvAll()

	c.JSON(http.StatusOK, env_setting_tbls)
}
