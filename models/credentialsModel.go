package models

import "golang.org/x/oauth2"

type UserInfo struct {
	Username   string        `binding:"required" json:"username"`
	FirstName  string        `binding:"required" json:"firstName"`
	LastName   string        `binding:"required" json:"lastName"`
	Enabled    string        `binding:"required" json:"enabled"`
	Email      string        `binding:"required" json:"email"`
	Attributes userAttribute `binding:"required" json:"attributes"`
}

type RegisterUserInfo struct {
	UserInfo
	Credential []userCredentilas `binding:"required" json:"credentials"`
}

type ResponseUserInfo struct {
	UserInfo
	CreatedTimestamp int `json:"CreatedTimestamp"`
}

type userAttribute struct {
	DepartmentNm string `json:"departmentNm" binding:"required"`
	Position     string `json:"position" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
}

type userCredentilas struct {
	Type      string `json:"type" binding:"required"`
	Value     string `json:"value" binding:"required"`
	Temporary bool   `json:"temporary" binding:"required"`
}

type Authdetails struct {
	APIClient    string
	APISecret    string
	UserName     string
	Password     string
	Account      string
	APIURL       string
	OrbitURL     string
	CurrentToken oauth2.TokenSource
	Transport    oauth2.Transport
}
