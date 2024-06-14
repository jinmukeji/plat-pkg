package errors

import (
	"fmt"

	"github.com/jinmukeji/plat-pkg/v4/micro/errors/codes"
)

type RpcError struct {
	// 错误码
	Code codes.Code `json:"code"`

	// 额外的错误提示消息
	Message string `json:"message"`

	// 导致本 error 的内部 error
	Cause error `json:"cause"`
}

func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	if re, ok := err.(*RpcError); ok {
		return re.GetCode()
	}
	return codes.Unknown
}

func New(c codes.Code, msg string) *RpcError {
	return &RpcError{
		Code:    c,
		Message: msg,
	}
}

func Newf(c codes.Code, format string, a ...interface{}) *RpcError {
	return New(c, fmt.Sprintf(format, a...))
}

func Error(c codes.Code, msg string) error {
	return New(c, msg)
}

func Errorf(c codes.Code, format string, a ...interface{}) error {
	return Newf(c, format, a...)
}

func ErrorWithCause(c codes.Code, cause error, msg string) error {
	re := New(c, msg)
	return re.WithCause(cause)
}

func ErrorfWithCause(c codes.Code, cause error, format string, a ...interface{}) error {
	re := Newf(c, format, a...)
	return re.WithCause(cause)
}

func (e *RpcError) WithCause(err error) *RpcError {
	if e != nil {
		e.Cause = err
	}
	return e
}

func (e *RpcError) leading() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("[errcode: %d] %s", e.Code, e.Code.Message())
}

func (e *RpcError) Error() string {
	if e == nil {
		return ""
	}

	if len(e.Message) > 0 {
		return fmt.Sprintf("%s: %s", e.leading(), e.Message)
	}

	return e.leading()
}

func (e *RpcError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Cause
}

func (e *RpcError) DetailedError() string {
	if e == nil {
		return ""
	}

	if e.Cause != nil {
		return fmt.Sprintf("%s ╭∩╮ %v", e.Error(), e.Cause)
	}

	return e.Error()
}

func (e *RpcError) GetCode() codes.Code {
	if e == nil {
		return codes.OK
	}

	return e.Code
}

func (e *RpcError) GetMessage() string {
	if e == nil {
		return ""
	}

	return e.Message
}

func (e *RpcError) GetCause() error {
	if e == nil {
		return nil
	}

	return e.Cause
}
