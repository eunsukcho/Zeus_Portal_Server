package httpd

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	envApi := r.Group("/api")
	{
		envApi.GET("/systemInfo", h.GetEnvData)
		//envApi.POST("/changeTheme", env.UpdateEnvData)
	}
	smtpApi := r.Group("/smtp")
	{
		smtpApi.POST("/register_smtp", h.Smtptest)
		smtpApi.POST("/smtpsave", h.SmtpSave)
	}
	menuApi := r.Group("/menu")
	{
		menuApi.GET("/topmenu", h.GetTopMenuData)
		menuApi.GET("/submenu", h.SubTopMenuData)
		menuApi.POST("/topmenusave", h.SaveTopMenu)
		menuApi.POST("/submenusave", h.SaveSubMenu)
		menuApi.POST("/topmenudelete", h.DeleteTopMenu)
		menuApi.POST("/submenudelete", h.DeleteSubMenu)
	}

	/*
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
	*/
	return r.Run(address)
}
