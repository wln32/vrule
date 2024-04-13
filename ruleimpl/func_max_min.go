package ruleimpl

import (
	"context"
	"errors"
	"reflect"

	"strings"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// int64 uint64 float64
type MaxRuleNumber[T Number] struct {
	Max T
	// 当前字段的转换函数
	FieldConvertFunc func(from reflect.Value) T
}

// 格式: max:max
// 说明：参数大小最大为max(支持整形和浮点类型参数)。
func (b *MaxRuleNumber[T]) Run(ctx context.Context, input RuleFuncInput) error {
	val := b.FieldConvertFunc(input.Value)
	if val > b.Max {
		if strings.Contains(input.Message, "{value}") {
			input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
				"{value}": gconv.String(val),
			})

		}
		return errors.New(input.Message)
	}
	return nil
}

// int64 uint64 float64
type MinRuleNumber[T Number] struct {
	Min T
	// 当前字段的转换函数
	FieldConvertFunc func(from reflect.Value) T
}

// 格式: max:max
// 说明：参数大小最大为max(支持整形和浮点类型参数)。
func (b *MinRuleNumber[T]) Run(ctx context.Context, input RuleFuncInput) error {

	val := b.FieldConvertFunc(input.Value)
	if val < b.Min {

		if strings.Contains(input.Message, "{value}") {
			input.Message = strings.Replace(input.Message, "{value}", gconv.String(val), 1)

		}
		return errors.New(input.Message)
	}
	return nil
}
