package main

import (
	"zeus/devices"
	"zeus/env"
	"zeus/menu"
	"zeus/smtpconnect"

	"github.com/gin-gonic/gin"
)

func CORSMiddelware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Cache-Control, Pragma, jsonType, Authorization,Origin")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddelware())

	r.GET("/get/project", devices.GetAllData)
	r.GET("/get/project/:manufacturer", devices.GetOneManufacturerData)

	r.GET("/api/systemInfo", env.GetAllEnvData)

	r.GET("/get/topmenu", menu.GetTopMenuData)
	r.GET("/get/submenu", menu.SubTopMenuData)
	r.POST("/get/topmenusave", menu.SaveTopMenu)
	r.POST("/get/submenusave", menu.SaveSubMenu)

	r.POST("/get/smtp", smtpconnect.Smtptest)
	r.POST("/get/smtpsave", smtpconnect.SmtpSave)

	r.Run("127.0.0.1:3000")
}
