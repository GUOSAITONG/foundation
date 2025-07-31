package response

// 标准返回结构（三要素）
type Result struct {
	Data    interface{} `json:"data"`    // 返回数据
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 错误信息
}

// 成功返回
func Success(data interface{}) Result {
	return Result{
		Data:    data,
		Code:    200,
		Message: "success",
	}
}

// 错误返回
func Error(code int, msg string) Result {
	return Result{
		Data:    nil,
		Code:    code,
		Message: msg,
	}
}
