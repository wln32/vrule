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
}

func (c *RegisterCustomRuleFunc) RunWithError(ctx context.Context, input RuleFuncInput) error {

	err := c.Fn(ctx, &CustomRuleInput{
		Value:     input.Value,
		Args:      c.Args,
		RuleName:  c.RuleName,
		FieldName: c.FieldName,
		FieldType: c.FieldType,
		StructPtr: input.StructPtr,
	})

	return err
}
