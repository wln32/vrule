package vrule

import (
	"fmt"
	"reflect"

	"github.com/wln32/vrule/ruleimpl"
)

// TODO：优化 字段类型转换函数和依赖字段类型转换函数，如果类型是兼容的，可以转到一个
// 比如  当前字段是int，比较字段是int32  可以统一到int64去比较

// 格式: same:field
// 说明：参数值必需与field字段参数的值相同。
// 示例：在用户注册时，提交密码Password和确认密码Password2必须相等（服务端校验）。
func getSameRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Same, ruleVals[0])
	return getEqRuleFunc(s, f, ruleVals)
}

// 格式: different:field
// 说明：参数值不能与field字段参数的值相同。
// 示例：备用邮箱OtherMailAddr和邮箱地址MailAddr必须不相同。
func getDifferentRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Different, ruleVals[0])
	return getNotEqRuleFunc(s, f, ruleVals)
}

// 格式: eq:field
// 说明：参数值必需与field字段参数的值相同。same规则的别名，功能同same规则。
// 版本：框架版本>=v2.2.0
func getEqRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Eq, ruleVals[0])
	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vf := &ruleimpl.EqRule[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](int64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}

		return ruleimpl.ValidFuncImpl(vf.EqNumber)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vf := &ruleimpl.EqRule[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](uint64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.EqNumber)
	case reflect.Float32, reflect.Float64:
		vf := &ruleimpl.EqRule[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](float64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.EqNumber)
	case reflect.String:
		field, ok := s.typ.FieldByName(ruleVals[0])
		if !ok {
			panic(fmt.Errorf("structure: %s has no fields: %s", s.longName, ruleVals[0]))
		}
		if field.Type.Kind() != reflect.String {
			panicUnsupportedTypeError("get assoc field convert func", field.Type)
		}

		vf := &ruleimpl.EqRule[string]{
			FieldName:       ruleVals[0],
			AssocFieldIndex: f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.EqString)
	case reflect.Bool:
		field, ok := s.typ.FieldByName(ruleVals[0])
		if !ok {
			panic(fmt.Errorf("structure: %s has no fields: %s", s.longName, ruleVals[0]))
		}
		if field.Type.Kind() != reflect.Bool {
			panicUnsupportedTypeError("get assoc field convert func", field.Type)
		}
		vf := &ruleimpl.EqRule[bool]{
			FieldName:       ruleVals[0],
			AssocFieldIndex: f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.EqBool)
	default:
		panicUnsupportedTypeError("get eq rule func", f.typ)
	}
	return nil

}

// 格式: not-eq:field
// 说明：参数值必需与field字段参数的值不相同。different规则的别名，功能同different规则。
// 版本：框架版本>=v2.2.0
func getNotEqRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.NotEq, ruleVals[0])
	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vf := &ruleimpl.NotEqRule[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](int64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.NotEqNumber)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vf := &ruleimpl.NotEqRule[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](uint64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.NotEqNumber)
	case reflect.Float32, reflect.Float64:
		vf := &ruleimpl.NotEqRule[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](float64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.NotEqNumber)
	case reflect.String:
		vf := &ruleimpl.NotEqRule[string]{
			FieldName:       ruleVals[0],
			AssocFieldIndex: f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.NotEqString)
	case reflect.Bool:
		vf := &ruleimpl.NotEqRule[bool]{
			FieldName:       ruleVals[0],
			AssocFieldIndex: f.requiredFieldsIndex[ruleVals[0]],
		}
		return ruleimpl.ValidFuncImpl(vf.NotEqBool)

	default:
		panicUnsupportedTypeError("get not-eq rule func", f.typ)
	}
	return nil
}

// 格式: gt:field
// 说明：参数值必需大于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func getGtRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Gt, ruleVals[0])
	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.GtRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](int64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.GtRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](uint64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.GtRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](float64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	default:
		panicUnsupportedTypeError("get gt rule func", f.typ)
	}

	return nil
}

// 格式: gte:field
// 说明：参数值必需大于等于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func getGteRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Gte, ruleVals[0])
	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.GteRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](int64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.GteRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](uint64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.GteRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](float64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	default:
		panicUnsupportedTypeError("get gte rule func", f.typ)
	}

	return nil
}

// 格式: lt:field
// 说明：参数值必需小于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func getLtRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Lt, ruleVals[0])
	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.LtRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](int64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.LtRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](uint64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.LtRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](float64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	default:
		panicUnsupportedTypeError("get lt rule func", f.typ)
	}

	return nil
}

// 格式: lte:field
// 说明：参数值必需小于等于给定字段对应的值。
// 版本：框架版本>=v2.2.0
func getLteRuleFunc(s *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Lte, ruleVals[0])
	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.LteRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](int64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.LteRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](uint64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.LteRuleNumber[float64]{
			FieldName:             ruleVals[0],
			AssocFieldConvertFunc: getAssocFieldTypeConvert[float64](s, ruleVals[0]),
			FieldConvertFunc:      getFieldReflectConvert[float64](float64(1)),
			AssocFieldIndex:       f.requiredFieldsIndex[ruleVals[0]],
		}
	default:
		panicUnsupportedTypeError("get lte rule func", f.typ)
	}

	return nil
}
