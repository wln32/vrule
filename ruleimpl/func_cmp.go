package ruleimpl

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 比较大小  T 的类型可以是 comparable 这样可以比较任意值类型
type EqRule[T Number | string | bool] struct {
	// TODO FieldName待删除，直接在解析阶段替换为字段名
	FieldName string
	// 依赖字段的转换函数
	AssocFieldConvertFunc func(from reflect.Value) T
	// 当前字段的转换函数，数字类型转到float64，字符串不转
	FieldConvertFunc func(from reflect.Value) T
	// 关联字段索引
	AssocFieldIndex int
}

// 格式: eq:field
// 说明：参数值必需与field字段参数的值相同。same规则的别名，功能同same规则。
// 版本：框架版本>=v2.2.0
func (e *EqRule[T]) EqNumber(ctx context.Context, input RuleFuncInput) error {
	thisVal := e.FieldConvertFunc(input.Value)
	fieldVal := e.AssocFieldConvertFunc(input.StructPtr.Field(e.AssocFieldIndex))

	if fieldVal == thisVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": e.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)

}

// 格式: eq:field
// 说明：参数值必需与field字段参数的值相同。same规则的别名，功能同same规则。
// 版本：框架版本>=v2.2.0
func (e *EqRule[T]) EqString(ctx context.Context, input RuleFuncInput) error {
	thisVal := input.Value.String()
	fieldVal := input.StructPtr.Field(e.AssocFieldIndex).String()

	if fieldVal == thisVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": e.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)

}

// 格式: eq:field
// 说明：参数值必需与field字段参数的值相同。same规则的别名，功能同same规则。
// 版本：框架版本>=v2.2.0
func (e *EqRule[T]) EqBool(ctx context.Context, input RuleFuncInput) error {
	thisVal := input.Value.Bool()
	fieldVal := input.StructPtr.Field(e.AssocFieldIndex).Bool()

	if fieldVal == thisVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": e.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)

}

// 格式: not-eq:field
// 说明：参数值必需与field字段参数的值不相同。different规则的别名，功能同different规则。
// 版本：框架版本>=v2.2.0
type NotEqRule[T Number | string | bool] struct {
	// TODO FieldName待删除，直接在解析阶段替换为字段名
	FieldName string
	// 依赖字段的转换函数
	AssocFieldConvertFunc func(from reflect.Value) T
	// 当前字段的转换函数，数字类型转到float64，字符串不转
	FieldConvertFunc func(from reflect.Value) T
	// 关联字段索引
	AssocFieldIndex int
}

// 格式: same:field
// 说明：参数值必需与field字段参数的值相同。
// 示例：在用户注册时，提交密码Password和确认密码Password2必须相等（服务端校验）。
func (e *NotEqRule[T]) NotEqNumber(ctx context.Context, input RuleFuncInput) error {
	thisVal := e.FieldConvertFunc(input.Value)
	fieldVal := e.AssocFieldConvertFunc(input.StructPtr.Field(e.AssocFieldIndex))

	if fieldVal != thisVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": e.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)
}

// 格式: same:field
// 说明：参数值必需与field字段参数的值相同。
// 示例：在用户注册时，提交密码Password和确认密码Password2必须相等（服务端校验）。
func (e *NotEqRule[T]) NotEqString(ctx context.Context, input RuleFuncInput) error {
	thisVal := input.Value.String()
	fieldVal := input.StructPtr.Field(e.AssocFieldIndex).String()

	if fieldVal != thisVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": e.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)
}

// 格式: same:field
// 说明：参数值必需与field字段参数的值相同。
// 示例：在用户注册时，提交密码Password和确认密码Password2必须相等（服务端校验）。
func (e *NotEqRule[T]) NotEqBool(ctx context.Context, input RuleFuncInput) error {
	thisVal := input.Value.Bool()
	fieldVal := input.StructPtr.Field(e.AssocFieldIndex).Bool()

	if fieldVal != thisVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": e.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)
}

type GtRuleNumber[T Number] struct {
	// TODO FieldName待删除，直接在解析阶段替换为字段名
	FieldName string
	// 依赖字段的转换函数
	AssocFieldConvertFunc func(from reflect.Value) T
	// 当前字段的转换函数，数字类型转到float64，字符串不转
	FieldConvertFunc func(from reflect.Value) T
	// 关联字段索引
	AssocFieldIndex int
}

// 格式: gt:field
// 说明：参数值必需大于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func (g *GtRuleNumber[T]) Run(ctx context.Context, input RuleFuncInput) error {
	const gtErrorMsg = "The {field} value `{value}` must be greater than field {field1} value `{value1}`"
	thisVal := g.FieldConvertFunc(input.Value)

	fieldVal := g.AssocFieldConvertFunc(input.StructPtr.Field(g.AssocFieldIndex))
	if thisVal > fieldVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": g.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)
}

type GteRuleNumber[T Number] struct {
	// TODO FieldName待删除，直接在解析阶段替换为字段名
	FieldName string
	// 依赖字段的转换函数
	AssocFieldConvertFunc func(from reflect.Value) T
	// 当前字段的转换函数，数字类型转到float64，字符串不转
	FieldConvertFunc func(from reflect.Value) T
	// 关联字段索引
	AssocFieldIndex int
}

// 格式: gte:field
// 说明：参数值必需大于或等于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func (g *GteRuleNumber[T]) Run(ctx context.Context, input RuleFuncInput) error {
	const gtErrorMsg = "The {field} value `{value}` must be greater than field {field1} value `{value1}`"
	thisVal := g.FieldConvertFunc(input.Value)
	fieldVal := g.AssocFieldConvertFunc(input.StructPtr.Field(g.AssocFieldIndex))
	if thisVal >= fieldVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": g.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)
}

type LtRuleNumber[T Number] struct {
	// TODO FieldName待删除，直接在解析阶段替换为字段名
	FieldName string
	// 依赖字段的转换函数
	AssocFieldConvertFunc func(from reflect.Value) T
	// 当前字段的转换函数，数字类型转到float64，字符串不转
	FieldConvertFunc func(from reflect.Value) T
	// 关联字段索引
	AssocFieldIndex int
}

// 格式: lt:field
// 说明：参数值必需小于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func (g *LtRuleNumber[T]) Run(ctx context.Context, input RuleFuncInput) error {
	const gtErrorMsg = "The {field} value `{value}` must be greater than field {field1} value `{value1}`"
	thisVal := g.FieldConvertFunc(input.Value)
	fieldVal := g.AssocFieldConvertFunc(input.StructPtr.Field(g.AssocFieldIndex))
	if thisVal < fieldVal {
		return nil
	}

	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": g.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)
}

type LteRuleNumber[T Number] struct {
	// TODO FieldName待删除，直接在解析阶段替换为字段名
	FieldName string
	// 依赖字段的转换函数
	AssocFieldConvertFunc func(from reflect.Value) T
	// 当前字段的转换函数，数字类型转到float64，字符串不转
	FieldConvertFunc func(from reflect.Value) T
	// 关联字段索引
	AssocFieldIndex int
}

// 格式: lte:field
// 说明：参数值必需小于或等于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func (g *LteRuleNumber[T]) Run(ctx context.Context, input RuleFuncInput) error {
	const gtErrorMsg = "The {field} value `{value}` must be greater than field {field1} value `{value1}`"
	thisVal := g.FieldConvertFunc(input.Value)
	fieldVal := g.AssocFieldConvertFunc(input.StructPtr.Field(g.AssocFieldIndex))
	if thisVal <= fieldVal {
		return nil
	}
	input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  gconv.String(thisVal),
		"{field1}": g.FieldName,
		"{value1}": gconv.String(fieldVal),
	})
	return fmt.Errorf(input.Message)
}
