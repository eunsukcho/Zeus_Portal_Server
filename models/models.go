package models

import (
	"gorm.io/gorm"
)

type SmtpInfo struct {
	gorm.Model
	AdminAddress string `gorm:"column:smtp_user" json:"AdminAddress" binding:"required"`
	SmtpAddress  string `gorm:"column:smtp_host" json:"SmtpAddress" binding:"required"`
	Port         string `gorm:"column:smtp_port" json:"Port" binding:"required"`
	Password     string `gorm:"column:smtp_password" json:"Password" binding:"required"`
}

func (SmtpInfo) TableName() string {
	return "smtp_setting_tbl"
}

type TopMenuInfo struct {
	gorm.Model
	Top_Menu_Code       string `gorm:"column:top_menu_code" json:"top_menu_code" binding:"required`
	Top_Menu_Name       string `gorm:"column:top_menu_name" json:"top_menu_name" binding:"required`
	Top_Menu_Target_Url string `gorm:"column:top_menu_target_url" json:"top_menu_target_url" binding:"required`
	Top_Menu_Order      string `gorm:"column:top_menu_order" json:"top_menu_order" binding:"required`
}

func (TopMenuInfo) TableName() string {
	return "top_menu_tbl_test"
}

type SubMenuInfo struct {
	gorm.Model
	Top_Menu_Code       string `gorm:"column:top_menu_code;ForeignKey:top_menu_code" json:"top_menu_code" binding:"required`
	Sub_Menu_Code       string `gorm:"column:sub_menu_code" json:"sub_menu_code" binding:"required`
	Top_Menu_Name       string `gorm:"column:top_menu_name" json:"top_menu_name" binding:"required`
	Sub_Menu_Name       string `gorm:"column:sub_menu_name" json:"sub_menu_name" binding:"required`
	Sub_Menu_Target_Url string `gorm:"column:sub_menu_target_url" json:"sub_menu_target_url" binding:"required`
	Sub_Menu_Order      string `gorm:"column:sub_menu_order" json:"sub_menu_order" binding:"required`
}

func (SubMenuInfo) TableName() string {
	return "sub_menu_tbl_test"
}

type Env_setting_Tbls struct {
	ThemeSettingVal string `gorm:"column:theme_setting_val" binding:"required" json:themeSettingVal`
	LangSettingVal  string `gorm:"column:lang_setting_val" binding:"required" json:langSettingVal`
	AutoLogoutVal   string `gorm:"column:auto_logout_val" binding:"required" json:autoLogoutVal`
	PortalVersion   int    `gorm:"column:portal_version" binding:"required" json:portalVersion`
}
type Envs struct {
	ThemeSettingVal string `binding:"required" json:themeSettingVal`
	LangSettingVal  string `binding:"required" json:langSettingVal`
	AutoLogoutVal   string `binding:"required" json:autoLogoutVal`
	PortalVersion   int    `binding:"required" json:portalVersion`
}

func (Env_setting_Tbls) TableName() string {
	return "env_setting_tbls"
}
