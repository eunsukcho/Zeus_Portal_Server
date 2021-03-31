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
	smtpApi := r.Group("/api/smtp")
	{
		smtpApi.POST("/register_smtp", h.Smtptest)
		smtpApi.POST("/smtpsave", h.SmtpSave)
		smtpApi.GET("/smtpget", h.SmtpGet)
		smtpApi.POST("/sendmail", h.SendMail)
		smtpApi.POST("/invitation", h.InvitationUser)
	}
	menuApi := r.Group("/api/menu")
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
		menuApi.POST("/topmenudeleteUrl", h.DeleteTopMenuUrl)
		menuApi.POST("/submenudeleteUrl", h.DeleteSubMenuUrl)
		menuApi.POST("/target", h.GetMenuTargetUrl)
		menuApi.POST("/toptarget", h.GetTopMenuTargetUrl)
		menuApi.POST("/topmenuupdate", h.UpdateTopMenuInfo)
		menuApi.POST("/submenuupdate", h.UpdateSubMenuInfo)
	}
	authApi := r.Group("/api/auth")
	{
		authApi.POST("/auth_list", h.AuthInfoData)
		authApi.POST("/save_auth", h.SaveAuthData)
	}
	userApi := r.Group("/api/user")
	userApi.Use(rh.UserClientInit)
	{
		userApi.POST("/infoInit", rh.UserClientInit)
		userApi.POST("/user_list/:id", rh.UserList)
		userApi.POST("/userListByGroup/:id/members", rh.UserListByGroup)
		userApi.POST("/register_user", rh.RegisterUser)
		userApi.POST("/delete_user/:id", rh.DeleteUser)
		userApi.POST("/update_user", rh.UpdateUser)
		userApi.POST("/update_userCredentials/:id", rh.UpdateUserCredentials)
	}
	groupApi := r.Group("/api/groups")
	groupApi.Use(rh.GroupClientInit)
	{
		groupApi.POST("/lt/:id", rh.GroupsList)
		groupApi.POST("/putKey", rh.RegisterToken)
	}


	return r.Run(address)
}
