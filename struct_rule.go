package vrule

import (
	"reflect"

	"github.com/wln32/vrule/ruleimpl"
)

type StructRule struct {
	ruleFields []*FieldRules
	// pkgpath + structName
	longName string
	// shortName = structName
	shortName string
	typ       reflect.Type
}

func (s *StructRule) StructType() reflect.Type {
	return s.typ
}

func (s *StructRule) LongName() string {
	return s.longName
}
func (s *StructRule) ShortName() string {
	return s.shortName
}

// 删除掉没有rule的字段
func (f *StructRule) deleteEmptyRuleField() {

	for i := 0; i < len(f.ruleFields); {
		field := f.ruleFields[i]
		// 只能删除基础类型或者time类型，没有校验规则的字段
		// 防止结构体里面有校验规则的字段就被删掉
		if field.structFields != nil {
			i++
			continue
		}
		if len(field.ruleArray) == 0 {

			f.ruleFields = append(f.ruleFields[:i], f.ruleFields[i+1:]...)

		} else {
			// 如果没删除，在自增，不然会有bug
			i++
		}
	}
}

// 设置结构体所有字段的验证函数
func (f *StructRule) setFieldsRuleValidFunc() {

	for i := 0; i < len(f.ruleFields); i++ {
		field := f.ruleFields[i]
		if len(field.ruleArray) != 0 {
			f.setFieldRuleValidFunc(field)
		}
	}
}

// 设置字段的所有验证函数
func (f *StructRule) setFieldRuleValidFunc(fieldRule *FieldRules) {
	if fieldRule.Funcs == nil {
		fieldRule.Funcs = make(map[string]ruleimpl.ValidFunc)
	}
	for ruleName, ruleVals := range fieldRule.ruleArray {

		registerFunc, ok := builtinRulesMapToFunc[ruleName]
		if ok {
			fieldRule.Funcs[ruleName] = registerFunc(f, fieldRule, ruleVals)
		} else {
			// 支持自定义规则
			fn := getCustomValidRuleFunc(f, fieldRule, ruleName, ruleVals)
			if fn == nil {
				panicInvalidRuleError(f.longName, fieldRule.fieldName, ruleName)
			}
			// 绑定到验证的字段上
			fieldRule.Funcs[ruleName] = fn
		}
	}
}

// 设置关联的字段索引，
func (f *StructRule) setIndexAssocFields() {

	for i := 0; i < len(f.ruleFields); i++ {
		field := f.ruleFields[i]
		if len(field.requiredFields) != 0 {
			f.setAssocFieldIndex(field)
		}
	}
}

// 设置关联的字段索引，
// TODO: 放到具体的验证函数里面去取值
func (f *StructRule) setAssocFieldIndex(fieldRule *FieldRules) {
	requiredFields := fieldRule.requiredFields
	typ, _ := isStructType(f.typ)

	if fieldRule.requiredFieldsIndex == nil {
		fieldRule.requiredFieldsIndex = make(map[string]int)
	}
	for _, name := range requiredFields {
		field, ok := typ.FieldByName(name)
		if ok {
			// 如果字段存在，设置索引
			fieldRule.setAssocFieldIndex(name, field.Index[0])
		}
	}
}
