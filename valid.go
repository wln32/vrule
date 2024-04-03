package vrule

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/i18n/gi18n"
)

var defaultValidator = New()

type Validator struct {
	cache *StructCache

	// 解析rule时使用的option
	parseRuleOption *ParseRuleOption
	// 校验rule时使用的option
	validRuleOption *ValidRuleOption
	// 获取msg时使用的option
	translationOption *TranslationOption

	i18n *gi18n.Manager
}

func New() *Validator {
	v := &Validator{}
	var cache = &StructCache{
		cache: map[string]*StructRule{},
	}

	var parse = &ParseRuleOption{
		// 过滤字段, 默认全部解析
		FilterFieldFunc: func(structType reflect.Type, field reflect.StructField) bool {
			return false
		},
		FieldNameFunc: func(structType reflect.Type, field reflect.StructField) string {
			return ""
		},
	}
	var valid = &ValidRuleOption{
		StopOnFirstError: false,
	}

	i18n := gi18n.New(gi18n.Options{
		Path:     getI18nPath(),
		Language: "en",
	})

	var trans = &TranslationOption{
		TranslateFunc: func(ctx context.Context, content string) string {
			return ""
		},
	}

	v.i18n = i18n
	v.cache = cache
	v.parseRuleOption = parse
	v.validRuleOption = valid
	v.translationOption = trans

	return v
}

func (v *Validator) Struct(a any) error {
	val := reflect.ValueOf(a)

	typ, ok := isStructType(val.Type())
	if !ok {
		panic("type is not struct")
	}
	structRule := v.cache.GetStructRuleOrCreate(typ, v)
	if structRule == nil {
		return fmt.Errorf("current structure: %s has no rules to verify", val.Type())
	}
	return structRule.Valid(context.TODO(), val, *v.validRuleOption)
}

func (v *Validator) StructCtx(ctx context.Context, a any) error {

	val := reflect.ValueOf(a)

	typ, ok := isStructType(val.Type())
	if !ok {
		panic("type is not struct")
	}
	structRule := v.cache.GetStructRuleOrCreate(typ, v)
	if structRule == nil {
		return fmt.Errorf("current structure: %s has no rules to verify", val.Type())
	}
	return structRule.Valid(ctx, val, *v.validRuleOption)
}

func (v *Validator) StructNotCache(a any) error {
	val := reflect.ValueOf(a)

	typ, ok := isStructType(val.Type())
	if !ok {
		panic("type is not struct")
	}

	newv := *v
	(newv).cache = nil
	structRule := v.cache.GetStructRuleOrCreate(typ, &newv)
	if structRule == nil {
		return fmt.Errorf("current structure: %s has no rules to verify", val.Type())
	}
	return structRule.Valid(context.TODO(), val, *v.validRuleOption)
}

// 过滤那些字段，主要是带有v规则的字段
func (v *Validator) SetFilterFieldFunc(f func(structType reflect.Type, field reflect.StructField) bool) {
	if f != nil {
		v.parseRuleOption.FilterFieldFunc = f
	}
}
func (v *Validator) SetFieldNameFunc(f func(structType reflect.Type, field reflect.StructField) string) {
	if f != nil {
		v.parseRuleOption.FieldNameFunc = f
	}
}

func (v *Validator) StopOnFirstError(b bool) {
	v.validRuleOption.StopOnFirstError = b
}
