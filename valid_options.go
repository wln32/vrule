package vrule

import (
	"context"
	"reflect"
)

type TranslationOption struct {
	TranslateFunc func(ctx context.Context, content string) string
}

type ParseRuleOption struct {
	// 过滤字段
	FilterFieldFunc func(structType reflect.Type, field reflect.StructField) bool
	FieldNameFunc   func(structType reflect.Type, field reflect.StructField) string
}

type ValidRuleOption struct {
	// 遇到第一个错误时是否停止
	StopOnFirstError bool
}
