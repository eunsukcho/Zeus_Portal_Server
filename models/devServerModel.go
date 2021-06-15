package models

import "time"

type DevServerModel struct {
	Dev_Server_Id uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Hostname      string `gorm:"column:hostname" json:"hostName"`
	ExternalIP    string `gorm:"column:external_ip" json:"externalIp"`
	InternalIP    string `gorm:"column:internal_ip" json:"internalIp"`
}

func (DevServerModel) TableName() string {
	return "dev_server_tbl"
}
