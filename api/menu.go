// 菜單接口
package api

import (
	"gin-admin-api/core"
	"gin-admin-api/global"
	"gin-admin-api/model"
	"gin-admin-api/result"
	util "gin-admin-api/utils"
	"strconv"
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

	// 菜單名稱不能重複
	menuByName := GetMenuByName(dto.MenuName)
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
		// {
		// 	"menu_icon": "test",
		// 	"menu_name": "目錄",
		// 	"menu_status": 1,
		// 	"menu_type": 1,
		// 	"parent_id": 0,
		// 	"sort": 0
		//   }
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
		// {
		// 	"menu_icon": "菜單icon",
		// 	"menu_name": "菜單",
		// 	"menu_status": 1,
		// 	"menu_type": 2,
		// 	"parent_id": 1,
		// 	"permssion_value": "/menu/test",
		// 	"sort": 1,
		// 	"url": "url"
		//   }
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
		// {
		// 	"menu_icon": "按鈕icon",
		// 	"menu_name": "按鈕",
		// 	"menu_status": 1,
		// 	"menu_type": 3,
		// 	"parent_id": 2,
		// 	"permssion_value": "/menu/test/test",
		// 	"sort": 1
		//   }
		core.Db.Create(&menu)
	}
	result.Success(c,true)
}

// 查詢菜單列表
// @Tags 菜單相關接口
// @Summary 查詢菜單列表接口
// @Produce json
// @Description 查詢菜單列表接口
// @Param menu_name query string false "菜單名稱"
// @Param menu_status query integer false "菜單狀態"
// @Success 200 {object} result.Result
// @router /api/menu/list [get]
func GetMenuList (c *gin.Context){
	menuName := c.Query("menu_name")
	menuStatus := c.Query("menu_status")

	global.Log.Info("hello"+menuName,menuStatus)

	var menuList []model.Menu
	curDb := core.Db.Table("menu").Order("sort")

	if menuName != ""{
		curDb = curDb.Where("menu_name = ?",menuName)
	}

	if menuStatus != ""{
		curDb = curDb.Where("menu_status = ?",menuStatus)
	}

	curDb.Find(&menuList)
	result.Success(c,menuList)
}

// 根據id查詢菜單
// @Tags 菜單相關接口
// @Summary 根據id查詢菜單接口
// @Produce json
// @Description 根據id查詢菜單接口
// @Param id query int true "菜單id"
// @Success 200 {object} result.Result
// @router /api/menu [get]
func GetMenuById (c *gin.Context){
	id,_ := strconv.Atoi(c.Query("id"))
	var menu model.Menu
	core.Db.First(&menu,id)
	result.Success(c,menu)
}

// 根據id更新菜單
// @Tags 菜單相關接口
// @Summary 根據id更新菜單接口
// @Produce json
// @Description 根據id更新菜單接口
// @Param data body model.UpdateMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/menu [put]
func UpdateMenu(c *gin.Context){
	var dto model.UpdateMenuDto
    _ = c.ShouldBind(&dto)

	var menu model.Menu
	core.Db.First(&menu,dto.ID)

	global.Log.Info(dto)

	menu.MenuIcon = dto.MenuIcon
	menu.MenuName = dto.MenuName	
	menu.ParentId = dto.ParentId
	menu.PermissionValue = dto.PermissionValue
	menu.MenuType = dto.MenuType
	menu.Url = dto.Url
	menu.MenuStatus = dto.MenuStatus
	menu.Sort = dto.Sort
	global.Log.Info(menu)

	// **重要：保存到資料庫**
	if err := core.Db.Save(&menu).Error; err != nil {
		global.Log.Error("更新菜單失敗:", err)
		return
	}
	result.Success(c,true)
}

// 根據id刪除菜單
// @Tags 菜單相關接口
// @Summary  根據id刪除菜單接口
// @Produce json
// @Description  根據id刪除菜單接口
// @Param data body model.DeleteMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/menu [delete]
func DeleteMenu(c *gin.Context){
	var dto model.DeleteMenuDto
	_ = c.ShouldBind(&dto)
	// 若菜單已分配 不能刪除
	roleMenu := GetRoleMenuById(dto.ID)
	if roleMenu.MenuId >0 {
		result.Failed()
	}
	core.Db.Delete(dto.ID)

}


func GetMenuByName(menuName string) (menu model.Menu) {
    global.Log.Infof("查詢菜單: %s", menuName)
    
    err := core.Db.Where("menu_name = ?", menuName).First(&menu).Error
    if err != nil {
		global.Log.Errorf("查詢菜單失敗: %v", err)
        return menu // 返回零值
    }
    
    return menu
}

func GetRoleMenuById(id uint)(roleMenu model.RoleMenu){
	core.Db.Where("menu = ?",id).First(&roleMenu)
	return roleMenu
} 