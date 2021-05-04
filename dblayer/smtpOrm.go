package dblayer

import "zeus/models"

//smtp setting
func (db *DBORM) SmtpInfoConnectionCheck() ([]models.SmtpInfo, error) {
	return nil, nil
}
func (db *DBORM) SmtpInfoSave(smtpinfo models.SmtpInfo) (models.SmtpInfo, error) {
	return smtpinfo, db.Model(&smtpinfo).Update(&smtpinfo).Error
}
func (db *DBORM) SmtpInfoTest() ([]models.SmtpInfo, error) {
	return nil, nil
}
func (db *DBORM) SmtpInfoGet() (smtpinfo []models.SmtpInfo, err error) {
	return smtpinfo, db.Find(&smtpinfo).Error
}
