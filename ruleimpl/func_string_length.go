package ruleimpl

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type StringSizeRule struct {
	Size int
}

// 格式: size:size
// 说明：参数长度为 size (长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func (s *StringSizeRule) Run(ctx context.Context, input RuleFuncInput) error {
	runes := []rune(input.Value.String())
	valueLen := len(runes)
	if valueLen != s.Size {

		if strings.Contains(input.Message, "{") {
			input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
				"{value}": gconv.String(input.Value),
				"{size}":  gconv.String(s.Size),
			})
			return fmt.Errorf(input.Message)
		}
		return fmt.Errorf(input.Message)
	}
	return nil
}

type StringLengthRule struct {
	Min int
	Max int
}

// 格式: length:min,max
// 说明：参数长度为min到max(长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func (s *StringLengthRule) Run(ctx context.Context, input RuleFuncInput) error {
	runes := []rune(input.Value.String())
	valueLen := len(runes)
	if valueLen < s.Min || valueLen > s.Max {
		if strings.Contains(input.Message, "{") {
			input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
				"{value}": gconv.String(input.Value),
			})

		}
		return fmt.Errorf(input.Message)

	}
	return nil
}

type StringMinLengthRule struct {
	Min int
}

// 格式: min-length:min
// 说明：参数长度最小为min(长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func (s *StringMinLengthRule) Run(ctx context.Context, input RuleFuncInput) error {
	runes := []rune(input.Value.String())

	if len(runes) < s.Min {
		if strings.Contains(input.Message, "{") {
			input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
				"{value}": gconv.String(input.Value),
				"{min}":   gconv.String(s.Min),
			})

		}
		return fmt.Errorf(input.Message)
	}
	return nil
}

type StringMaxLengthRule struct {
	Max int
}

// 格式: max-length:max
// 说明：参数长度最大为max(长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func (s *StringMaxLengthRule) Run(ctx context.Context, input RuleFuncInput) error {
	runes := []rune(input.Value.String())
	if len(runes) > s.Max {
		if strings.Contains(input.Message, "{") {
			input.Message = gstr.ReplaceByMap(input.Message, map[string]string{
				"{value}": gconv.String(input.Value),
				"{max}":   gconv.String(s.Max),
			})

		}
		return fmt.Errorf(input.Message)
	}
	return nil
}
