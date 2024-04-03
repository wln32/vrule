package ruleimpl

import (
	"context"
	"errors"
	"reflect"
)

// required: string len==0
// required: ptr == nil
// required: struct 递归
// required: slice  len==0
// required: map len==0
func requiredEmpty(input RuleFuncInput) error {
	//typ := input.ValueType
	//val := reflect.ValueOf(input.Value)
	//
	//switch typ.Kind() {
	//case reflect.String, reflect.Slice, reflect.Map:
	//	if val.Len() == 0 {
	//		return errors.New(input.Message)
	//	}
	//case reflect.Ptr:
	//	if val.IsNil() {
	//		return errors.New(input.Message)
	//	}
	//}

	return nil
}

// TODO : 后续传递Value的时候，可以不用再调用Interface()，直接传递reflect.Value会更好
func RequiredPtrFunc(ctx context.Context, input RuleFuncInput) error {
	val := reflect.ValueOf(input.Value)
	if val.IsNil() {
		return errors.New(input.Message)
	}
	return nil
}

// TODO : 后续传递Value的时候，可以不用再调用Interface()，直接传递reflect.Value会更好
func RequiredLengthFunc(ctx context.Context, input RuleFuncInput) error {

	val := reflect.ValueOf(input.Value)
	if val.Len() == 0 {
		return errors.New(input.Message)
	}
	return nil
}

// TODO : 后续传递Value的时候，可以不用再调用Interface()，直接传递reflect.Value会更好
func RequiredStringLengthFunc(ctx context.Context, input RuleFuncInput) error {

	l := input.Value.(string)
	if len(l) == 0 {
		return errors.New(input.Message)
	}
	return nil
}

type RequiredIfRuleArg struct {
	AssocFieldIndex int
	Value           any
}

type RequiredIfRule struct {
	// filed => value
	// 校验依赖字段是否有值的函数
	// map[依赖字段]=>  val
	//AssocFieldValues map[string]any
	// 判定当前字段是否有值的函数
	IsEmpty ValidFunc
	// TODO: 直接存储依赖字段的索引，如果提前满足条件，可以跳过后面的，延迟求值
	AssocFields []RequiredIfRuleArg
}

// Run RequiredIfFunc 格式: required-if:field,value,...
// 说明：必需参数(当field=value时，当前字段必须有值)。多个字段以,号分隔。
func (r *RequiredIfRule) Run(ctx context.Context, input RuleFuncInput) error {

	//for field, val := range r.AssocFieldValues {
	//	v, _ := input.AssocFieldValues[field]
	//	// 如果field=value时，判断当前字段是否有值
	//	if v == val {
	//		return r.IsEmpty.Run(ctx, input)
	//	}
	//}
	for _, assocField := range r.AssocFields {
		v := input.StructPtr.Field(assocField.AssocFieldIndex).Interface()
		if v == assocField.Value {
			return r.IsEmpty.Run(ctx, input)
		}
	}

	return nil
}

type RequiredUnlessRule struct {
	// filed => value
	// 校验依赖字段是否有值的函数
	// map[依赖字段]=> 依赖字段验证是否有值的函数
	//AssocFieldValues map[string]any
	// 判定当前字段是否有值的函数
	IsEmpty ValidFunc

	// TODO: 直接存储依赖字段的索引，如果提前满足条件，可以跳过后面的，延迟求值
	AssocFields []RequiredIfRuleArg
}

// Run 格式: required-unless:field,value,...
// 说明：必需参数(当field!=value时，当前字段必须有值)。多个字段以,号分隔。
func (r *RequiredUnlessRule) Run(ctx context.Context, input RuleFuncInput) error {
	//for field, val := range r.AssocFieldValues {
	//	v, _ := input.AssocFieldValues[field]
	//	// 当field!=value时，当前字段必须有值
	//	if v != val {
	//		return r.IsEmpty.Run(ctx, input)
	//	}
	//
	//}
	for _, assocField := range r.AssocFields {
		v := input.StructPtr.Field(assocField.AssocFieldIndex).Interface()
		if v != assocField.Value {
			return r.IsEmpty.Run(ctx, input)
		}
	}

	return nil
}

type RequiredWithRuleArg struct {
	// 关联字段的索引
	AssocFieldIndex int
	// 关联字段的校验函数
	AssocFieldValidFunc ValidFunc
}

type RequiredFieldsRule struct {
	//required-without-all: field1,field2,...
	//required-without: field1,field2,...
	//required-with-all: field1,field2,...
	//required-with: field1,field2,...
	// 校验依赖字段是否有值的函数
	// map[依赖字段]=> 依赖字段验证是否有值的函数
	// AssocFieldValidFunc map[string]ValidFunc
	// 依赖的字段名字
	// 判定当前字段是否有值的函数
	IsEmpty ValidFunc
	// TODO: 直接存储依赖字段的索引，如果提前满足条件，可以跳过后面的，延迟求值

	AssocFields []RequiredWithRuleArg
}

// RequiredWith 格式: required-with:field1,field2,...
// 说明：必需参数(只要有一个字段有值)。当前字段必须有值
func (r *RequiredFieldsRule) RequiredWith(ctx context.Context, input RuleFuncInput) (err error) {

	//for field, fn := range r.AssocFieldValidFunc {
	//	err = fn.Run(ctx, RuleFuncInput{
	//		Value:   input.AssocFieldValues[field],
	//		Message: field,
	//	})
	//	// 没有错误就是有值
	//	if err == nil {
	//		// 只要有一个字段有值
	//		break
	//	}
	//
	//}

	for _, assoc := range r.AssocFields {
		val := input.StructPtr.Field(assoc.AssocFieldIndex).Interface()
		err = assoc.AssocFieldValidFunc.Run(ctx, RuleFuncInput{
			Value:   val,
			Message: "1",
		})
		if err == nil {
			// 只要有一个字段有值
			break
		}
	}

	// 只要有一个字段有值
	if err == nil {
		// 判断当前字段是否有值
		err = r.IsEmpty.Run(ctx, RuleFuncInput{
			Value:   input.Value,
			Message: "1",
		})
		if err != nil {
			return errors.New(input.Message)
		}

	}

	return nil
}

// RequiredWithAll 格式: required-with-all:field1,field2,...
// 说明：必须参数(全部字段都得有值)。当前字段必须有值
// 示例：当Id，Name，Gender，WifeName全部不为空时，HusbandName必须不为空。
func (r *RequiredFieldsRule) RequiredWithAll(ctx context.Context, input RuleFuncInput) (err error) {
	//for field, fn := range r.AssocFieldValidFunc {
	//	err = fn.Run(ctx, RuleFuncInput{
	//		Value:   input.AssocFieldValues[field],
	//		Message: field,
	//	})
	//	// 如果给定的字段有一个没有值
	//	if err != nil {
	//		return nil
	//	}
	//}
	for _, assoc := range r.AssocFields {
		val := input.StructPtr.Field(assoc.AssocFieldIndex).Interface()
		err = assoc.AssocFieldValidFunc.Run(ctx, RuleFuncInput{
			Value:   val,
			Message: "1",
		})
		if err != nil {
			return nil
		}
	}

	// 所有字段都有值
	// 如果给定的字段没有值
	if r.IsEmpty.Run(ctx, RuleFuncInput{
		Value:   input.Value,
		Message: "required-with-all",
	}) != nil {
		return errors.New(input.Message)
	}

	return nil
}

// RequiredWithout 格式: required-without:field1,field2,...
// 说明：必需参数(只要有一个字段为空)。当前字段必须有值
// 示例：当Id或WifeName为空时，HusbandName必须不为空
func (r *RequiredFieldsRule) RequiredWithout(ctx context.Context, input RuleFuncInput) (err error) {
	//for field, fn := range r.AssocFieldValidFunc {
	//	err = fn.Run(ctx, RuleFuncInput{
	//		Value:   input.AssocFieldValues[field],
	//		Message: field,
	//	})
	//	// 如果给定的字段有一个为空
	//	if err != nil {
	//		break
	//	}
	//}
	for _, assoc := range r.AssocFields {
		val := input.StructPtr.Field(assoc.AssocFieldIndex).Interface()
		err = assoc.AssocFieldValidFunc.Run(ctx, RuleFuncInput{
			Value:   val,
			Message: "1",
		})
		if err != nil {
			break
		}
	}

	if err != nil && r.IsEmpty.Run(ctx,
		RuleFuncInput{
			Value:   input.Value,
			Message: "required-without",
		}) != nil {
		return errors.New(input.Message)
	}

	return nil
}

// RequiredWithoutAll 格式: required-without-all:field1,field2,...
// 说明：必须参数(所有字段都为空时)。当前字段必须有值
// 示例：当Id和WifeName都为空时，HusbandName必须不为空。当前字段必须有值
func (r *RequiredFieldsRule) RequiredWithoutAll(ctx context.Context, input RuleFuncInput) (err error) {

	//for field, fn := range r.AssocFieldValidFunc {
	//	err = fn.Run(ctx, RuleFuncInput{
	//		Value:   input.AssocFieldValues[field],
	//		Message: field,
	//	})
	//	if err == nil {
	//		// 如果给定的字段有一个不为空，当前字段不需要有值
	//		return nil
	//	}
	//}

	for _, assoc := range r.AssocFields {
		val := input.StructPtr.Field(assoc.AssocFieldIndex).Interface()
		err = assoc.AssocFieldValidFunc.Run(ctx, RuleFuncInput{
			Value:   val,
			Message: "1",
		})
		if err == nil {
			return nil
		}
	}

	// 给定的所有字段都没有值
	// 判断当前字段是否有值
	if r.IsEmpty.Run(ctx,
		RuleFuncInput{
			Value:   input.Value,
			Message: "required-without-all",
		}) != nil {
		return errors.New(input.Message)
	}
	return nil
}