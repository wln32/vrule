package ruleimpl

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
)

const betWeenMsg = "The {field} value `#{value}` must be between {min} and {max}"

type BetweenRuleNumber[T Number] struct {
	Min T
	Max T
	// 当前字段的转换函数
	FieldConvertFunc func(from reflect.Value) T
}

// 这里的input按照顺序就是min，max，其他规则同理
// 格式: between:min,max
// 说明：参数大小为min到max(支持整形和浮点类型参数)。
func (b *BetweenRuleNumber[T]) Run(ctx context.Context, input RuleFuncInput) error {
	val := b.FieldConvertFunc(input.Value)

	if val < b.Min || val > b.Max {
		if strings.Contains(input.Message, "{value}") {
			input.Message = strings.ReplaceAll(input.Message, "{value}", gconv.String(val))
		}
		return errors.New(input.Message)
	}
	return nil
}

// Format: not-in:value1,value2,...
// Format: in:value1,value2,...
type RangeRule[T Number] struct {
	Values []T
	// 当前字段的转换函数
	FieldConvertFunc func(from reflect.Value) T
}

// 格式: in:value1,value2,...
// 说明：参数值应该在value1,value2,...中（字符串匹配）。
func (r *RangeRule[T]) In(ctx context.Context, input RuleFuncInput) error {
	val := r.FieldConvertFunc(input.Value)

	for i := 0; i < len(r.Values); i++ {
		if r.Values[i] == val {
			return nil
		}
	}
	if strings.Contains(input.Message, "{value}") {
		input.Message = strings.ReplaceAll(input.Message, "{value}", gconv.String(val))
	}
	return errors.New(input.Message)

}

// 格式: not-in:value1,value2,...
// 说明：参数值不应该在value1,value2,...中（字符串匹配）。
func (r *RangeRule[T]) NotIn(ctx context.Context, input RuleFuncInput) error {
	val := r.FieldConvertFunc(input.Value)
	for i := 0; i < len(r.Values); i++ {
		if r.Values[i] == val {
			if strings.Contains(input.Message, "{value}") {
				input.Message = strings.ReplaceAll(input.Message, "{value}", gconv.String(val))
			}
			return errors.New(input.Message)
		}
	}
	return nil
}
