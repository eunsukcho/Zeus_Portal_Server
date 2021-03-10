package dblayer

import "zeus/models"

type DBLayer interface {
	// init zeus env
	GetAllEnvData() ([]models.Env_setting_Tbls, error)
	UpdateEnvData(models.Envs) (models.Env_setting_Tbls, error)

	// menu setting
	GetAllTopMenu() ([]models.TopMenuInfo, error)
	GetAllSubMenu() ([]models.SubMenuInfo, error)
	SaveTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
	SaveSubMenuInfo(models.SubMenuInfo) (models.SubMenuInfo, error)
	DeleteTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
	DeleteSubMenuInfo(models.SubMenuInfo) (models.SubMenuInfo, error)

	//smtp setting
	SmtpInfoConnectionCheck() ([]models.SmtpInfo, error)
	SmtpInfoSave(models.SmtpInfo, []byte) (models.SmtpInfo, error)
	SmtpInfoTest() ([]models.SmtpInfo, error)
}
