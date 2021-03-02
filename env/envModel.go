package env

import model "zeus/initModel"

type Env_setting_Tbls struct {
	ThemeSettingVal string `gorm:"column:theme_setting_val" binding:"required" json:themeSettingVal`
	LangSettingVal  string `gorm:"column:lang_setting_val" binding:"required" json:langSettingVal`
	AutoLogoutVal   string `gorm:"column:auto_logout_val" binding:"required" json:autoLogoutVal`
	PortalVersion   int    `gorm:"column:portal_version" binding:"required" json:portalVersion`
}
type Env struct {
	ThemeSettingVal string `binding:"required" json:themeSettingVal`
	LangSettingVal  string `binding:"required" json:langSettingVal`
	AutoLogoutVal   string `binding:"required" json:autoLogoutVal`
	PortalVersion   int    `binding:"required" json:portalVersion`
}

func (Env_setting_Tbls) TableName() string {
	return "env_setting_tbls"
}

func GetEnvAll() *[]Env_setting_Tbls {
	db := model.DbInit()
	defer db.Close()

	var env_setting_tbls []Env_setting_Tbls
	db.Find(&env_setting_tbls)

	return &env_setting_tbls
}

func UpdateEnvVal(env *Env) *[]Env_setting_Tbls {
	db := model.DbInit()
	defer db.Close()

	var env_setting_tbls []Env_setting_Tbls

	db.Model(&env_setting_tbls).Update("theme_setting_val", env.ThemeSettingVal)

	return &env_setting_tbls
}

func GetDeviceByManufactureID(manufacturer string) *Env_setting_Tbls {
	db := model.DbInit()
	defer db.Close()

	var env_setting_tbls Env_setting_Tbls
	db.Find(&env_setting_tbls)

	return &env_setting_tbls
}
