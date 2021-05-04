package dblayer

import "zeus/models"

type SmtpLayer interface {

	//smtp setting
	SmtpInfoConnectionCheck() ([]models.SmtpInfo, error)
	SmtpInfoSave(models.SmtpInfo) (models.SmtpInfo, error)
	SmtpInfoTest() ([]models.SmtpInfo, error)
	SmtpInfoGet() ([]models.SmtpInfo, error)
}
