// core/migrate.go
package core

import (
	"gin-admin-api/global"
	"gin-admin-api/model"
	util "gin-admin-api/utils"
	"time"
)

func RunMigrations() error {
    global.Log.Info("[migrate]開始數據庫遷移")
    
    // 定義所有需要遷移的模型
    models := []any{
        &model.Menu{},
        // 其他模型可以在這裡添加
        // &model.User{},
        // &model.Role{},
    }
    
    // 執行自動遷移
    err := Db.AutoMigrate(models...)
    if err != nil {
        global.Log.Errorf("[migrate]遷移失敗: %v", err)
        return err
    }
    
    global.Log.Info("[migrate]數據庫遷移完成")
    return nil
}

func SeedData() error {
    global.Log.Info("[seed]開始添加種子數據")
    
    // 檢查是否已有數據
    var count int64
    Db.Model(&model.Menu{}).Count(&count)
    if count > 0 {
        global.Log.Info("[seed]數據已存在，跳過種子數據")
        return nil
    }
    
    // 添加初始菜單數據
    now := time.Now()
    menus := []model.Menu{
        {
            ParentId:    0,
            MenuName:    "系統管理",
            MenuIcon:    "system",
            MenuType:    1,
            MenuStatus:  1,
            Sort:        1,
            CreateTime:  util.HTime{Time: now}, // ✅ 設置當前時間
        },
        {
            ParentId:    1,
            MenuName:    "菜單管理",
            MenuIcon:    "menu",
            MenuType:    2,
            Url:         "/system/menu",
            PermissionValue: "system:menu:list",
            MenuStatus:  1,
            Sort:        1,
            CreateTime:  util.HTime{Time: now}, // ✅ 設置當前時間
        },
    }
    
    for _, menu := range menus {
        if err := Db.Create(&menu).Error; err != nil {
            global.Log.Errorf("[seed]創建菜單失敗: %v", err)
            return err
        }
    }
    
    global.Log.Info("[seed]種子數據添加完成")
    return nil
}