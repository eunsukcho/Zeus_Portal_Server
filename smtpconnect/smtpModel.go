package smtpconnect

import "github.com/jinzhu/gorm"

type SmtpInfo struct {
	gorm.Model
	AdminAddress string `gorm:"column:smtp_user" json:"AdminAddress" binding:"required"`
	SmtpAddress  string `gorm:"column:smtp_host" json:"SmtpAddress" binding:"required"`
	Port         string `gorm:"column:smtp_port" json:"Port" binding:"required"`
	Password     string `gorm:"column:smtp_password" json:"Password" binding:"required"`
}
