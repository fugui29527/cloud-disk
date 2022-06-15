package helper

import "cloud-disk/core/internal/types"

var SuccessCode = "200"
var FailCode = "9999"

//数据已存在
var FailDoubleCode = "1001"

//token错误码
var FailAuthCode = "401"

func NewSuccessResult(code string, data interface{}) *types.Result {
	return &types.Result{
		Code: code,
		Data: data,
	}
}
func NewFailResult(code string, msg string) *types.Result {
	return &types.Result{
		Code: code,
		Msg:  msg,
	}
}
