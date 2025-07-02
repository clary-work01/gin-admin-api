package core

import (
	"fmt"
	"gin-admin-api/config"
	"gin-admin-api/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// mysql配置初始化

var Db *gorm.DB

func MysqlInit() error{
	var dbConfig = config.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
	dbConfig.Username,
	dbConfig.Password,
	dbConfig.Host,
	dbConfig.Port,
	dbConfig.Db,
	dbConfig.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating:true,
	})
	if err != nil {
        global.Log.Errorf("[mysql]連接失敗: %v", err)
        return err
    }
    if db.Error != nil {
        global.Log.Errorf("[mysql]數據庫錯誤: %v", db.Error)
        return db.Error
    }
	Db = db

    sqlDb, err := db.DB()
    if err != nil {
        global.Log.Errorf("[mysql]獲取SQL DB失敗: %v", err)
        return err
    }
    
    sqlDb.SetMaxIdleConns(dbConfig.MaxIdle)
    sqlDb.SetMaxOpenConns(dbConfig.MaxOpen)
    
    global.Log.Infof("[mysql]連接成功")

	// 運行遷移
	if err := RunMigrations(); err != nil {
		panic("數據庫遷移失敗: " + err.Error())
	}
		
	// 可選：添加種子數據
	if err := SeedData(); err != nil {
		global.Log.Errorf("種子數據添加失敗: %v", err)
	}
	return nil
}