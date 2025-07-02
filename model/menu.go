package model

import util "gin-admin-api/utils"

// 菜單model

type Menu struct {
	ID uint `gorm:"column:id;comment:'主鍵';primaryKey;NOT NULL" json:"id"`
	ParentId uint `gorm:"column:parent_id;comment:'父菜單ID'" json:"parent_id"`
	MenuName string `gorm:"column:menu_name;varchar(100);comment:'菜單名稱'" json:"menu_name"`
	MenuIcon string `gorm:"column:menu_icon;varchar(100);comment:'菜單圖標'" json:"menu_icon"`
	PermissionValue string `gorm:"column:permission_value;varchar(100);comment:'權限值'" json:"permssion_value"`
	MenuType uint `gorm:"column:menu_type;comment:'菜單類型: 1->目錄 2->菜單 3->按鈕'" json:"menu_type"`
	Url string `gorm:"column:url;varchar(100);comment:'菜單URL'" json:"url"`
	MenuStatus uint `gorm:"column:menu_status;comment:'啟用狀態: 1->啟用 2->禁用'" json:"menu_status"`
	Sort uint `gorm:"column:sort;comment:'排序'" json:"sort"`
	CreateTime util.HTime  `gorm:"column:create_time;comment:'創建時間'" json:"create_time"`
	// Children []Menu  `gorm:"_" json:"children"`
	 Children        []Menu    `gorm:"foreignKey:ParentId;references:ID" json:"children"`
}

func (Menu) TableName() string{
	return "menu"
}

// CreateMenuDto 新增菜單參數
type CreateMenuDto struct{
	ParentId uint `json:"parent_id"`
	MenuName string `json:"menu_name"`
	MenuIcon string `json:"menu_icon"`
	PermissionValue string `json:"permssion_value"`
	MenuType uint `json:"menu_type"`
	Url string `json:"url"`
	MenuStatus uint `json:"menu_status"`
	Sort uint `json:"sort"`
}

// UpdateMenuDto 更新菜單參數
type UpdateMenuDto struct{
	ID uint `json:"id"`
	ParentId uint `json:"parent_id"`
	MenuName string `json:"menu_name"`
	MenuIcon string `json:"menu_icon"`
	PermissionValue string `json:"permssion_value"`
	MenuType uint `json:"menu_type"`
	Url string `json:"url"`
	MenuStatus uint `json:"menu_status"`
	Sort uint `json:"sort"`
}

type DeleteMenuDto struct{
	ID uint `json:"id"`
}