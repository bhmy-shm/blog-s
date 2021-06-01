package errcode

import (
	"fmt"
	"net/http"
)

type Error struct{
	code int `json:"code"`	//状态码
	msg string `json:"msg"`	//错误提示
	details []string `json:"details"`	//详细的错误信息
}

var codes = map[int]string{}

func NewError(code int,msg string) *Error{
	if _,ok := codes[code] ; ok {
		panic(fmt.Sprintf("错误码：%d已经存在，请更换",code))
	}
	codes[code] = msg
	return &Error{code: code,msg: msg}
}

//返回简单的错误信息
func (e *Error) Error() string{
	return fmt.Sprintf("错误码：%d,错误信息：%s",e.Code(),e.Msg())
}
//只返回错误编码
func (e *Error) Code() int{
	return e.code
}
//只返回错误信息
func (e *Error) Msg() string{
	return e.msg
}
//返回多个错误信息
func (e *Error) MsgS(args []interface{}) string{
	return fmt.Sprintf(e.msg,args...)
}
//返回详细的错误信息
func (e *Error) Details() []string {
	return e.details
}
//返回多个详细错误信息的集合
func (e *Error) WithDetails(details ...string) *Error{
	e.details = []string{}
	for _,d := range details{
		e.details = append(e.details,d)
	}
	return e
}

//将自定义的错误码和http响应码做一个转换，
//不同的错误码在http响应码中有不同的含义，所以将其区分，方便抓取监控和报警系统识别/
func (e *Error)StatusCode() int{
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServiceError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.code:
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeOut.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
