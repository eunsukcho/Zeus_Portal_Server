package models

type Authdetails struct {
	ClientId     string `gorm:"column:client_id" json:"clientId" binding:"required"`
	ClientSecret string `gorm:"column:client_secret" json:"clientSecret" binding:"required"`
	AdminId      string `gorm:"column:admin_id" json:"adminId" binding:"required"`
	AdminPw      string `gorm:"column:admin_pw" json:"adminPw" binding:"required"`
	TokenUrl     string `gorm:"column:token_url" json:"tokenUrl" binding:"required"`
}

func (Authdetails) TableName() string {
	return "admin_auth_tbl"
}

// Binding struct
type ReqOption struct {
	Body AdminAPIInfo `json:"body"`
}
type AdminAPIInfo struct {
	User  RegisterUserInfo `json:"user"`
	Admin Authdetails      `binding:"required" json:"admin"`
}
type GroupAdminAPIInfo struct {
	Groups ReqToken    `json:"groups"`
	Admin  Authdetails `binding:"required" json:"admin"`
}

// User Struct
type UserListData struct {
	Id    string      `json:"id"`
	Admin Authdetails `binding:"required" json:"admin"`
}

type RegisterUserInfo struct {
	ID          string `json:"id"`
	Dev_User_Id uint   `json:"dev_user_id,omitempty"`
	UserInfo
	Groups     []string          `json:"groups" gorm:"column:groups"`
	Attributes userAttribute     `json:"attributes"`
	Credential []userCredentilas `json:"credentials"`
}

func (RegisterUserInfo) TableName() string {
	return "dev_tmp_tbls"
}

type ResponseUserInfo struct {
	Id string `json:"id"`
	UserInfo
	Attributes       resUserAttributes `binding:"required" json:"attributes"`
	CreatedTimestamp int               `json:"createdTimestamp"`
}

type UserInfo struct {
	Username  string `json:"username" gorm:"column:username"`
	FirstName string `json:"firstName" gorm:"column:firstname"`
	LastName  string ` json:"lastName" gorm:"column:lastname"`
	Enabled   bool   `json:"enabled" gorm:"column:enabled"`
	Email     string `json:"email" gorm:"column:email"`
}

type userAttribute struct {
	DepartmentNm []string `json:"departmentNm" gorm:"column:theme_setting_val"`
	Position     []string `json:"position" gorm:"column:theme_setting_val"`
	PhoneNumber  []string `json:"phoneNumber" gorm:"column:theme_setting_val"`
}
type resUserAttributes struct {
	DepartmentNm []string `json:"departmentNm" binding:"required"`
	Position     []string `json:"position" binding:"required"`
	PhoneNumber  []string `json:"phoneNumber" binding:"required"`
}

type userCredentilas struct {
	Type      string `json:"type" gorm:"column:credentials_type"`
	Value     string `json:"value" gorm:"column:credentials_value"`
	Temporary bool   `json:"temporary" gorm:"column:credentials_temporary"`
}

// Groups Struct
type ReqToken struct {
	Id         string          `json:"id" `
	Name       string          `json:"name" `
	Attributes groupAttributes `json:"attributes" `
}

type ResGroupInfo struct {
	Id         string          `json:"id" binding:"required"`
	Name       string          `json:"name" binding:"required"`
	Path       string          `json:"path" binding:"required"`
	Attributes groupAttributes `json:"attributes" binding:"required"`
}

type groupAttributes struct {
	TokenVal []string `json:"token"`
}
