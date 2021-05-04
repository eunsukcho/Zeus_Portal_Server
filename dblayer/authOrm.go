package dblayer

import (
	"encoding/json"
	"fmt"
	"zeus/models"
)

func (db *DBORM) GetAllAuthData() (auth []models.Authdetails, err error) {

	return auth, db.Find(&auth).Error
}
func (db *DBORM) SaveAuthData(authBinding models.Authdetails) (auth []models.Authdetails, err error) {
	err = db.Create(&authBinding).Error
	if err != nil {
		return nil, err
	}

	return auth, db.Find(&auth).Error
}

func (db *DBORM) SaveDevUserInfo(devinfo models.Dev_Info) (models.Dev_Info, error) {

	return devinfo, db.Create(&devinfo).Error
}

func (db *DBORM) GetDevUserInfo(group string) ([]models.RegisterUserInfo, error) {
	var tmp []models.Dev_Info
	var devInfo []models.RegisterUserInfo
	err := db.Where("groupname=? and enabled=?", group, false).Find(&tmp).Error
	if err != nil {
		panic(err)
	}

	for _, info := range tmp {
		var registerInfo = models.RegisterUserInfo{}
		err := json.Unmarshal([]byte(info.Dev_info), &registerInfo)
		if err != nil {
			fmt.Println(err)
		}
		devInfo = append(devInfo, registerInfo)
	}

	fmt.Println("devInfo : ", devInfo)
	return devInfo, nil
}
func (db *DBORM) AcceptUpdateUser(user string) (dev models.Dev_Info, err error) {

	return dev, db.Model(&dev).Where("email = ?", user).Update(models.Dev_Info{Enabled: true}).Error
}
