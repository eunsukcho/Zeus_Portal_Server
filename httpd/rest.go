package httpd

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandlerWithParams()
	if err != nil {
		return err
	}

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
	menuApi.Use(h.DBConnectionCheck)
	{
		menuApi.POST("/topMenuInfoByName/:topCodeName", h.GetTopMenuInfoByName)
		menuApi.GET("/topmenu", h.GetTopMenuData)
		menuApi.GET("/submenu", h.SubTopMenuData)
		menuApi.GET("/topmenuicon", h.GetIcon)
		menuApi.POST("/ckDuplicateTopMenu/:topCode/:order", h.CkDuplicateTopMenu)
		menuApi.POST("/topmenusave", h.SaveTopMenu)
		menuApi.POST("/ckDuplicateSubMenu/:topCode/:subCode/:order", h.CkDuplicateSubMenu)
		menuApi.POST("/submenusave", h.SaveSubMenu)
		menuApi.POST("/topmenudelete", h.DeleteTopMenu)
		menuApi.POST("/submenudelete", h.DeleteSubMenu)
		menuApi.POST("/topmenusaveUrl", h.SaveUrlLink)
		menuApi.POST("/submenusaveUrl", h.SaveUrlSubLink)
		menuApi.POST("/topmenudeleteUrl", h.DeleteTopMenuUrl)
		menuApi.POST("/deleteSubMenuByTopMenuCode/:topCodeName", h.DeleteSubMenuByTopCodeUrl)
		menuApi.POST("/submenudeleteUrl", h.DeleteSubMenuUrl)
		menuApi.POST("/target", h.GetMenuTargetUrl)
		menuApi.POST("/toptarget", h.GetTopMenuTargetUrl)
		menuApi.POST("/topmenuupdate", h.UpdateTopMenuInfo)
		menuApi.POST("/submenuupdate", h.UpdateSubMenuInfo)
		menuApi.POST("/updateSubMenuInfoByTopMenuCode/:topCode/:topCodeName", h.UpdateSubMenuTopCodeName)
	}
	authApi := r.Group("/api/auth")
	{
		authApi.POST("/auth_list", h.AuthInfoData)
		authApi.POST("/save_auth", h.SaveAuthData)
	}
	devApi := r.Group("/api/dev")
	devApi.Use(h.DBConnectionCheck)
	{
		devApi.POST("/userListByGroupTmp/:id", h.GetDevUser)
		devApi.POST("/register_user_dev", h.CreateDevUser)
		devApi.POST("/acceptDev/:id", h.AcceptUser)
	}
	adminApi := r.Group("/api/admin")
	adminApi.Use(rh.UserClientInit)
	{
		adminApi.POST("/infoInit", rh.UserClientInit)
		adminApi.POST("/user_list/:id", rh.UserList)
		adminApi.POST("/userListByGroup/:id/members", rh.UserListByGroup)
		adminApi.POST("/register_user", rh.RegisterUser)
		adminApi.POST("/delete_user/:id", rh.DeleteUser)
		adminApi.POST("/update_user", rh.UpdateUser)
		adminApi.POST("/update_userCredentials/:id", rh.UpdateUserCredentials)
	}
	groupApi := r.Group("/api/groups")
	groupApi.Use(rh.GroupClientInit)
	{
		groupApi.POST("/lt/:id", rh.GroupsList)
		groupApi.POST("/putKey", rh.RegisterToken)
	}

	druidApi := r.Group("/api/druid")
	{
		druidApi.GET("/:table", h.GetColumnSearchInfo)
		druidApi.POST("/val", h.GetLogValue)
	}
	return r.Run(address)
}
