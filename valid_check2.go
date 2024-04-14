package vrule

import (
	"context"
	"reflect"

	"github.com/wln32/vrule/ruleimpl"
)

func (s *StructRule) Valid(ctx context.Context, val reflect.Value, validOption ValidRuleOption) error {

	// 做一些初始化的准备工作
	verr := NewValidationError(s.ShortName)

	s.CheckStruct(ctx, val, verr, validOption)

	return verr
}

// struct: field*
func (s *StructRule) CheckStruct(ctx context.Context, structPtr reflect.Value,
	verr *ValidationError, validOption ValidRuleOption,
) {

	// 如果给的值是空的，不需要校验
	// 这里有可能是checkBasic进来的
	// 如果是 说明字段校验过了，但是还得验证结构体的规则
	// 就比如以下情况，b字段是满足规则的
	// 但是b是结构体，有可能B结构体里面也有v规则字段，
	// 需要递归校验
	/*
		type StructA struct {
			aa int `v:"required"`
			bb int `v:"required"`
		}
		type Example struct {
			A StructA
			B string `v:"required-without:a"`
		}
		foo := &Example{
			A: &StructA{aa:1,bb:2},
		}
	*/
	// 如果通过校验，（**struct | *struct | struct ） -> struct
	// 是不是空指针
	structVal, ok := valueIsValid(structPtr)
	if !ok {
		return
	}

	for i := 0; i < len(s.RuleFields); i++ {
		field := s.RuleFields[i]
		fieldVal := structVal.Field(field.FieldIndex)
		ok = field.CheckStructField(ctx, fieldVal, structVal, verr, validOption)
		if !ok {
			return
		}
	}

	return
}

// fieldVal 字段的值
// structPtr 结构体指针
// verr 错误
// validOption 校验选项
// 返回一个bool值，表示是否继续下个字段的校验
// field: basic | struct | slice | map
func (f *FieldRules) CheckStructField(ctx context.Context, fieldVal reflect.Value,
	structPtr reflect.Value, verr *ValidationError, validOption ValidRuleOption) bool {
	//type User struct {
	//	Name string `v:"required"`
	//  Address []string `v:"required"`
	//  Orders []*Orders `v:"required"`
	//}
	// required Address
	// 需要校验Orders 是否满足规则，
	var ruleErrors BasicFieldError
	if len(f.RuleArray) != 0 {

		for ruleName, fn := range f.Funcs {
			err := fn.Run(ctx, ruleimpl.RuleFuncInput{
				Value: fieldVal,

				StructPtr: structPtr,
				Message:   f.MsgArray[ruleName],
			})
			if err != nil {
				if ruleErrors.fieldErrors == nil {
					ruleErrors = BasicFieldError{
						fieldName:   f.FieldName,
						fieldErrors: make(map[string]error),
					}
				}
				ruleErrors.fieldErrors[ruleName] = err
				// ruleErrors.AddRuleError(ruleName, err)
				// 遇到第一个错误，就退出，不继续校验后面的规则
				if validOption.StopOnFirstError {
					break
				}
			}
		}
	}

	if ruleErrors.fieldErrors != nil {
		// 延迟初始化
		if verr.fieldErrors == nil {
			verr.fieldErrors = make(map[string]error, 1)
		}

		verr.AddFieldError(f.FieldName, ruleErrors)
		// 遇到第一个错误，就退出，不继续校验后面的规则
		if validOption.StopOnFirstError {
			return false
		}

	}

	whetherContinue := true

	switch f.kind {
	case BasicFiled, TimeField:
		return true

	case SliceFiled:
		arrLen := fieldVal.Len()
		// TODO 长度校验
		if arrLen == 0 {
			break
		}
		//sliceItemVerr := ValidationError{}
		// 进入切片，初始化当前错误的child字段
		var sliceErrors = []ValidationError{
			ValidationError{},
		}
		sliceErrIndex := 0
		for i := 0; i < arrLen; i++ {

			item := fieldVal.Index(i)
			f.StructRule.CheckStruct(ctx, item, &sliceErrors[sliceErrIndex], validOption)
			if sliceErrors[sliceErrIndex].fieldErrors != nil {

				// 遇到第一个错误，就退出，不继续校验后面的规则
				if validOption.StopOnFirstError {
					whetherContinue = false
					break
				}
				sliceErrIndex++
				if i+1 < arrLen {
					sliceErrors = append(sliceErrors, ValidationError{})

				}

			}
		}

		if sliceErrIndex != 0 {
			// 延迟初始化
			if verr.fieldErrors == nil {
				verr.fieldErrors = make(map[string]error, 1)
			}

			verr.AddFieldError(f.FieldName, SliceFieldError{
				fieldName:   f.FieldName,
				fieldErrors: sliceErrors,
			})
		}
		return whetherContinue
	case MapField:

		// 延迟初始化
		var mapErrors map[reflect.Value]ValidationError
		// struct
		mapItemVerr := ValidationError{}

		mapIter := fieldVal.MapRange()
		for mapIter.Next() {
			f.StructRule.CheckStruct(ctx, mapIter.Value(), &mapItemVerr, validOption)
			if mapItemVerr.fieldErrors != nil {
				if mapErrors == nil {
					// 延迟初始化
					mapErrors = make(map[reflect.Value]ValidationError, 1)
				}
				mapErrors[mapIter.Key()] = mapItemVerr
				mapItemVerr = ValidationError{}

			}
			// 遇到第一个错误，就退出，不继续校验后面的规则
			if validOption.StopOnFirstError {
				whetherContinue = false
				break
			}
		}
		if mapErrors != nil {
			// 延迟初始化
			if verr.fieldErrors == nil {
				verr.fieldErrors = make(map[string]error, 1)
			}

			verr.AddFieldError(f.FieldName, MapFieldError{
				fieldName:   f.FieldName,
				fieldErrors: mapErrors,
			})
		}
		return whetherContinue
	case StructrField:
		structErrors := &ValidationError{
			structName: f.FieldName,
		}
		// TODO: 增加一个错误合并的
		f.StructRule.CheckStruct(ctx, fieldVal, structErrors, validOption)
		if structErrors.fieldErrors != nil {
			// 延迟初始化
			if verr.fieldErrors == nil {
				verr.fieldErrors = make(map[string]error, 1)
			}
			verr.AddFieldError(f.FieldName, structErrors)
			if validOption.StopOnFirstError {
				whetherContinue = false
				// whetherContinue = !validOption.Bail
			}
		}

	}
	return whetherContinue
}
