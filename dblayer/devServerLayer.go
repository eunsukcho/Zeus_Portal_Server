package dblayer

import "zeus/models"

type DevServerLayer interface {
	GetAllDevServerInfoData() ([]models.DevServerModel, error)
	GetDevServerInfoDataById(uint) (models.DevServerModel, error)

	SaveDevServerInfo(models.DevServerModel) (models.DevServerModel, error)

	UpdateDevServerInfo(models.DevServerModel, uint) (models.DevServerModel, error)
	DeleteDevServerInfo(uint) error
}
