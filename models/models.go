package models

import (
	"time"

	"gorm.io/gorm"
)

type Install_Device_Tbls struct {
	//gorm.Model
	Manufacturer        string    `gorm:"column:manufacturer" json:manufacturer`
	DeviceCode          string    `gorm:"column:device_code" json:deviceCode`
	ProductSerialNumber string    `gorm:"column:product_serial_number" json:productSerialNumber`
	TunnelNumber        int       `gorm:"column:tunnel_number" json:tunnelNumber`
	TunnelName          string    `gorm:"column:tunnel_name" json:tunnelName`
	ModelStatus         int       `gorm:"column:model_status" json:modelStatus`
	InstallDt           time.Time `gorm:"column:install_dt" json:installDt`
}

func (Install_Device_Tbls) TableName() string {
	return "install_device_tbls"
}

type UserInfo struct {
	Username   string        `binding:"required" json:"username"`
	FirstName  string        `binding:"required" json:"firstName"`
	LastName   string        `binding:"required" json:"lastName"`
	Enabled    string        `binding:"required" json:"enabled"`
	Email      string        `binding:"required" json:"email"`
	Attributes userAttribute `binding:"required" json:"attributes"`
}

type userAttribute struct {
	DepartmentNm string `json:"departmentNm, string" binding:"required"`
	Position     string `json:"position, string" binding:"required"`
	PhoneNumber  string `json:"phoneNumber, string" binding:"required"`
}

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
	Icon_Code           string `gorm:"column:icon_code" json:"icon_code" binding:"required"`
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
	Icon_Code           string `gorm:"column:icon_code" json:"icon_code" binding:"required"`
}

func (SubMenuInfo) TableName() string {
	return "sub_menu_tbl_test"
}

type TopMenuIcon struct {
	Icon_Code string `gorm:"column:icon_code" json:"icon_code" binding:"required"`
	Icon_Name string `gorm:"column:icon_name" json:"icon_name" binding:"required"`
}

func (TopMenuIcon) TableName() string {
	return "top_menu_icon_tbl"
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
