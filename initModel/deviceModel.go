package model

import (
	"time"
)

type Install_Device_Tbls struct {
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

func GetAll() *[]Install_Device_Tbls {
	db := DbInit()
	defer db.Close()

	var install_device_tbls []Install_Device_Tbls
	db.Find(&install_device_tbls)

	return &install_device_tbls
}

func GetDeviceByManufactureID(manufacturer string) *Install_Device_Tbls {
	db := DbInit()
	defer db.Close()

	var install_device_tbls Install_Device_Tbls
	db.Where("manufacturer = ?", manufacturer).Find(&install_device_tbls)

	return &install_device_tbls
}
