package dblayer

import "zeus/models"

type DBLayer interface {
	// init zeus env
	GetAllEnvData() ([]models.Env_setting_Tbls, error)
	UpdateEnvData(models.Env_setting_Tbls) (models.Env_setting_Tbls, error)

	// menu setting
	GetAllTopMenu() ([]models.TopMenuInfo, error)
	GetAllSubMenu() ([]models.SubMenuInfo, error)
	SaveTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
	SaveSubMenuInfo(models.SubMenuInfo) (models.SubMenuInfo, error)
	DeleteTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
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
}
