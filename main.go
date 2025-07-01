package main

import (
	"fmt"
	"gin-admin-api/config"
	"gin-admin-api/core"
	_ "gin-admin-api/docs"
	"gin-admin-api/global"
	"gin-admin-api/router"
)

// @title admin-api
// @version 1.0
// @description admin-api文件
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main(){
	// 初始化 logger
	global.Log = core.InitLogger()
	// 初始化 mysql
	core.MysqlInit()
	// 初始化 redis
	core.RedisInit()
	// 初始化 router
	router := router.RouterInit()
	address := fmt.Sprintf("%s:%s",config.Config.System.Host,config.Config.System.Port)
	global.Log.Infof("系統啟動成功,運行在：%s",address)
	router.Run(address)
}
