package errs

import "fmt"

// ErrCode is the error code send to front end
type ErrCode int

const (
	successCode ErrCode = 0
	successMsg  string  = "success"
	failCode    ErrCode = 1
	failMsg     string  = "fail"
)

type Error struct {
	Code ErrCode
	Msg  string
}

// Error() return the format string of error
func (e Error) Error() string {
	return fmt.Sprintf("[%d]%s", e.Code, e.Msg)
}

func newError(code ErrCode, msg string) Error {
	return Error{
		Code: code,
		Msg:  msg,
	}
}

// NewError return an error with code failCode and message failMsg
func NewError(msg string) Error {
	return Error{
		Code: failCode,
		Msg:  fmt.Sprintf("%s: %s", failMsg, msg),
	}
}

// WrapError return an error with the same code as input error e, and it's message is msg appended to e's message
func WrapError(e Error, msg string) Error {
	e.Msg = fmt.Sprintf("%s: %s", e.Msg, msg)
	return e
}

// 基本错误
var (
	Success     = newError(successCode, successMsg)
	FailedError = newError(failCode, failMsg)
)

// 请求相关错误
var (
	RequestParseError = newError(1000, "failed to parse the request")
	RequestParamError = newError(1001, "illegal param")
)

// db相关错误
var (
	DBOperationError   = newError(2000, "failed to operate db")
	DBRecordExistError = newError(2001, "found duplicated key")
	DBNotFoundError    = newError(2002, "can not found the record in db")
)
