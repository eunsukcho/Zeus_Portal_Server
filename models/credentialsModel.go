package models

type AdminAPIInfo struct {
	User RegisterUserInfo `binding:"required" json:"user"`
	Admin Authdetails `binding:"required" json:"admin"`
}

type Authdetails struct {
	ClientId    string `json:"clientId" binding:"required"`
	ClientSecret    string `json:"clientSecret" binding:"required"`
	AdminId     string `json:"adminId" binding:"required"`
	AdminPw     string `json:"adminPw" binding:"required"`
	TokenUrl       string `json:"tokenUrl" binding:"required"`
}

type RegisterUserInfo struct {
	UserInfo
	Groups []string `binding:"required" json:"groups"`
	Attributes userAttribute `binding:"required" json:"attributes"`
	Credential []userCredentilas `binding:"required" json:"credentials"`
}

type ResponseUserInfo struct {
	Id		   		string		 `json:"id"`
	UserInfo
	Attributes resUserAttributes `binding:"required" json:"attributes"`
	CreatedTimestamp int `json:"CreatedTimestamp"`
}

type UserInfo struct {
	Username   string        `binding:"required" json:"username"`
	FirstName  string        `binding:"required" json:"firstName"`
	LastName   string        `binding:"required" json:"lastName"`
	Enabled    string        `binding:"required" json:"enabled"`
	Email      string        `binding:"required" json:"email"`
}

type userAttribute struct {
	DepartmentNm string `json:"departmentNm" binding:"required"`
	Position     string `json:"position" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
}
type resUserAttributes struct {
	DepartmentNm []string `json:"departmentNm" binding:"required"`
	Position     []string `json:"position" binding:"required"`
	PhoneNumber  []string `json:"phoneNumber" binding:"required"`
}

type userCredentilas struct {
	Type      string `json:"type" binding:"required"`
	Value     string `json:"value" binding:"required"`
	Temporary bool   `json:"temporary" binding:"required"`
}

type ReqToken struct {
	Id string `json:"id" `
	Attributes groupAttributes `json:"attributes" `
}

type ResGroupInfo struct {
	Id string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
	Attributes groupAttributes `json:"attributes" binding:"required"`
}

type groupAttributes struct {
	TokenVal []string `json:"token"`
}