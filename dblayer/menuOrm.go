package dblayer

import "zeus/models"

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
func (db *DBORM) CkDuplicateTopMenu(topcode string) (rst int64, err error) {
	var top models.TopMenuInfo
	var cnt int64
	return cnt, db.Model(&top).Where("top_menu_code = ? ", topcode).Count(&cnt).Error
}
func (db *DBORM) CkDuplicateTopMenuOrder(order int) (rst int64, err error) {
	var top models.TopMenuInfo
	var cnt int64
	return cnt, db.Model(&top).Where("top_menu_order = ? ", order).Count(&cnt).Error
}
func (db *DBORM) SaveTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Create(&top).Error
}
func (db *DBORM) CkDuplicateSubMenu(topcode string, subcode string) (rst int64, err error) {
	var sub models.SubMenuInfo
	var cnt int64
	return cnt, db.Model(&sub).Where("sub_menu_code = ? and top_menu_code = ?", subcode, topcode).Count(&cnt).Error
}
func (db *DBORM) CkDuplicateSubMenuOrder(topcode string, subcode string, order int) (rst int64, err error) {
	var sub models.SubMenuInfo
	var cnt int64
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
	return top, db.Model(&top).Where("top_menu_code = ?", top.Top_Menu_Code).Updates(models.TopMenuInfo{Top_Menu_Target_Url: top.Top_Menu_Target_Url, New_Window: top.New_Window}).Error
}
func (db *DBORM) SaveUrlSubLink(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Model(&sub).Where("sub_menu_code = ?", sub.Sub_Menu_Code).Updates(models.SubMenuInfo{Sub_Menu_Target_Url: sub.Sub_Menu_Target_Url, New_Window: sub.New_Window}).Error
}
func (db *DBORM) DeleteTopMenuUrl(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Model(&top).Where("top_menu_code = ?", top.Top_Menu_Code).Updates(models.TopMenuInfo{Top_Menu_Target_Url: top.Top_Menu_Target_Url}).Error
}
func (db *DBORM) DeleteSubMenuUrl(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Model(&sub).Where("sub_menu_code = ?", sub.Sub_Menu_Code).Updates(models.SubMenuInfo{Sub_Menu_Target_Url: sub.Sub_Menu_Target_Url}).Error
}
func (db *DBORM) GetMenuTargetUrl(menuCode models.SubMenuInfo) (urlCode models.SubMenuInfo, err error) {
	return urlCode, db.Where("top_menu_code=? and sub_menu_code=?", menuCode.Top_Menu_Code, menuCode.Sub_Menu_Code).Find(&urlCode).Error
}
func (db *DBORM) GetTopMenuTargetUrl(menuCode models.TopMenuInfo) (urlCode models.TopMenuInfo, err error) {
	return urlCode, db.Where("top_menu_code=?", menuCode.Top_Menu_Code).Find(&urlCode).Error
}
func (db *DBORM) UpdateTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error) {
	return top, db.Model(&top).Where("top_menu_code=?", top.Top_Menu_Code).Updates(models.TopMenuInfo{Top_Menu_Name: top.Top_Menu_Name, Top_Menu_Order: top.Top_Menu_Order, Icon_Code: top.Icon_Code}).Error
}
func (db *DBORM) UpdateSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error) {
	return sub, db.Model(&sub).Where("sub_menu_code=?", sub.Sub_Menu_Code).Updates(models.SubMenuInfo{Sub_Menu_Name: sub.Sub_Menu_Name, Sub_Menu_Order: sub.Sub_Menu_Order, Top_Menu_Code: sub.Top_Menu_Code, Top_Menu_Name: sub.Top_Menu_Name, Icon_Code: sub.Icon_Code}).Error
}
func (db *DBORM) UpdateSubMenuTopCodeName(topCode string, topName string) (sub models.SubMenuInfo, err error) {
	return sub, db.Model(&sub).Where("top_menu_code=?", topCode).Updates(models.SubMenuInfo{Top_Menu_Name: topName}).Error
}
