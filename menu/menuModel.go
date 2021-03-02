package menu

import "time"

type TopMenuInfo struct {
	Top_Menu_Code       string    `gorm:"column:top_menu_code" json:"top_menu_code" binding:"required`
	Top_Menu_Name       string    `gorm:"column:top_menu_name" json:"top_menu_name" binding:"required`
	Top_Menu_Target_Url string    `gorm:"column:top_menu_target_url" json:"top_menu_target_url" binding:"required`
	Top_Menu_Order      string    `gorm:"column:top_menu_order" json:"top_menu_order" binding:"required`
	Created_Dt          time.Time `gorm:"column:created_dt" json:"created_dt" binding:"required`
	Updated_Dt          time.Time `gorm:"column:updated_dt" json:"updated_dt" binding:"required`
}

type SubMenuInfo struct {
	Top_Menu_Code       string    `gorm:"column:top_menu_code;ForeignKey:top_menu_code" json:"top_menu_code" binding:"required`
	Sub_Menu_Code       string    `gorm:"column:sub_menu_code" json:"sub_menu_code" binding:"required`
	Top_Menu_Name       string    `gorm:"column:top_menu_name" json:"top_menu_name" binding:"required`
	Sub_Menu_Name       string    `gorm:"column:sub_menu_name" json:"sub_menu_name" binding:"required`
	Sub_Menu_Target_Url string    `gorm:"column:sub_menu_target_url" json:"sub_menu_target_url" binding:"required`
	Sub_Menu_Order      string    `gorm:"column:sub_menu_order" json:"sub_menu_order" binding:"required`
	Created_Dt          time.Time `gorm:"column:created_dt" json:"created_dt" binding:"required`
	Updated_Dt          time.Time `gorm:"column:updated_dt" json:"updated_dt" binding:"required`
}
