package result

// 狀態碼和狀態資料定義

type Codes struct{
	Message map[uint]string
	Success uint
	Failed uint
	MenuIsExists uint
}

var ApiCode = &Codes{
	Success: 200,
	Failed: 501,
	MenuIsExists: 600,
}

// 狀態資料初始化
func init(){
	ApiCode.Message = map[uint]string{
		ApiCode.Success:"成功",
		ApiCode.Failed:"失敗",
		ApiCode.MenuIsExists :"菜單名稱已存在,請重新輸入",
	}
}

// getMessage 提供外部調用
func (c *Codes) GetMessage(code uint) string{
	message ,ok := c.Message[code]
	if !ok{
		return ""
	}
	return message
}