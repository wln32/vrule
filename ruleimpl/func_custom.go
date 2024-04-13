package ruleimpl

import (
	"context"
	"reflect"
)

type CustomRuleInput struct {
	Value     reflect.Value
	Args      string
	RuleName  string
	FieldName string

	FieldType reflect.Type
	// 支持直接传递结构体指针，用来获取结构体中其他字段值
	StructPtr reflect.Value
}

type CustomValidRuleFunc func(ctx context.Context, args *CustomRuleInput) error

type RegisterCustomRuleFunc struct {
	Args      string
	RuleName  string
	FieldName string

	FieldType reflect.Type
	Fn        CustomValidRuleFunc
	// 支持动态替换错误信息中的 {value} 字段
	// ValidRuleMessageTranslationFunc func(ctx context.Context, message error, args string, StructPtr reflect.Value) error
}

func (c *RegisterCustomRuleFunc) RunWithError(ctx context.Context, input RuleFuncInput) error {

	err := c.Fn(ctx, &CustomRuleInput{
		Value:    input.Value,
		Args:     c.Args,
		RuleName: c.RuleName,

		FieldName: c.FieldName,
		FieldType: c.FieldType,
		StructPtr: input.StructPtr,
	})
	// TODO: 自定义规则的错误信息 替换动态值 例如 {value} 替换到动态的字段值
	//if c.ValidRuleMessageTranslationFunc != nil {
	//	err = c.ValidRuleMessageTranslationFunc(ctx, err, c.Args, input.StructPtr)
	//}

	return err
}
