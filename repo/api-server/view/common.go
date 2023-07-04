package view

import "go-scribblog/repo/api-server/view/errs"

// Response is the response object send to front end
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RespWithError build a response with error e and data
func RespWithError(e errs.Error, data interface{}) *Response {
	return &Response{
		Code:    int(e.Code),
		Message: e.Msg,
		Data:    data,
	}
}

// Success build a success response with error e and data
func Success(data interface{}) *Response {
	return RespWithError(errs.Success, data)
}

// Fail build a fail response with error e and data
func Fail(data interface{}) *Response {
	return RespWithError(errs.FailedError, data)
}

// ResponseWithMsg build a fail response with ths specific message
func ResponseWithMsg(msg string) *Response {
	return RespWithError(errs.WrapError(errs.FailedError, msg), nil)
}

// RespSuccessWithMsg build a success response with ths specific message
func RespSuccessWithMsg(msg string) *Response {
	return RespWithError(errs.WrapError(errs.Success, msg), nil)
}
