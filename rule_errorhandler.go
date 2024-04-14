package vrule

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type ErrorHandler interface {
	ErrorHandler(v, v1 any) error
}

type staticErrorHandler string

func (h staticErrorHandler) ErrorHandler(_, _ any) error {
	return errors.New(string(h))
}

type valueErrorHandler struct {
	msg string
}

func (v *valueErrorHandler) ErrorHandler(value, _ any) error {
	s := gstr.ReplaceByMap(v.msg, map[string]string{
		"{value}": gconv.String(value),
	})
	return errors.New(s)
}

type value1ErrorHandler struct {
	msg string
}

func (v *value1ErrorHandler) ErrorHandler(value, _ any) error {
	s := gstr.ReplaceByMap(v.msg, map[string]string{
		"{value1}": gconv.String(value),
	})
	return errors.New(s)
}

type errorHandler struct {
	msg string
}

func (v *errorHandler) ErrorHandler(value, value1 any) error {
	s := gstr.ReplaceByMap(v.msg, map[string]string{
		"{value1}": gconv.String(value1),
		"{value}":  gconv.String(value),
	})
	return errors.New(s)
}

func parseErrorMsg(msg string) ErrorHandler {
	hasValue := strings.Contains(msg, "{value}")
	hasValue1 := strings.Contains(msg, "{value1}")

	if hasValue == false && hasValue1 {
		return &value1ErrorHandler{msg: msg}
	}
	if hasValue && hasValue1 == false {
		return &valueErrorHandler{msg: msg}
	}
	if hasValue && hasValue1 {
		return &errorHandler{msg: msg}
	}
	return staticErrorHandler(msg)
}
