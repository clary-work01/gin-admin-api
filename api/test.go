// 測試api相關接口

package api

import (
	"gin-admin-api/result"

	"github.com/gin-gonic/gin"
)

// Success 成功測試
// @Tags 測試相關接口
// @Summary 成功的測試接口
// @Produce json
// @Description 成功測試接口
// @Success 200 {object} result.Result
// @router /api/success [get]
func Success(c *gin.Context){
	result.Success(c,200)
}

// Failed 失敗測試
// @Tags 測試相關接口
// @Summary 失敗的測試接口
// @Produce json
// @Description 失敗測試接口
// @Success 200 {object} result.Result
// @router /api/failed [get]
func Failed(c *gin.Context){
	result.Failed(c,int(result.ApiCode.Failed),result.ApiCode.GetMessage(result.ApiCode.Failed))
}