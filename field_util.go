package vrule

import (
	"fmt"
	"reflect"
	"unicode"

	"github.com/wln32/vrule/ruleimpl"

	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
)

func isTimeType(typ reflect.Type) bool {
	var timeTypeStrings = []string{
		"gtime.Time",
		"time.Time",
	}
	typString := typ.String()
	for _, s := range timeTypeStrings {
		if s == typString {
			return true
		}
	}
	return false
}

func isStructType(typ reflect.Type) (reflect.Type, bool) {
	// *[]*Struct  map[*]*Struct
	// *T -> T
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	// []T -> T
	if typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Map {
		typ = typ.Elem()
	}
	// *T -> T
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Struct {
		return typ, true
	}
	return nil, false
}

func isBasicType(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	//==========================================================
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	//==========================================================
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	case reflect.Bool:
		return true
	//==========================================================
	case reflect.String:
		return true

	case reflect.Array, reflect.Slice, reflect.Map, reflect.Ptr:
		return isBasicType(typ.Elem())
	default:
		return false
	}

}

// 判断是不是有效的值
func valueIsValid(value reflect.Value) (reflect.Value, bool) {
	// *[]*Struct  map[k]*Struct
	// *T -> T
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return value, false
		}
		value = value.Elem()
	}

	switch value.Kind() {

	case reflect.Array, reflect.Slice, reflect.Map:
		// 判断空值
		if value.Len() == 0 {
			return value, false
		}
		// 如果长度不为0，直接返回
		return value, true
	}
	// 判断是不是二级指针
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return value, false
		}
		value = value.Elem()
	}
	if value.Kind() == reflect.Struct {
		return value, true
	}

	return value, false

}

/*
// 验证是不是递归结构体

	type Example struct {
		Name string `v:"required"`
		Other *Example `v:"required"`
	}

以上结构体的Other字段要求不能为nil，
进入到Other字段里面，又会导致Other.Other不能为空，
会导致递归验证失败，
*/
func (f *FieldRules) checkRecursionRequiredRuleStruct(fieldTyp, structTyp reflect.Type) error {
	hasRequired := false
	// 判断是不是递归结构体
	if structTyp == fieldTyp {
		_, ok := f.hasRule("required")
		if ok {
			hasRequired = true
		}
	}

	if hasRequired {
		var err = fmt.Errorf(`invalid rule: Structure: %s and field: %s have the same type as the structure and are modified by the required rule, which will lead to infinite recursion`, structTyp.String(), fieldTyp.String())
		panic(err)
	}
	return nil
}

/*
required-if:field,value,...
required-unless:field,value,...

required-with:field1,field2,...
required-with-all:field1,field2,...
required-without:field1,field2,...
required-without-all:field1,field2,...
*/
// 那些规则具有变量参数
var relatedRuleName = []string{
	// required
	// "required-if","required-unless","required-with","required-without","required-without-all",
	// cmp
	ruleimpl.Lte, ruleimpl.Lt, ruleimpl.Gte, ruleimpl.Gt,
	ruleimpl.NotEq, ruleimpl.Eq, ruleimpl.Different, ruleimpl.Same,

	// time
	ruleimpl.Before, ruleimpl.After, ruleimpl.BeforeEqual, ruleimpl.AfterEqual,
}

var relatedRuleNameMap = make(map[string]struct{})

func init() {

	for _, name := range relatedRuleName {
		relatedRuleNameMap[name] = struct{}{}
	}
}

// 校验required规则的参数是否正确
func (f *FieldRules) checkRequiredRulesIsValid(structName string, requiredRule string, ruleVals []string) []string {
	// required-if -> if
	switch requiredRule {
	case "required":
		if len(ruleVals) > 0 {

			panicRuleParameterWithMsgError(structName, f.FieldName, requiredRule, "without parameters", gconv.String(len(ruleVals)))
		}
	case "required-if", "required-unless":
		if len(ruleVals)%2 == 0 {
			return ruleVals
		}

		panicRuleParameterWithMsgError(structName, f.FieldName, requiredRule, "The number of parameters must be a multiple of 2", gconv.String(len(ruleVals)))
	case "required-with", "required-with-all", "required-with-out", "required-without-all":
		if len(ruleVals) == 0 {
			panicRuleParameterWithMsgError(structName, f.FieldName, requiredRule, "at least one", "0")
		}
		return ruleVals
	}
	return ruleVals
}

/*
before:field
before-equal:field
after:field
after-equal:field

lte:field
lt:field
gte:field
gt:field
not-eq:field
eq:field
different:field
same:field
*/
// 校验关联规则的参数是否正确
func (f *FieldRules) checkAssocRulesIsValid(structName string, ruleName string, ruleVals []string) []string {
	if len(ruleVals) > 1 {
		panicRuleParameterError(structName, f.FieldName, ruleName, 1, len(ruleVals))

	}
	_, ok := relatedRuleNameMap[ruleName]
	if !ok {
		panicInvalidRuleError(structName, f.FieldName, ruleName)
	}
	switch ruleName {
	case ruleimpl.Lte:
	case ruleimpl.Lt:
	case ruleimpl.Gte:
	case ruleimpl.Gt:
	case ruleimpl.NotEq:
	case ruleimpl.Eq:
	case ruleimpl.Different:
	case ruleimpl.Same:

	}

	return ruleVals
}

// 递归获取结构体的所有字段
func getStructFields(typ reflect.Type) []gstructs.Field {
	var fields []gstructs.Field
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 判断是不是公开的字段
		if unicode.IsLower(rune(field.Name[0])) {
			continue
		}
		// 判断是不是接口
		if field.Type.Kind() == reflect.Interface {
			continue
		}

		// 判断有没有带v或者valid 标签的
		tag := field.Tag.Get("v")
		if tag == "" {
			tag = field.Tag.Get("valid")
		}
		if tag == "" {
			// 如果没有tag，判断是不是结构体
			// slice map ptr 都解引用
			_, ok := isStructType(field.Type)
			if !ok {
				continue
			}
		}
		fields = append(fields, gstructs.Field{
			Value: reflect.Value{},
			Field: field,
		})
	}

	return fields
}
