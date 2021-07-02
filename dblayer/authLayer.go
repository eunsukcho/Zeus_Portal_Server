package dblayer

import "zeus/models"

type AuthLayer interface {
	//auth setting
	GetAllAuthData() ([]models.Authdetails, error)
	SaveAuthData(models.Authdetails) ([]models.Authdetails, error)

	SaveDevUserInfo(models.Dev_Info) (models.Dev_Info, error)
	GetDevUserInfo(string) ([]models.RegisterUserInfo, error)
	AcceptUpdateUser(uint) (models.Dev_Info, error)
	DeleteUser(string, uint) error
	CkDuplicateTmpDev(string) (int64, error)
}
