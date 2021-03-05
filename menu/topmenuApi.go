package menu

import (
	"net/http"
	model "zeus/initModel"

	"github.com/gin-gonic/gin"
)

func (TopMenuInfo) TableName() string {
	return "top_menu_tbl"
}

func (SubMenuInfo) TableName() string {
	return "sub_menu_tbl"
}

func GetTopMenu() *[]TopMenuInfo {
	var topmenuinfo []TopMenuInfo

	db := model.DbInit()
	defer db.Close()
	db.Order("top_menu_order asc").Find(&topmenuinfo)

	return &topmenuinfo
}

func GetSubMenu() *[]SubMenuInfo {
	var submenuinfo []SubMenuInfo

	db := model.DbInit()
	defer db.Close()
	db.Order("sub_menu_order asc").Find(&submenuinfo)

	return &submenuinfo
}

func SaveTopMenu(c *gin.Context) {
	var topmenuinfo TopMenuInfo
	db := model.DbInit()
	defer db.Close()
	err := c.BindJSON(&topmenuinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	db.Create(&topmenuinfo)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   topmenuinfo,
	})
}

func SaveSubMenu(c *gin.Context) {
	var submenuinfo SubMenuInfo
	db := model.DbInit()
	defer db.Close()
	err := c.BindJSON(&submenuinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	db.Create(&submenuinfo)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   submenuinfo,
	})

}

func DeleteTopMenu(c *gin.Context) {
	var topmenuinfo TopMenuInfo
	db := model.DbInit()
	defer db.Close()
	err := c.BindJSON(&topmenuinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	db.Where("top_menu_code = ? ", topmenuinfo.Top_Menu_Code).Delete(&topmenuinfo)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   topmenuinfo,
	})
}

func DeleteSubMenu(c *gin.Context) {
	var submenuinfo SubMenuInfo
	db := model.DbInit()
	defer db.Close()
	err := c.BindJSON(&submenuinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	db.Where("sub_menu_code = ? ", submenuinfo.Sub_Menu_Code).Delete(&submenuinfo)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   submenuinfo,
	})
}

func GetTopMenuData(c *gin.Context) {
	topmenuinfo := GetTopMenu()
	c.JSON(http.StatusOK, topmenuinfo)
}

func SubTopMenuData(c *gin.Context) {
	submenuinfo := GetSubMenu()
	c.JSON(http.StatusOK, submenuinfo)
}
