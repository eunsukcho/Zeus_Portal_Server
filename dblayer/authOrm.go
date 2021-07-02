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

	fmt.Println("GetDevUserInfo : ", tmp)
	for _, info := range tmp {
		var registerInfo = models.RegisterUserInfo{}
		err := json.Unmarshal([]byte(info.Dev_info), &registerInfo)
		if err != nil {
			fmt.Println(err)
		}
		registerInfo.Dev_User_Id = info.Dev_User_Id
		devInfo = append(devInfo, registerInfo)
	}

	fmt.Println("devInfo : ", devInfo)
	return devInfo, nil
}
func (db *DBORM) AcceptUpdateUser(user uint) (dev models.Dev_Info, err error) {
	fmt.Println("AcceptUpdateUser : ", user)
	var updateTbl models.Dev_Info
	return dev, db.Model(&updateTbl).Where("dev_user_id = ?", user).Updates(models.Dev_Info{Enabled: true}).Error
}

func (db *DBORM) DeleteUser(user string, id uint) error {
	var deleteTbl models.Dev_Info
	return db.Model(&deleteTbl).Where("email = ? and dev_user_id=?", user, id).Unscoped().Delete(&deleteTbl).Error
}

func (db *DBORM) CkDuplicateTmpDev(user string) (rst int64, err error) {
	var target models.Dev_Info
	var cnt int64
	return cnt, db.Model(&target).Where("email = ? and enabled=?", user, false).Count(&cnt).Error
}
