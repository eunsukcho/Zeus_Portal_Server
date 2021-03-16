package httpd

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	rh, err := NewRequestHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h, rh)
}

func RunAPIWithHandler(address string, h HandlerInterface, rh RequestHandlerInterface) error {
	r := gin.Default()

	envApi := r.Group("/api")
	{
		envApi.GET("/systemInfo", h.GetEnvData)
		envApi.POST("/changeTheme", h.UpdateEnvData)
	}
	smtpApi := r.Group("/smtp")
	{
		smtpApi.POST("/register_smtp", h.Smtptest)
		smtpApi.POST("/smtpsave", h.SmtpSave)
		smtpApi.GET("/smtpget", h.SmtpGet)
	}
	menuApi := r.Group("/menu")
	{
		menuApi.GET("/topmenu", h.GetTopMenuData)
		menuApi.GET("/submenu", h.SubTopMenuData)
		menuApi.GET("/topmenuicon", h.GetIcon)
		menuApi.POST("/topmenusave", h.SaveTopMenu)
		menuApi.POST("/submenusave", h.SaveSubMenu)
		menuApi.POST("/topmenudelete", h.DeleteTopMenu)
		menuApi.POST("/submenudelete", h.DeleteSubMenu)
		menuApi.POST("/topmenusaveUrl", h.SaveUrlLink)
		menuApi.POST("/submenusaveUrl", h.SaveUrlSubLink)
		menuApi.POST("/target", h.GetMenuTargetUrl)
	}
	userApi := r.Group("/user")
	{
		userApi.POST("/infoInit", rh.UserClientInit)
		userApi.POST("/user_list", rh.UserList)
		userApi.POST("/register_user", rh.RegisterUser)
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
