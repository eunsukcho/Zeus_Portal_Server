package models

type AdminAPIInfo struct {
	User RegisterUserInfo `binding:"required" json:"user"`
	Admin Authdetails `binding:"required" json:"admin"`
}

type RegisterUserInfo struct {
	UserInfo
	Attributes userAttribute `binding:"required" json:"attributes"`
	Credential []userCredentilas `binding:"required" json:"credentials"`
}

type ResponseUserInfo struct {
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
	ClientRoles string 		 `binding:"required" json:"clientRoles"`
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

type Authdetails struct {
	ClientId    string `json:"clientId" binding:"required"`
	ClientSecret    string `json:"clientSecret" binding:"required"`
	AdminId     string `json:"adminId" binding:"required"`
	AdminPw     string `json:"adminPw" binding:"required"`
	TokenUrl       string `json:"tokenUrl" binding:"required"`
}
