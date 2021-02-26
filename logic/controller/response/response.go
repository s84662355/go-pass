package response

// 请求返回值
const (
	CodeSuccess   = 0 // 成功返回
	CodeFail      = 1
	CodeFailRetry = 2 //需要改参数重试
)

// Response 用户响应数据
type RespData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//请求成功，返回
func Success(data interface{}) RespData {
	return RespData{
		Code:    CodeSuccess,
		Message: "成功",
		Data:    data,
	}
}

//请求成功，返回
func Fail(msg string) RespData {
	return RespData{
		Code:    CodeFail,
		Message: msg,
		Data:    nil,
	}
}

//失败
func FailCode(msg string, code int) RespData {
	return RespData{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}
