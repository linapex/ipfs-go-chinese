
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:38</date>
//</624460159820763136>

package e

import (
	"fmt"
	"runtime/debug"
)

//typeerr返回一个错误，其字符串解释了预期的错误和收到的错误。
func TypeErr(expected, actual interface{}) error {
	return fmt.Errorf("expected type %T, got %T", expected, actual)
}

//编译时类型检查handlerError是否为错误
var _ error = New(nil)

//handlerError向错误添加堆栈跟踪
type HandlerError struct {
	Err   error
	Stack []byte
}

//错误使handlerError实现错误
func (err HandlerError) Error() string {
	return fmt.Sprintf("%s in:\n%s", err.Err.Error(), err.Stack)
}

//new返回新的handlerError
func New(err error) HandlerError {
	return HandlerError{Err: err, Stack: debug.Stack()}
}

