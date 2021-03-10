package dblayer

import (
	"fmt"
	"zeus/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBORM struct {
	*gorm.DB
}

func NewDBInit() (*DBORM, error) {
	dbConfig := settingDB()

	postgres_conn_name := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		dbConfig.db_host, dbConfig.db_port, dbConfig.db_username, dbConfig.db_name, dbConfig.db_password)
	fmt.Println("conname \n", postgres_conn_name)

	db, err := gorm.Open("postgres", postgres_conn_name)

	return &DBORM{
		DB: db,
	}, err
}

// init zeus env
func (db *DBORM) GetAllEnvData() (envs []models.Env_setting_Tbls, err error) {

	return envs, db.Find(&envs).Error
}
func (db *DBORM) UpdateEnvData(envs models.Envs) (envInfo models.Env_setting_Tbls, err error) {

	theme := envs.ThemeSettingVal
	lang := envs.LangSettingVal
	autoLogout := envs.AutoLogoutVal
	version := envs.PortalVersion

	return envInfo, db.Model(&envInfo).Updates(models.Env_setting_Tbls{ThemeSettingVal: theme, LangSettingVal: lang, AutoLogoutVal: autoLogout, PortalVersion: version}).Error
}

// menu setting
func (db *DBORM) GetAllTopMenu() (top []models.TopMenuInfo, err error) {
	return top, db.Order("top_menu_order asc").Find(&top).Error
}
func (db *DBORM) GetAllSubMenu() (sub []models.SubMenuInfo, err error) {
	return sub, db.Order("sub_menu_order asc").Find(&sub).Error
}
func (db *DBORM) SaveTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Create(&top).Error
}
func (db *DBORM) SaveSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Create(&sub).Error
}
func (db *DBORM) DeleteTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Where("top_menu_code = ? ", top.Top_Menu_Code).Delete(&top).Error
}
func (db *DBORM) DeleteSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Where("sub_menu_code = ? ", sub.Sub_Menu_Code).Delete(&sub).Error
}

//smtp setting
func (db *DBORM) SmtpInfoConnectionCheck() ([]models.SmtpInfo, error) {
	return nil, nil
}
func (db *DBORM) SmtpInfoSave(smtp models.SmtpInfo, password []byte) (models.SmtpInfo, error) {
	db.Create(&smtp)
	return smtp, db.Model(&smtp).Where("smtp_password = ?", smtp.Password).Update("smtp_password", password).Error
}
func (db *DBORM) SmtpInfoTest() ([]models.SmtpInfo, error) {
	return nil, nil
}
