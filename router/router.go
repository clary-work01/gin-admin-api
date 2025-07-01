// 初始化路由及註冊

package router

import (
	"gin-admin-api/api"
	"gin-admin-api/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterInit() *gin.Engine{
	// 設置啟動模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	// 跌機時恢復
	router.Use(gin.Recovery())
	// 註冊
	register(router)
	return router
}
// 註冊路由接口
func register(router *gin.Engine){
	// todo後續路由接口
	router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/success",api.Success)
	router.GET("/api/failed",api.Failed)
	router.POST("/api/menu",api.CreateMenu)
}