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
		menuApi.GET("/mainView", h.GetMainView)
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
		menuApi.POST("/ckMainUrl", h.CkDuplicateIsMain)
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

	/*
		authApi : keycloak 연동을 위한 Authentication Code Control
	*/
	authApi := r.Group("/api/auth")
	{
		authApi.POST("/auth_list", h.AuthInfoData) /* Get Authentication Info */
		authApi.POST("/save_auth", h.SaveAuthData) /* Save Authentication Code */
	}

	/*
		devApi : 개발자 등록 시 devuser_tmp_tbls control
	*/
	devApi := r.Group("/api/developer")
	devApi.Use(h.DBConnectionCheck)
	{
		devApi.GET("/cnt/:id", h.CkDuplicateTmpDev)      /* 개발자 등록 시 중복 등록 체크 */
		devApi.POST("/users/:id", h.GetDevUser)          /* param - group name*/
		devApi.POST("/user", h.CreateDevUser)            /* param - user json, group name, enable, email(duplicate check) */
		devApi.POST("/commit/user/:reqId", h.AcceptUser) /* param - 임시 등록 ROW ID, devuser_tmp_tbls update */
		devApi.DELETE("/user/:id", h.DeleteTmpUser)      /* param - 임시 등록 ROW ID, devuser_tmp_tbls delete */
	}

	/*
		adminApi : 관리자 정보 control (Keycloak API)
	*/
	adminApi := r.Group("/api/admin")
	adminApi.Use(rh.UserClientInit) /* Authentication 코드를 통한 Token Setting */
	{
		/*
			- 관리자 리스트
			- param - all/id (all-전체 List/id-사용자 개인 정보)
		*/
		adminApi.POST("/users/:id", rh.UserList)

		/*
			- 해당 그룹의 관리자 List
			- param - group id
		*/
		adminApi.POST("/group/user/:id/members", rh.UserListByGroup)

		/*
			- 관리자 등록
			- param - models.RegisterUserInfo
		*/
		adminApi.POST("/user", rh.RegisterUser)

		/*
			- 관리자 삭제
			- param - id (관리자 keycloak id)
		*/
		adminApi.POST("/user/:id", rh.DeleteUser)

		/*
			- 관리자 정보 수정
			- param - models.RegisterUserInfo
		*/
		adminApi.PUT("/user", rh.UpdateUser)

		/*
			- 관리자 비밀번호 수정 요청 (keycloak에서 mail)
			- param - models.RegisterUserInfo
		*/
		adminApi.POST("/credentials/:id", rh.UpdateUserCredentials)
	}

	/*
		groupApi : Group 정보 control (Keycloak API)
	*/
	groupApi := r.Group("/api/groups")
	groupApi.Use(rh.GroupClientInit) /* Authentication 코드를 통한 Token Setting */
	{
		/*
			- Keycloak Group List
			- param - id : all (Group의 전체 List) / id (administrator 그룹명) -> get keycloak group id
		*/
		groupApi.POST("/group/:id", rh.GroupsList)

		/*
			- Keycloak Group Key 등록
			- param - models.ReqToken
		*/
		groupApi.POST("/key", rh.RegisterToken)
	}

	druidApi := r.Group("/api/druid")
	{
		druidApi.GET("/:table", h.GetColumnSearchInfo)
		druidApi.POST("/val", h.GetLogValue)
	}

	k8sNamespaceApi := r.Group("/api/k8s/namespace")
	k8sNamespaceApi.Use(h.BindingModel)
	{
		k8sNamespaceApi.POST("", h.CreateRequestProject)
		k8sNamespaceApi.POST("/resourcequotas", h.CreateRequestResourceQuota)
		k8sNamespaceApi.POST("/serviceAccount", h.CreateServiceAccount)
		k8sNamespaceApi.POST("/roles", h.CreateRole)
		k8sNamespaceApi.POST("/rolesBinding", h.CreateRoleBinding)

		k8sNamespaceApi.POST("/userToken", h.GetUserToken)
	}
	k8sNamespaceApiDelete := r.Group("/api/k8s/namespace")
	k8sNamespaceApiDelete.DELETE("/deleteNamespace/:namespace", h.DeleteNamespace)

	devServerApi := r.Group("/api/devServer")
	{
		devServerApi.GET("", h.GetAllDevServerInfoData)
		devServerApi.GET("/:reqId", h.GetDevServerInfoDataById)

		devServerApi.POST("", h.SaveDevServerInfo)
		devServerApi.PUT("/:reqId", h.UpdateDevServerInfo)
		devServerApi.DELETE("/:reqId", h.DeleteDevServerInfo)
	}

	return r.Run(address)
}
