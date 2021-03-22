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
	userApi := r.Group("/api/user")
	{
		userApi.POST("/infoInit", rh.UserClientInit)
		userApi.POST("/user_list", rh.UserList)
		userApi.POST("/register_user", rh.RegisterUser)
	}
	groupApi := r.Group("/api/groups")
	{
		groupApi.POST("/lt", rh.GroupsList)
		groupApi.POST("/putKey", rh.RegisterToken)
	}

	return r.Run(address)
}
