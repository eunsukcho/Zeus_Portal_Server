package env

import model "zeus/initModel"

type Env_setting_Tbls struct {
	ThemeSettingVal string `gorm:"column:theme_setting_val" json:themeSettingVal`
	LangSettingVal  string `gorm:"column:lang_setting_val" json:langSettingVal`
	AutoLogoutVal   string `gorm:"column:auto_logout_val" json:autoLogoutVal`
	PortalVersion   int    `gorm:"column:portal_version" json:portalVersion`
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

func GetDeviceByManufactureID(manufacturer string) *Env_setting_Tbls {
	db := model.DbInit()
	defer db.Close()

	var env_setting_tbls Env_setting_Tbls
	db.Find(&env_setting_tbls)

	return &env_setting_tbls
}
