package dblayer

import "zeus/models"

func (db *DBORM) GetAllDevServerInfoData() (dev []models.DevServerModel, err error) {

	return dev, db.Find(&dev).Error
}

func (db *DBORM) GetDevServerInfoDataById(id uint) (dev models.DevServerModel, err error) {
	return dev, db.Where("id=?", id).Find(&dev).Error
}

func (db *DBORM) SaveDevServerInfo(dev models.DevServerModel) (models.DevServerModel, error) {
	return dev, db.Create(&dev).Error
}

func (db *DBORM) UpdateDevServerInfo(dev models.DevServerModel, id uint) (models.DevServerModel, error) {
	var updateTbl models.DevServerModel
	return dev, db.Model(&updateTbl).Where("id=?", id).Update(dev).Error
}

func (db *DBORM) DeleteDevServerInfo(id uint) error {
	var dev models.DevServerModel
	return db.Where("id=?", id).Unscoped().Delete(&dev).Error
}
