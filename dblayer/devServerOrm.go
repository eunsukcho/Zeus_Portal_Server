package dblayer

import "zeus/models"

func (db *DBORM) GetAllDevServerInfoData() (dev []models.DevServerModel, err error) {

	return dev, db.Order("dev_server_id desc").Find(&dev).Error
}

func (db *DBORM) GetDevServerInfoDataById(id uint) (dev models.DevServerModel, err error) {
	return dev, db.Where("dev_server_id=?", id).Find(&dev).Error
}

func (db *DBORM) SaveDevServerInfo(dev models.DevServerModel) (models.DevServerModel, error) {
	return dev, db.Create(&dev).Error
}

func (db *DBORM) UpdateDevServerInfo(dev models.DevServerModel, id uint) (models.DevServerModel, error) {
	var updateTbl models.DevServerModel
	return dev, db.Model(&updateTbl).Where("dev_server_id=?", id).Updates(models.DevServerModel{InternalIP: dev.InternalIP, ExternalIP: dev.ExternalIP, Hostname: dev.Hostname}).Error
}

func (db *DBORM) DeleteDevServerInfo(id uint) error {
	var dev models.DevServerModel
	return db.Where("dev_server_id=?", id).Unscoped().Delete(&dev).Error
}
