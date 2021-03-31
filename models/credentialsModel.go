package models

type Authdetails struct {
	ClientId     string `gorm:"column:client_id" json:"clientId" binding:"required"`
	ClientSecret string `gorm:"column:client_secret" json:"clientSecret" binding:"required"`
	AdminId      string `gorm:"column:admin_id" json:"adminId" binding:"required"`
	AdminPw      string `gorm:"column:admin_pw" json:"adminPw" binding:"required"`
	TokenUrl     string `gorm:"column:token_url" json:"tokenUrl" binding:"required"`
}

func (Authdetails) TableName() string {
	return "admin_auth_tbls"
}

// Binding struct
type AdminAPIInfo struct {
	User  RegisterUserInfo `json:"user"`
	Admin Authdetails      `binding:"required" json:"admin"`
}
type GroupAdminAPIInfo struct {
	Groups ReqToken    `json:"groups"`
	Admin  Authdetails `binding:"required" json:"admin"`
}

// Binding Uri
type Uri struct {
	Id string `uri:"id" binding:"required"`
}

// User Struct
type UserListData struct {
	Id    string      `json:"id"`
	Admin Authdetails `binding:"required" json:"admin"`
}

type RegisterUserInfo struct {
	ID string `json:"id"`
	UserInfo
	Groups     []string          `json:"groups"`
	Attributes userAttribute     `json:"attributes"`
	Credential []userCredentilas `json:"credentials"`
}

type ResponseUserInfo struct {
	Id string `json:"id"`
	UserInfo
	Attributes       resUserAttributes `binding:"required" json:"attributes"`
	CreatedTimestamp int               `json:"createdTimestamp"`
}

type UserInfo struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string ` json:"lastName"`
	Enabled   bool `json:"enabled"`
	Email     string `json:"email"`
}

type userAttribute struct {
	DepartmentNm []string `json:"departmentNm"`
	Position     []string `json:"position"`
	PhoneNumber  []string `json:"phoneNumber" `
}
type resUserAttributes struct {
	DepartmentNm []string `json:"departmentNm" binding:"required"`
	Position     []string `json:"position" binding:"required"`
	PhoneNumber  []string `json:"phoneNumber" binding:"required"`
}

type userCredentilas struct {
	Type      string `json:"type"`
	Value     string `json:"value" `
	Temporary bool   `json:"temporary" `
}

// Groups Struct
type ReqToken struct {
	Id         string          `json:"id" `
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
