package model

type RoleMenu struct{
	RoleId uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"role_id"`
	MenuId uint `gorm:"column:menu_id;comment:'菜單id';NOT NULL" json:"menu_id"`
}

func (RoleMenu) TableName() string{
	return "role_menu"
}	