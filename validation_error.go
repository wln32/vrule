package vrule

import (
	"fmt"
	"reflect"
	"strings"
)

type Errors []error

func (e Errors) GetErrorWithIndex(i int) error {
	return e[i]
}

func (e Errors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}
	var errStr strings.Builder
	for i := 0; i < len(e); i++ {
		errStr.WriteString(e[i].Error() + "\n")
	}
	return errStr.String()
}

type DetailError interface {
	DetailError() string
}

type BasicFieldError struct {
	fieldName   string
	fieldErrors map[string]error // 规则名 -> 错误
}

func (e BasicFieldError) GetRuleError(rule string) error {
	return e.fieldErrors[rule]
}

func (e BasicFieldError) Error() string {
	var errStr strings.Builder
	for _, msg := range e.fieldErrors {
		errStr.WriteString(fmt.Sprintf("%s", msg))

	}
	return errStr.String()
}

func (e BasicFieldError) Errors() []error {
	var errs []error
	for _, msg := range e.fieldErrors {
		errs = append(errs, msg)

	}
	return errs
}

func (e BasicFieldError) DetailError() string {
	var errStr strings.Builder
	for rule, msg := range e.fieldErrors {
		errStr.WriteString(fmt.Sprintf("%s:[%s: %s]", e.fieldName, rule, msg.Error()))

	}
	return errStr.String()
}

type MapFieldError struct {
	fieldName   string
	fieldErrors map[reflect.Value]ValidationError
}

func (e MapFieldError) GetError(key string) ValidationError {

	for k, v := range e.fieldErrors {
		kStr := k.String()
		if kStr == key {
			return v
		}
	}
	return ValidationError{}
}

func (e MapFieldError) DetailError() string {
	var errStr strings.Builder
	for k, v := range e.fieldErrors {
		errStr.WriteString(fmt.Sprintf("%s[%v]: %s", e.fieldName, k.String(), v.DetailError()))

	}
	return errStr.String()
}

func (e MapFieldError) Errors() []error {
	var errs []error
	for _, msg := range e.fieldErrors {
		errs = append(errs, msg.Errors()...)

	}
	return errs
}

func (e MapFieldError) Error() string {
	var errStr strings.Builder
	for _, v := range e.fieldErrors {
		errStr.WriteString(fmt.Sprintf("%s", v.Error()))

	}
	return errStr.String()
}

type SliceFieldError struct {
	fieldName   string
	fieldErrors []ValidationError
}

func (e SliceFieldError) GetError(index int) ValidationError {
	return e.fieldErrors[index]
}

func (e SliceFieldError) DetailError() string {
	var errStr strings.Builder
	for k, v := range e.fieldErrors {
		errStr.WriteString(fmt.Sprintf("%s[%v]: %s", e.fieldName, k, v.DetailError()))

	}
	return errStr.String()
}

func (e SliceFieldError) Errors() []error {
	var errs []error
	for _, msg := range e.fieldErrors {
		errs = append(errs, msg.Errors()...)

	}
	return errs
}

func (e SliceFieldError) Error() string {
	var errStr strings.Builder
	for _, v := range e.fieldErrors {
		errStr.WriteString(fmt.Sprintf("%s", v.Error()))
	}
	return errStr.String()
}

type ValidationError struct {
	// 结构体名字
	structName string

	// key = FieldName
	// err = SliceFieldError | MapFieldError | ValidationError(Struct) | BasicFieldError
	fieldErrors map[string]error
}

func NewValidationError(structName string) *ValidationError {
	verr := &ValidationError{
		structName: structName,
		// 需要延迟初始化，因为可能没有错误
		// fieldErrors: make(map[string]BasicFieldError),
	}

	return verr
}
func (v ValidationError) GetFieldError(name string) BasicFieldError {
	e, _ := v.fieldErrors[name].(BasicFieldError)

	return e
}

func (v ValidationError) GetStructFieldError(name string) *ValidationError {
	return v.fieldErrors[name].(*ValidationError)
}

func (v ValidationError) GetMapFieldError(name string) MapFieldError {
	return v.fieldErrors[name].(MapFieldError)
}
func (v ValidationError) GetSliceFieldError(name string) SliceFieldError {
	return v.fieldErrors[name].(SliceFieldError)
}

func (v *ValidationError) AddFieldError(name string, fieldErr error) {
	// TODO: 延迟初始化
	//if v.fieldErrors == nil {
	//	v.fieldErrors = make(map[string]error, 1)
	//}

	v.fieldErrors[name] = fieldErr
}

func (v *ValidationError) DetailError() string {
	var errStr strings.Builder
	for field, msg := range v.fieldErrors {

		switch e := msg.(type) {
		case BasicFieldError:
			errStr.WriteString(fmt.Sprintf("%s\n", e.DetailError()))
		case MapFieldError:
			errStr.WriteString(fmt.Sprintf("%s\n", e.DetailError()))
		case SliceFieldError:
			errStr.WriteString(fmt.Sprintf("%s\n", e.DetailError()))
		case *ValidationError:
			errStr.WriteString(fmt.Sprintf("%s: %s\n", field, e.DetailError()))
		}
	}
	return errStr.String()
}

func (v *ValidationError) Errors() Errors {
	var errs []error
	for _, msg := range v.fieldErrors {
		switch e := msg.(type) {
		case BasicFieldError:
			errs = append(errs, e.Errors()...)
		case MapFieldError:
			errs = append(errs, e.Errors()...)
		case SliceFieldError:
			errs = append(errs, e.Errors()...)
		case *ValidationError:
			errs = append(errs, e.Errors()...)
		}
	}
	return errs
}

func (v *ValidationError) Error() string {
	var errStr strings.Builder
	// 默认返回第一条错误
	for _, msg := range v.fieldErrors {
		errStr.WriteString(msg.Error())
		break
	}
	return errStr.String()
}
