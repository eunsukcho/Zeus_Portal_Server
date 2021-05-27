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

	var updateTbl models.Env_setting_Tbls
	return envs, db.Model(&updateTbl).Updates(envs).Error
}
func (db *DBORM) GetLogCode() (logCode []models.LogType_Code, err error) {

	return logCode, db.Find(&logCode).Error
}
