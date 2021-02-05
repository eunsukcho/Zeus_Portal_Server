package devices

import (
	"net/http"
	"zeus/model"

	"github.com/gin-gonic/gin"
)

func GetAllData(c *gin.Context) {
	install_device_tbls := model.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    install_device_tbls,
		"message": nil,
	})
}

func GetOneManufacturerData(c *gin.Context) {
	manufacturer := c.Param("manufacturer")
	install_device_tbls := model.GetDeviceByManufactureID(manufacturer)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    install_device_tbls,
		"message": nil,
	})
}
