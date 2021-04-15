package dblayer

import "zeus/models"

type DBLayer interface {
	// init zeus env
	GetAllEnvData() ([]models.Env_setting_Tbls, error)
	UpdateEnvData(models.Env_setting_Tbls) (models.Env_setting_Tbls, error)

	// menu setting
	GetTopMenuInfoByName(string) (models.TopMenuInfo, error)
	GetAllTopMenu() ([]models.TopMenuInfo, error)
	GetAllSubMenu() ([]models.SubMenuInfo, error)
	CkDuplicateTopMenu(string) (int, error)
	CkDuplicateTopMenuOrder(order int) (rst int, err error)
	SaveTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
	CkDuplicateSubMenu(string, string) (int, error)
	CkDuplicateSubMenuOrder(string, string, int) (rst int, err error)
	SaveSubMenuInfo(models.SubMenuInfo) (models.SubMenuInfo, error)
	DeleteTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
	DeleteSubMenuByTopCodeUrl(string) (models.SubMenuInfo, error)
	DeleteSubMenuInfo(models.SubMenuInfo) (models.SubMenuInfo, error)
	GetAllIcon() ([]models.TopMenuIcon, error)
	SaveUrlLink(models.TopMenuInfo) (models.TopMenuInfo, error)
	SaveUrlSubLink(models.SubMenuInfo) (models.SubMenuInfo, error)
	DeleteTopMenuUrl(models.TopMenuInfo) (models.TopMenuInfo, error)
	DeleteSubMenuUrl(sub models.SubMenuInfo) (models.SubMenuInfo, error)
	GetMenuTargetUrl(models.SubMenuInfo) (models.SubMenuInfo, error)
	GetTopMenuTargetUrl(models.TopMenuInfo) (models.TopMenuInfo, error)
	UpdateTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error)
	UpdateSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error)
	
	//smtp setting
	SmtpInfoConnectionCheck() ([]models.SmtpInfo, error)
	SmtpInfoSave(models.SmtpInfo) (models.SmtpInfo, error)
	SmtpInfoTest() ([]models.SmtpInfo, error)
	SmtpInfoGet() ([]models.SmtpInfo, error)

	//auth setting
	GetAllAuthData() ([]models.Authdetails, error)
	SaveAuthData(models.Authdetails) ([]models.Authdetails, error)

	SaveDevUserInfo(models.Dev_Info) (models.Dev_Info, error)
	GetDevUserInfo(string) ([]models.RegisterUserInfo, error)
	AcceptUpdateUser(string) (models.Dev_Info, error)
}
