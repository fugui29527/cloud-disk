package helper

var SuccessCode = "200"
var FailCode = "9999"

//token错误码
var FailAuthCode = "401"

type Result struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewSuccessResult(code string, data interface{}) *Result {
	return &Result{
		Code: code,
		Data: data,
	}
}
func NewFailResult(code string, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
	}
}
