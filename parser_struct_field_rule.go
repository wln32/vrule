package vrule

import (
	"fmt"
	"reflect"
	"unicode"
)

func (v *Validator) parseFieldRule(fieldName, tag string, structTyp, originalTyp reflect.Type, fieldIndex int, cache *StructCache) (*FieldRules, error) {

	var f = &FieldRules{
		Type:       originalTyp,
		FieldName:  fieldName,
		FieldIndex: fieldIndex,
		kind:       BasicFiled,
	}
	if tag != "" {
		// 如果有v 或者valid标签的，防止有的字段类型是struct但是没有v 或者valid标签的
		name, rule, msg := parseTagValue(tag)
		if name == "" {
			if v.parseRuleOption.FieldNameFunc != nil {
				// 如果用户指定了字段名
				byName, _ := structTyp.FieldByName(fieldName)
				name = v.parseRuleOption.FieldNameFunc(structTyp, byName)
			}
		}
		if name == "" {
			name = fieldName
		}
		f.Name = name
		// 把rule 和msg 解析，以 | 号隔开
		f.parseRuleAndMsg(rule, msg, v)
		// 设置关联校验的字段
		f.setAssocFields(structTyp.String())
	}

	var elemFieldTyp = originalTyp
	// 判断类型是什么
	// *T -> T
	if elemFieldTyp.Kind() == reflect.Ptr {
		elemFieldTyp = elemFieldTyp.Elem()
	}

	switch elemFieldTyp.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:

		f.kind = SliceFiled
		if elemFieldTyp.Kind() == reflect.Map {
			f.kind = MapField
		}

		// []T -> T
		elemFieldTyp = elemFieldTyp.Elem()
		// *T -> T
		if elemFieldTyp.Kind() == reflect.Ptr {
			elemFieldTyp = elemFieldTyp.Elem()
		}
		if elemFieldTyp.Kind() == reflect.Struct {
			// 需要校验 字段类型是不是与当前对象类型是否相同 切片 map都得校验
			err := f.checkRecursionRequiredRuleStruct(elemFieldTyp, structTyp)
			if err != nil {
				return nil, err
			}
			f.StructRule = v.ParseStruct(elemFieldTyp, cache)
		}
	case reflect.Struct:
		// 支持time gtime, time类型不在属于结构体类型，属于基本类型
		if isTimeType(elemFieldTyp) {
			f.kind = TimeField
			// 删除掉一些无效的时间规则
			f.removeInvalidTimeRule()
		} else {
			// 需要校验 字段类型是不是与当前对象类型是否相同 切片 map都得校验
			// 如果相同，判断是否有required规则，如果有直接panic
			// 这样做的动机是 当字段类型和当前类型一样的时候，并且有required规则，
			// 会导致无限递归验证，导致失败
			/*
				type Order struct {
						Name string
						Id int
						Orders []*Order  `v:"required"` // panic
						OtherOrder *Order   `v:"required"` // panic
						MapOrder map[string]*Order   `v:"required"` // panic
				}
				// 字段类型和结构体类型相同，且当前结构体内有required字段，panic
			*/
			err := f.checkRecursionRequiredRuleStruct(elemFieldTyp, structTyp)
			if err != nil {
				return nil, err
			}
			f.StructRule = v.ParseStruct(elemFieldTyp, cache)
			f.kind = StructrField
		}

	default:
		// 当一个字段是值类型时
		// 例如required规则下的 值类型字段不需要校验，包括struct
		f.removeInvalidRules(originalTyp)
	}

	return f, nil

}

func (v *Validator) ParseStruct(a any, cache *StructCache) *StructRule {

	var typ reflect.Type
	if atype, ok := a.(reflect.Type); ok {
		typ = atype
	} else {
		typ = reflect.TypeOf(a)
		// *struct -> struct
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
	}

	// struct
	if typ.Kind() != reflect.Struct {
		panic(fmt.Errorf("must be of struct type, other types are not supported by %v", typ))
	}

	if cache != nil {
		// 判断有没有已经创建过了，如果有直接返回
		structRule := cache.GetStructRule(typ)
		if structRule != nil {
			return structRule
		}
	}

	// 递归获取，获取全部的字段，防止以下情况时没有验证的字段
	// 字段名没有v或者valid ，判断下类型是不是基础类型
	// 如果不是基础类型，需要递归
	/*/
	type Params struct {
		Age *uint `v:"min:18"`
	}
	type Req struct {
		Name   string
		Params *Params
	}
	*/
	tagFields := getStructFields(typ)
	if len(tagFields) == 0 {
		return nil
	}

	structRule := &StructRule{
		RuleFields: make([]*FieldRules, 0, 4),
		Type:       typ,
		LongName:   getStructName(typ),
		ShortName:  typ.String(),
	}

	var structTyp = typ
	var err error

	for _, field := range tagFields {

		// 判断是不是公开的字段
		if unicode.IsLower(rune(field.Field.Name[0])) {
			continue
		}
		// 判断是不是接口
		if field.Kind() == reflect.Interface {
			continue
		}

		var (
			rule *FieldRules
		)

		tagValue := field.Tag("v")
		if tagValue == "" {
			tagValue = field.Tag("valid")
		}
		// 如果当前字段有v 规则
		if tagValue != "" {
			//  用户自定义的过滤字段函数，内置默认的返回false
			// 只能过滤掉带有v或者valid规则的字段
			if v.parseRuleOption != nil {
				if v.parseRuleOption.FilterFieldFunc(typ, field.Field) {
					continue
				}
			}

			rule, err = v.parseFieldRule(field.Name(), tagValue, structTyp, field.Field.Type, field.Field.Index[0], cache)
			if err != nil {
				return nil
			}
		}
		/*
			rule.FieldIndex = field.Field.Index[0]
			// 防止内嵌结果体的字段索引
			type Pass struct {
					Pass1 string `valid:"password1@required|same:password2#请输入您的密码|您两次输入的密码不一致"`
					Pass2 string `valid:"password2@required|same:password1#请再次输入您的密码|您两次输入的密码不一致"`
				}
			type User struct {
				Id   int
				Name string `valid:"LongName@required#请输入您的姓名"`
				Pass
			}
		*/

		// 判断字段类型是不是结构体，进行递归
		// fieldTyp 只是解引用用来判断是不是结构体
		fieldTyp, ok := isStructType(field.Field.Type)
		if ok {
			/*
				type User struct {
					UsersPtr *User
					ID       *int   `v:"required" `
					Name     string `v:"required" `
				}
			*/
			// 判断字段类型和结构体类型是否相同
			// 如果相同，直接跳过，不需要添加验证的字段,或者是gtime.Time, time.Time
			if fieldTyp == structTyp {
				continue
			}
			// 不要忘记 如果时间类型的字段带有rule，就添加到ruleFields中去
			if isTimeType(fieldTyp) && rule != nil {
				structRule.RuleFields = append(structRule.RuleFields, rule)
				continue
			}

			// 这里要用原来的字段类型field.Field.Type
			rule, err = v.parseFieldRule(field.Name(), tagValue, structTyp, field.Field.Type, field.Field.Index[0], cache)
			if err != nil {
				return nil
			}
			rule.StructRule = v.ParseStruct(fieldTyp, cache)
		}

		if rule != nil {
			structRule.RuleFields = append(structRule.RuleFields, rule)
		}
	}

	/*
		type UserApiSearch247 struct {
			Uid      int64  `v:"required"`
			Nickname string `v:"required-with:uid"`
		}
	*/
	// 例如以上结构体Nickname字段的关联校验字段uid是不存在的，默认不会进行大小写转换去匹配所有字段
	// 需要判断 每个字段的关联规则的字段是否存在，以及使用索引来访问关联字段
	structRule.setIndexAssocFields()

	// 设置字段的校验函数
	structRule.setFieldsRuleValidFunc()

	// 删除掉没有rules的字段,前面已经过滤掉无效的规则了，
	structRule.deleteEmptyRuleField()

	if len(structRule.RuleFields) != 0 {
		if cache != nil {
			cache.AddStructRule(structRule)
		}
	}
	return structRule
}
