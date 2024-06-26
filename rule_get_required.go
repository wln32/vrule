package vrule

import (
	"context"
	"fmt"
	"reflect"

	"github.com/wln32/vrule/ruleimpl"

	"github.com/gogf/gf/v2/util/gconv"
)

/*
 */
// required
func getRequiredRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	// 判断当前类型
	// 指针
	// slice map
	// string

	return getRequiredWithFieldValidFunc(f.Type)

}

// RequiredIfFunc 格式: required-if:field,value,...
// 说明：必需参数(当任意所给定字段值与所给值相等时，即：当field字段的值为value时，当前验证字段为必须参数)。多个字段以,号分隔。
// 示例：当Gender字段为1时WifeName字段必须不为空， 当Gender字段为2时HusbandName字段必须不为空。
// 当前字段类型: 符合required规则即可
// field类型：基础类型(数字类型，布尔类型，字符串类型)
// value类型：必须是字面量类型，例如1 false hello，不能是变量
func getRequiredIfRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {

	args := make([]ruleimpl.RequiredIfRuleArg, 0)

	for i := 0; i < len(ruleVals); i += 2 {
		field, ok := s.Type.FieldByName(ruleVals[i])
		if ok {
			//fns[ruleVals[i]] = getRequiredFieldValue(field.Type.Kind(), ruleVals[i+1])
			val, isstring := getRequiredFieldValue(field.Type, ruleVals[i+1])

			arg := ruleimpl.RequiredIfRuleArg{
				AssocFieldIndex: int32(field.Index[0]),
				IsString:        isstring,
				Value:           val,
			}
			args = append(args, arg)
		} else {
			// 不存在的字段
		}
	}

	return &ruleimpl.RequiredIfRule{
		IsEmpty: getRequiredWithFieldValidFunc(f.Type),
		// AssocFieldValues: fns,
		AssocFields: args,
	}
}

// 格式: required-unless:field,value,...
// 说明：必需参数(当所给定字段值与所给值都不相等时，即：当field字段的值不为value时，当前验证字段为必须参数)。多个字段以,号分隔。
// 示例：当Gender不等于0且Gender不等于2时，WifeName必须不为空；当Id 不等于0且 Gender 不等于2时， HusbandName 必须不为空。
// 当前字段类型: 符合required规则即可
// field类型：基础类型(数字类型，布尔类型，字符串类型)
// value类型：必须是字面量类型，例如1 false hello，不能是变量
func getRequiredUnlessRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	args := make([]ruleimpl.RequiredIfRuleArg, 0)

	for i := 0; i < len(ruleVals); i += 2 {
		field, ok := s.Type.FieldByName(ruleVals[i])
		if ok {
			// fns[ruleVals[i]] = getRequiredFieldValue(field.Type.Kind(), ruleVals[i+1])
			val, isstring := getRequiredFieldValue(field.Type, ruleVals[i+1])

			args = append(args, ruleimpl.RequiredIfRuleArg{
				AssocFieldIndex: int32(field.Index[0]),
				IsString:        isstring,
				Value:           val,
			})
		} else {
			// 不存在的字段
		}
	}

	return &ruleimpl.RequiredUnlessRule{
		IsEmpty: getRequiredWithFieldValidFunc(f.Type),
		// AssocFieldValues: fns,
		AssocFields: args,
	}
}

// 格式: required-with:field1,field2,...
// 说明：必需参数(当所给定任意字段值其中之一不为空时)。
// 示例：当WifeName不为空时，HusbandName必须不为空。
// required-with: field1,field2,field3
func getRequiredWithRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	// 直接调用required函数来做判断
	// 需要判断每个字段是什么类型，然后给RequiredFieldsRule 这个结构体传递一个required函数数组即可
	ruleFunc := &ruleimpl.RequiredFieldsRule{
		//AssocFieldValidFunc: getRequiredFuncs(s, ruleVals),
		IsEmpty:     getRequiredWithFieldValidFunc(f.Type),
		AssocFields: getRequiredWithAssocFieldFuncs(s, ruleVals),
	}

	return ruleimpl.ValidFuncImpl(ruleFunc.RequiredWith)
}

// 格式: required-with-all:field1,field2,...
// 说明：必须参数(当所给定所有字段值全部都不为空时)。
// 示例：当Id，Name，Gender，WifeName全部不为空时，HusbandName必须不为空。
// required-with-all: field1,field2,field3
func getRequiredWithAllRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {

	ruleFunc := &ruleimpl.RequiredFieldsRule{
		//AssocFieldValidFunc: getRequiredFuncs(s, ruleVals),
		IsEmpty:     getRequiredWithFieldValidFunc(f.Type),
		AssocFields: getRequiredWithAssocFieldFuncs(s, ruleVals),
	}

	return ruleimpl.ValidFuncImpl(ruleFunc.RequiredWithAll)
}

// 格式: required-without:field1,field2,...
// 说明：必需参数(当所给定任意字段值其中之一为空时)。
// 示例：当Id或WifeName为空时，HusbandName必须不为空
// required-without: field1,field2,field3
func getRequiredWithoutRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	ruleFunc := &ruleimpl.RequiredFieldsRule{
		//AssocFieldValidFunc: getRequiredFuncs(s, ruleVals),
		IsEmpty:     getRequiredWithFieldValidFunc(f.Type),
		AssocFields: getRequiredWithAssocFieldFuncs(s, ruleVals),
	}
	return ruleimpl.ValidFuncImpl(ruleFunc.RequiredWithout)
}

// 格式: required-without-all:field1,field2,...
// 说明：必须参数(当所给定所有字段值全部都为空时)。
// 示例：当Id和WifeName都为空时，HusbandName必须不为空。
// required-without-all: field1,field2,field3
func getRequiredWithoutAllRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	ruleFunc := &ruleimpl.RequiredFieldsRule{
		//AssocFieldValidFunc: getRequiredFuncs(s, ruleVals),
		IsEmpty:     getRequiredWithFieldValidFunc(f.Type),
		AssocFields: getRequiredWithAssocFieldFuncs(s, ruleVals),
	}
	return ruleimpl.ValidFuncImpl(ruleFunc.RequiredWithoutAll)
}

// 获取每个字段的类型 ptr slice map string 生成具体的验证函数，更快的实现
// 输入：字段名组成的数组，类型从结构体里面获取
// 输出：required数组函数
// TODO ： 直接存储依赖字段的索引，延迟求值，减少函数调用
func getRequiredFuncs(f *StructRule, assocFields []string) map[string]ruleimpl.ValidFunc {
	funcs := make(map[string]ruleimpl.ValidFunc, 0)
	for _, relatedFiled := range assocFields {
		field, ok := f.Type.FieldByName(relatedFiled)
		if ok {
			funcs[field.Name] = getRequiredWithFieldValidFunc(field.Type)
		}
	}

	return funcs
}

func getRequiredWithAssocFieldFuncs(f *StructRule, assocFields []string) []ruleimpl.RequiredWithRuleArg {

	funcs := make([]ruleimpl.RequiredWithRuleArg, 0)
	for _, assocField := range assocFields {
		field, ok := f.Type.FieldByName(assocField)
		if ok {
			funcs = append(funcs, ruleimpl.RequiredWithRuleArg{
				AssocFieldIndex:     field.Index[0],
				AssocFieldValidFunc: getRequiredWithFieldValidFunc(field.Type),
			})
		}
	}

	return funcs
}

func getRequiredWithFieldValidFunc(typ reflect.Type) ruleimpl.ValidFunc {
	switch typ.Kind() {
	case reflect.Ptr:
		return ruleimpl.ValidFuncImpl(ruleimpl.RequiredPtrFunc)
	case reflect.Slice, reflect.Map, reflect.Array:
		return ruleimpl.ValidFuncImpl(ruleimpl.RequiredLengthFunc)
	case reflect.String:
		return ruleimpl.ValidFuncImpl(ruleimpl.RequiredLengthFunc)
	case reflect.Struct:
		// struct值类型不生效
		return ruleimpl.ValidFuncImpl(func(ctx context.Context, input ruleimpl.RuleFuncInput) error {
			return nil
		})

	default:
		// 只要不报错
		getRequiredFieldValue(typ, "0")
		return ruleimpl.ValidFuncImpl(func(ctx context.Context, input ruleimpl.RuleFuncInput) error {
			return nil
		})
		// 如果是其他的基础值类型，默认永远有值，
		// 等后续吧这些required-withxxx规则优化，
		// 依赖的是值类型时，
		// 直接把当前字段优化为required规则
		// 可以减少一次函数调用，
	}

}

type requiredCmpFunc = func(field any, value any) bool

// 主要是required-if，required-unless使用
func getRequiredFieldValue(typ reflect.Type, val string) (a any, isstring bool) {
	switch typ.String() {
	case "int":
		return gconv.Int(val), false
	case "int8":
		return gconv.Int8(val), false
	case "int16":
		return gconv.Int16(val), false
	case "int32":
		return gconv.Int32(val), false
	case "int64":
		return gconv.Int64(val), false
	case "uint":
		return gconv.Uint(val), false
	case "uint8":
		return gconv.Uint8(val), false
	case "uint16":
		return gconv.Uint16(val), false
	case "uint32":
		return gconv.Uint32(val), false
	case "uint64":
		return gconv.Uint64(val), false
	case "float32":
		return gconv.Float32(val), false
	case "float64":
		return gconv.Float64(val), false
	case "bool":
		return gconv.Bool(val), false
	case "string":
		return val, true
	default:
		// 如果不是基础类型
		// 查看底层是否为基础类型
		// 构造一个自定义的基础类型
		v := reflect.New(typ)

		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v.Elem().SetInt(gconv.Int64(val))
			a = v.Elem().Interface()
			return

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v.Elem().SetUint(gconv.Uint64(val))
			a = v.Elem().Interface()
			return

		case reflect.Float32, reflect.Float64:
			v.Elem().SetFloat(gconv.Float64(val))
			a = v.Elem().Interface()
			return
		case reflect.Bool:
			v.Elem().SetBool(gconv.Bool(val))
			a = v.Elem().Interface()
			return
		case reflect.String:
			v.Elem().SetString(val)
			// 如果是string，则直接返回，不需要返回空接口
			a = v.Elem().String()
			isstring = true
			return
		default:
			panic(fmt.Errorf("getRequiredFieldValue: Unsupported parameter type: %s", typ))
		}

	}

	return

}
