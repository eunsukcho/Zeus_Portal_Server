package main

import (
	"zeus/devices"
	"zeus/env"
	"zeus/menu"
	"zeus/smtpconnect"
	"zeus/user"

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

	envApi := r.Group("/api")
	{
		envApi.GET("/systemInfo", env.GetAllEnvData)
		envApi.POST("/changeTheme", env.UpdateEnvData)
	}

	smtpApi := r.Group("/smtp")
	{
		smtpApi.POST("/register_smtp", smtpconnect.Smtptest)
		smtpApi.POST("/smtpsave", smtpconnect.SmtpSave)
	}

	userApi := r.Group("/user")
	{
		userApi.POST("/register_user", user.Register_user)
	}

	menuApi := r.Group("/menu")
	{
		menuApi.GET("/topmenu", menu.GetTopMenuData)
		menuApi.GET("/submenu", menu.SubTopMenuData)
		menuApi.POST("/topmenusave", menu.SaveTopMenu)
		menuApi.POST("/submenusave", menu.SaveSubMenu)
		menuApi.POST("/topmenudelete", menu.DeleteTopMenu)
		menuApi.POST("/submenudelete", menu.DeleteSubMenu)
	}

	r.Run("127.0.0.1:3000")
}
