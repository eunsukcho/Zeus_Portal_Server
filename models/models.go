package models

import (
	"gorm.io/gorm"
)

type Invitation struct {
	AccessAuth        string `json:"AccessAuth"`
	InvitationAddress string `json:"InvitationAddress"`
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
	DeletedAt           gorm.DeletedAt `gorm:"-"`
	Top_Menu_Code       string         `gorm:"column:top_menu_code" json:"top_menu_code" binding:"required"`
	Top_Menu_Name       string         `gorm:"column:top_menu_name" json:"top_menu_name"`
	Top_Menu_Target_Url string         `gorm:"column:top_menu_target_url" json:"top_menu_target_url"`
	Top_Menu_Order      string         `gorm:"column:top_menu_order" json:"top_menu_order"`
	Icon_Code           string         `gorm:"column:icon_code" json:"icon_code"`
	New_Window          string         `gorm:"column:new_window" json:"new_window"`
	Is_Main             *bool          `gorm:"column:is_main" json:"is_main"`
}

func (TopMenuInfo) TableName() string {
	return "top_menu_tbl"
}

type SubMenuInfo struct {
	gorm.Model
	Top_Menu_Code       string `gorm:"column:top_menu_code;ForeignKey:top_menu_code" json:"top_menu_code"`
	Sub_Menu_Code       string `gorm:"column:sub_menu_code" json:"sub_menu_code" binding:"required"`
	Top_Menu_Name       string `gorm:"column:top_menu_name" json:"top_menu_name"`
	Sub_Menu_Name       string `gorm:"column:sub_menu_name" json:"sub_menu_name"`
	Sub_Menu_Target_Url string `gorm:"column:sub_menu_target_url" json:"sub_menu_target_url"`
	Sub_Menu_Order      string `gorm:"column:sub_menu_order" json:"sub_menu_order"`
	Icon_Code           string `gorm:"column:icon_code" json:"icon_code"`
	New_Window          string `gorm:"column:new_window" json:"new_window"`
}

func (SubMenuInfo) TableName() string {
	return "sub_menu_tbl"
}

type TopMenuIcon struct {
	Icon_Code        string `gorm:"column:icon_code" json:"icon_code" binding:"required"`
	Icon_Name        string `gorm:"column:icon_name" json:"icon_name" binding:"required"`
	Icon_Description string `gorm:"column:icon_description" json:"icon_description" binding:"required"`
}

func (TopMenuIcon) TableName() string {
	return "top_menu_icon_tbl"
}

type Env_setting_Tbls struct {
	ID               uint   `gorm:"primarykey"`
	ThemeSettingVal  string `gorm:"column:theme_setting_val" json:"themeSettingVal"`
	LangSettingVal   string `gorm:"column:lang_setting_val"  json:"langSettingVal"`
	AutoLogoutVal    bool   `gorm:"column:auto_logout_val"  json:"autoLogoutVal"`
	PortalVersion    int    `gorm:"column:portal_version" json:"portalVersion"`
	UserRegisterAuth bool   `gorm:"column:user_register_auth" json:"userRegisterAuth"`
	GrafanaToken     string `gorm:"column:grafana_token" json:"grafanaToken"`
}
type Envs struct {
	ThemeSettingVal  string `binding:"required" json:"themeSettingVal"`
	LangSettingVal   string `binding:"required" json:"langSettingVal"`
	AutoLogoutVal    bool   `binding:"required" json:"autoLogoutVal"`
	PortalVersion    int    `binding:"required" json:"zoneVersionportalVersion"`
	UserRegisterAuth bool   `binding:"required" json:"userRegisterAuth"`
	GrafanaToken     string `gorm:"column:grafana_token" json:"grafanaToken"`
}

func (Env_setting_Tbls) TableName() string {
	return "env_setting_tbl"
}

type Dev_Info struct {
	Dev_User_Id uint   `gorm:"primarykey"`
	Dev_info    string `gorm:"column:dev_info" json:"devInfo"`
	Enabled     bool   `gorm:"column:enabled" json:"enabled"`
	GroupName   string `gorm:"column:groupname" json:"groupName"`
	Email       string `gorm:"column:email" json:"email"`
}

func (Dev_Info) TableName() string {
	return "devuser_tmp_tbl"
}

type Res_Dev_Info struct {
	Dev_info []string `json:"devInfo"`
}

func (resDevInfo *Res_Dev_Info) AddInfo(info string) []string {
	resDevInfo.Dev_info = append(resDevInfo.Dev_info, info)
	return resDevInfo.Dev_info
}

type LogType_Code struct {
	LogType     string `gorm:"column:logtype" json:"logType"`
	LogTypeName string `gorm:"column:logtype_nm" json:"logTypeNm"`
}

func (LogType_Code) TableName() string {
	return "log_code_tbl"
}

// Binding Uri
type Uri struct {
	Id          string `uri:"id"`
	ReqId       uint   `uri:"reqId"`
	TopCode     string `uri:"topCode"`
	SubCode     string `uri:"subCode"`
	Order       int    `uri:"order"`
	TopCodeName string `uri:"topCodeName"`
	Table       string `uri:"table"`
}
