// 請求回應結構定義
package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data any `json:"data"`
}

func Success(c *gin.Context, data any){
	if data == nil{
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.Success)
	res.Message = ApiCode.GetMessage(ApiCode.Success)
	res.Data = data
	c.JSON(http.StatusOK,res)
}

func Failed(c *gin.Context, code int, message string){
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK,res)
}
