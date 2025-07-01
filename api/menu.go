// 菜單接口
package api

import (
	"fmt"
	"gin-admin-api/core"
	"gin-admin-api/model"
	"gin-admin-api/result"
	util "gin-admin-api/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 新增菜單
// @Tags 菜單相關接口
// @Summary 新增菜單接口
// @Produce json
// @Description 新增菜單接口
// @Param data body model.CreateMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/menu [post]
func CreateMenu(c *gin.Context){
	// 綁定參數
	var dto model.CreateMenuDto
	_ = c.BindJSON(&dto)

	// 查詢菜單不能重複
	menuByName := GetMenuByName(dto.MenuName)
	fmt.Println("menuByName",menuByName)
	if menuByName.ID != 0 {
		result.Failed(c, int(result.ApiCode.MenuIsExists),result.ApiCode.GetMessage(result.ApiCode.MenuIsExists))
		return 
	}

	// 目錄
	if dto.MenuType == 1 {
		menu := model.Menu{
			ParentId:0,
			MenuName: dto.MenuName,
			MenuIcon: dto.MenuIcon,
			MenuType: dto.MenuType,
			MenuStatus: dto.MenuStatus,
			Sort: dto.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		core.Db.Create(&menu)
	// 菜單
	}else if dto.MenuType == 2{
		menu := model.Menu{
			ParentId: dto.ParentId,
			MenuName: dto.MenuName,
			MenuIcon: dto.MenuIcon,
			MenuType: dto.MenuType,
			Url: dto.Url,
			PermissionValue: dto.PermissionValue,
			MenuStatus: dto.MenuStatus,
			Sort: dto.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		core.Db.Create(&menu)
	// 按鈕
	}else if dto.MenuType == 3{
		menu := model.Menu{
			ParentId: dto.ParentId,
			MenuName: dto.MenuName,
			MenuType: dto.MenuType,
			PermissionValue: dto.PermissionValue,
			MenuStatus: dto.MenuStatus,
			Sort: dto.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		core.Db.Create(&menu)
	}
	result.Success(c,true)

}

func GetMenuByName(menuName string) (menu model.Menu){
	core.Db.Where("menu_name = ?",menuName).First(&menu)
	return menu
}