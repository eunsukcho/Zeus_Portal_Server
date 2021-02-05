package main

import (
	"zeus/devices"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/get/project", devices.GetAllData)
	r.GET("/get/project/:manufacturer", devices.GetOneManufacturerData)

	r.Run()

}
