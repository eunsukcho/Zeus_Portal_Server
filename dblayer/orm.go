package dblayer

import (
	"encoding/json"
	"fmt"
	"zeus/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBORM struct {
	*gorm.DB
}

func NewDBInit() (*DBORM, error) {
	dbConfig := settingDB()

	postgres_conn_name := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		dbConfig.db_host, dbConfig.db_port, dbConfig.db_username, dbConfig.db_name, dbConfig.db_password)
	fmt.Println("conname \n", postgres_conn_name)

	db, err := gorm.Open("postgres", postgres_conn_name)

	return &DBORM{
		DB: db,
	}, err
}

// init zeus env
func (db *DBORM) GetAllEnvData() (envs []models.Env_setting_Tbls, err error) {

	return envs, db.Find(&envs).Error
}
func (db *DBORM) UpdateEnvData(envs models.Env_setting_Tbls) (envInfo models.Env_setting_Tbls, err error) {

	theme := envs.ThemeSettingVal
	lang := envs.LangSettingVal
	autoLogout := envs.AutoLogoutVal
	version := envs.PortalVersion
	userAuth := envs.UserRegisterAuth

	return envInfo, db.Model(&envInfo).Updates(map[string]interface{}{"ThemeSettingVal": theme, "LangSettingVal": lang, "AutoLogoutVal": autoLogout, "PortalVersion": version, "UserRegisterAuth": userAuth}).Error
}

// menu setting
func (db *DBORM) GetTopMenuInfoByName(topCodeName string) (top models.TopMenuInfo, err error) {
	return top, db.Where("top_menu_name = ? ", topCodeName).Find(&top).Error
}
func (db *DBORM) GetAllTopMenu() (top []models.TopMenuInfo, err error) {
	return top, db.Order("top_menu_order asc").Find(&top).Error
}
func (db *DBORM) GetAllSubMenu() (sub []models.SubMenuInfo, err error) {
	return sub, db.Order("sub_menu_order asc").Find(&sub).Error
}
func (db *DBORM) CkDuplicateTopMenu(topcode string) (rst int, err error) {
	var top models.TopMenuInfo
	var cnt int
	return cnt, db.Model(&top).Where("top_menu_code = ? ", topcode).Count(&cnt).Error
}
func (db *DBORM) CkDuplicateTopMenuOrder(order int) (rst int, err error) {
	var top models.TopMenuInfo
	var cnt int
	return cnt, db.Model(&top).Where("top_menu_order = ? ", order).Count(&cnt).Error
}
func (db *DBORM) SaveTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Create(&top).Error
}
func (db *DBORM) CkDuplicateSubMenu(topcode string, subcode string) (rst int, err error) {
	var sub models.SubMenuInfo
	var cnt int
	return cnt, db.Model(&sub).Where("sub_menu_code = ? and top_menu_code = ?", subcode, topcode).Count(&cnt).Error
}
func (db *DBORM) CkDuplicateSubMenuOrder(topcode string, subcode string, order int) (rst int, err error) {
	var sub models.SubMenuInfo
	var cnt int
	return cnt, db.Model(&sub).Where("top_menu_code = ? and sub_menu_order = ? ", topcode, order).Count(&cnt).Error
}
func (db *DBORM) SaveSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Create(&sub).Error
}
func (db *DBORM) DeleteTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Where("top_menu_code = ? ", top.Top_Menu_Code).Unscoped().Delete(&top).Error
}
func (db *DBORM) DeleteSubMenuByTopCodeUrl(top string) (models.SubMenuInfo, error) {
	var sub models.SubMenuInfo
	return sub, db.Model(&sub).Where("top_menu_code=?", top).Unscoped().Delete(&sub).Error
}

func (db *DBORM) DeleteSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Where("sub_menu_code = ? and top_menu_code=?", sub.Sub_Menu_Code, sub.Top_Menu_Code).Unscoped().Delete(&sub).Error
}
func (db *DBORM) GetAllIcon() (icon []models.TopMenuIcon, err error) {
	return icon, db.Find(&icon).Error
}
func (db *DBORM) SaveUrlLink(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Model(&top).Where("top_menu_code = ?", top.Top_Menu_Code).Update(models.TopMenuInfo{Top_Menu_Target_Url: top.Top_Menu_Target_Url, New_Window: top.New_Window}).Error
}
func (db *DBORM) SaveUrlSubLink(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Model(&sub).Where("sub_menu_code = ?", sub.Sub_Menu_Code).Update(models.SubMenuInfo{Sub_Menu_Target_Url: sub.Sub_Menu_Target_Url, New_Window: sub.New_Window}).Error
}
func (db *DBORM) DeleteTopMenuUrl(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Model(&top).Where("top_menu_code = ?", top.Top_Menu_Code).Update(models.TopMenuInfo{Top_Menu_Target_Url: top.Top_Menu_Target_Url}).Error
}
func (db *DBORM) DeleteSubMenuUrl(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Model(&sub).Where("sub_menu_code = ?", sub.Sub_Menu_Code).Update(models.SubMenuInfo{Sub_Menu_Target_Url: sub.Sub_Menu_Target_Url}).Error
}
func (db *DBORM) GetMenuTargetUrl(menuCode models.SubMenuInfo) (urlCode models.SubMenuInfo, err error) {
	return urlCode, db.Where("top_menu_code=? and sub_menu_code=?", menuCode.Top_Menu_Code, menuCode.Sub_Menu_Code).Find(&urlCode).Error
}
func (db *DBORM) GetTopMenuTargetUrl(menuCode models.TopMenuInfo) (urlCode models.TopMenuInfo, err error) {
	return urlCode, db.Where("top_menu_code=?", menuCode.Top_Menu_Code).Find(&urlCode).Error
}
func (db *DBORM) UpdateTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Model(&top).Where("top_menu_code=?", top.Top_Menu_Code).Update(models.TopMenuInfo{Top_Menu_Name: top.Top_Menu_Name, Top_Menu_Order: top.Top_Menu_Order, Icon_Code: top.Icon_Code}).Error
}
func (db *DBORM) UpdateSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Model(&sub).Where("sub_menu_code=?", sub.Sub_Menu_Code).Update(models.SubMenuInfo{Sub_Menu_Name: sub.Sub_Menu_Name, Sub_Menu_Order: sub.Sub_Menu_Order, Top_Menu_Code: sub.Top_Menu_Code, Top_Menu_Name: sub.Top_Menu_Name, Icon_Code: sub.Icon_Code}).Error
}
func (db *DBORM) UpdateSubMenuTopCodeName(topCode string, topName string) (sub models.SubMenuInfo, err error) {
	return sub, db.Model(&sub).Where("top_menu_code=?", topCode).Update(models.SubMenuInfo{Top_Menu_Name: topName}).Error
}

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
