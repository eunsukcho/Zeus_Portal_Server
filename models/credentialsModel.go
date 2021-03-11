package models

import "golang.org/x/oauth2"

type RegisterUserInfo struct {
	Username   string        `binding:"required" json:"username"`
	FirstName  string        `binding:"required" json:"firstName"`
	LastName   string        `binding:"required" json:"lastName"`
	Enabled    string        `binding:"required" json:"enabled"`
	Email      string        `binding:"required" json:"email"`
	Attributes userAttribute `binding:"required" json:"attributes"`
}

type ResponseUserInfo struct {
	RegisterUserInfo
	CreatedTimestamp int `json:"CreatedTimestamp"`
}

type userAttribute struct {
	DepartmentNm string `json:"departmentNm" binding:"required"`
	Position     string `json:"position" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
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
