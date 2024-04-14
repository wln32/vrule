package vrule

import (
	"reflect"
	"strings"

	ruleimpl "github.com/wln32/vrule/ruleimpl"
)

type FieldKind int

const (
	BasicFiled FieldKind = iota + 1
	// []struct []*struct
	SliceFiled
	// map[k]struct map[k]*struct
	MapField
	StructrField
	// time gtime
	TimeField
	// []int []string
	sliceBasicFiled
	// map[k]int
	mapBasicFiled
)

type FieldRules struct {
	// v:"[LongName@]RuleArray[#msg]"
	Name string

	// 按 | 解析的rule和msg
	// rule=>[value1,value2]   example: length:6,16 => length=>{6,16}
	RuleArray map[string][]string
	// rule=> msg
	MsgArray map[string]string

	// 字段类型
	Type reflect.Type

	FieldName string
	// 字段在结构体中的索引，反射的时候直接根据索引获取
	FieldIndex int

	kind FieldKind

	// fields 如果kind 是struct才有效，其他的无效
	StructRule *StructRule

	// 关联校验的字段名，用于required-xxx以及比较系列规则
	requiredFields []string
	// 关联字段的名字=》索引，快速访问
	// FieldName => filedIndex
	requiredFieldsIndex map[string]int

	// 指针类型 requiredPtr
	// slice map array  requiredLen
	// func(ctx context.Context,in fieldValidFuncInput)
	Funcs map[string]ruleimpl.ValidFunc
}

func (f *FieldRules) removeInvalidTimeRule() {
	// 写一个闭包函数，用来判断是不是时间类型，如果是，就把对应的规则和消息删除掉
	deleteTimeRuleAndMsg := func(ruleName string) {
		_, ok := f.hasRule(ruleName)
		if ok {
			delete(f.RuleArray, ruleName)
			// 错误提示也需要删除
			_, ok := f.hasMsg(ruleName)
			if ok {
				delete(f.MsgArray, ruleName)
			}
		}
	}
	// 如果是date，datetime，date-format，规则的，并且字段类型是time.Time,gtime.Time的不用验证了
	deleteTimeRuleAndMsg(ruleimpl.Date)
	deleteTimeRuleAndMsg(ruleimpl.DateTime)
	deleteTimeRuleAndMsg(ruleimpl.DataFormat)
}

// 删掉当前字段的无效规则
func (f *FieldRules) removeInvalidRules(typ reflect.Type) {

	switch typ.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Ptr:
		// 不处理
	default:
		// 如果是其他值类型，需要查看是否有required规则，如果有，直接去除验证
		for ruleName, _ := range f.RuleArray {
			switch {
			case strings.Contains(ruleName, "required"):
				delete(f.RuleArray, ruleName)
				// 错误提示也需要删除
				_, ok := f.hasMsg(ruleName)
				if ok {
					delete(f.MsgArray, ruleName)
				}
			}
		}
	}
}

// 设置关联字段的索引
func (f *FieldRules) setAssocFieldIndex(name string, index int) {
	f.requiredFieldsIndex[name] = index
}

/*
	required-if:field,value,...
	required-unless:field,value,...
	required-with:field1,field2,...
	required-with-all:field1,field2,...
	required-without:field1,field2,...
	required-without-all:field1,field2,...

	lte: field √
	lt: field  √
	gte: field  √
	gt: field   √
*/
// 设置关联规则的关联字段
// 比如 lte: field，的类型如果是intxx系列的话，需要把field转成int64，
// 如果依赖的是值类型的字段，可以把这个校验字段删掉，因为总是有值，这个在后面直接做了返回nil代表永远有值
func (f *FieldRules) setAssocFields(structName string) {
	var fields []string
	for ruleName, ruleVals := range f.RuleArray {

		index := strings.Index(ruleName, ruleimpl.Required)
		if index != -1 {
			fields = append(fields, f.checkRequiredRulesIsValid(structName, ruleName, ruleVals)...)
		} else {
			// 主要是cmp系列的规则
			_, ok := relatedRuleNameMap[ruleName]
			if ok {
				fields = append(fields, f.checkAssocRulesIsValid(structName, ruleName, ruleVals)...)
			}
		}
		// TODO: 支持自定义解析规则的函数
	}
	if len(fields) > 0 {
		f.requiredFields = fields
	}

}

func (f *FieldRules) FieldType() reflect.Type {
	return f.Type
}
