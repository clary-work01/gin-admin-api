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
	if err != nil{
		return err
	}
	if db.Error != nil {
		return err
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDb.SetMaxOpenConns(dbConfig.MaxOpen)
	global.Log.Infof("[mysql]連接成功")
	return nil
}