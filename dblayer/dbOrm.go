package dblayer

import (
	"zeus/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// init zeus env
func (db *DBORM) GetAllEnvData() (envs []models.Env_setting_Tbls, err error) {

	return envs, db.Find(&envs).Error
}
func (db *DBORM) UpdateEnvData(envs models.Env_setting_Tbls) (envInfo models.Env_setting_Tbls, err error) {

	var id uint
	id = 1
	theme := envs.ThemeSettingVal
	lang := envs.LangSettingVal
	autoLogout := envs.AutoLogoutVal
	version := envs.PortalVersion
	userAuth := envs.UserRegisterAuth

	return envInfo, db.Model(&envInfo).Where("id=?", id).Updates(map[string]interface{}{"ThemeSettingVal": theme, "LangSettingVal": lang, "AutoLogoutVal": autoLogout, "PortalVersion": version, "UserRegisterAuth": userAuth}).Error

}
func (db *DBORM) GetLogCode() (logCode []models.LogType_Code, err error) {

	return logCode, db.Find(&logCode).Error
}
