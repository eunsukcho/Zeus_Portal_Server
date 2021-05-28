package dblayer

import "zeus/models"

type MenuLayer interface {

	// menu setting
	GetTopMenuInfoByName(string) (models.TopMenuInfo, error)
	GetAllTopMenu() ([]models.TopMenuInfo, error)
	GetAllSubMenu() ([]models.SubMenuInfo, error)
	CkDuplicateTopMenu(string) (int64, error)
	CkDuplicateTopMenuOrder(order int) (rst int64, err error)
	SaveTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
	CkDuplicateSubMenu(string, string) (int64, error)
	CkDuplicateSubMenuOrder(string, string, int) (rst int64, err error)
	SaveSubMenuInfo(models.SubMenuInfo) (models.SubMenuInfo, error)
	DeleteTopMenuInfo(models.TopMenuInfo) (models.TopMenuInfo, error)
	DeleteSubMenuByTopCodeUrl(string) (models.SubMenuInfo, error)
	DeleteSubMenuInfo(models.SubMenuInfo) (models.SubMenuInfo, error)
	GetAllIcon() ([]models.TopMenuIcon, error)
	SaveUrlLink(models.TopMenuInfo) (models.TopMenuInfo, error)
	SaveUrlSubLink(models.SubMenuInfo) (models.SubMenuInfo, error)
	DeleteTopMenuUrl(models.TopMenuInfo) (models.TopMenuInfo, error)
	DeleteSubMenuUrl(sub models.SubMenuInfo) (models.SubMenuInfo, error)
	GetMenuTargetUrl(models.SubMenuInfo) (models.SubMenuInfo, error)
	GetTopMenuTargetUrl(models.TopMenuInfo) (models.TopMenuInfo, error)
	UpdateTopMenuInfo(top models.TopMenuInfo) (models.TopMenuInfo, error)
	UpdateSubMenuInfo(sub models.SubMenuInfo) (models.SubMenuInfo, error)
	UpdateSubMenuTopCodeName(string, string) (models.SubMenuInfo, error)
}
