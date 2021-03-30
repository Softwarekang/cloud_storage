package errors

import (
	"fmt"
	"strconv"
	"strings"
)

// Errors 做成导出变量, 方便类型断言到这个类型
type Errors struct {
	code  int32
	msg   string
	cause error
}

// Error 根据code生成Errors
func Error(code int32) error {
	return &Errors{code: code}
}

// Errorf code+string 生成Errors
func Errorf(code int32, format string, args ...interface{}) error {
	return &Errors{code: code, msg: fmt.Sprintf(format, args...)}
}

// Wrap error+code 生成Errors
func Wrap(cause error, code int32) error {
	if cause == nil {
		return Error(code)
	}
	return &Errors{cause: cause, code: code}
}

// Wrapf error+code+string 生成Errors
func Wrapf(cause error, code int32, format string, args ...interface{}) error {
	if cause == nil {
		return Errorf(code, format, args...)
	}
	return &Errors{cause: cause, code: code, msg: fmt.Sprintf(format, args...)}
}

// Error 获得error string
func (e *Errors) Error() string {
	sb := strings.Builder{}

	sb.WriteString("(")
	sb.WriteString(strconv.Itoa(int(e.code)))
	sb.WriteString(")")

	if len(e.msg) != 0 {
		sb.WriteString(e.msg)
	}
	if e.cause != nil {
		sb.WriteString(":")
		sb.WriteString(e.cause.Error())
	}

	return sb.String()
}

// Translate 国际化string
func (e *Errors) Translate() string {
	// TODO: 错误信息国际化还没想清楚
	return strconv.Itoa(int(e.code))
}

// Code 返回code
func (e *Errors) Code() int32 {
	return e.code
}

// Code 返回code
func (e *Errors) Message() string {
	return e.msg
}

// Cause 返回error
func (e *Errors) Cause() error {
	return e.cause
}

// Cause 返回cause
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return err
}

// Code 返回code
func Code(err error) int32 {
	type coder interface {
		Code() int32
	}

	if errors, ok := err.(coder); ok {
		return errors.Code()
	}

	return 500
}

// Message 返回msg
func Message(err error) string {
	type coder interface {
		Message() string
	}

	if errors, ok := err.(coder); ok {
		return errors.Message()
	}

	return err.Error()
}
