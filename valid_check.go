package vrule

import (
	"context"
	"reflect"

	"github.com/wln32/vrule/ruleimpl"
)

// struct: field*
func (f *StructRule) checkStructByName(ctx context.Context, ptr any) error {
	structPtr := reflect.ValueOf(ptr)

	// TODO： 用来验证无限递归required，可以在注册阶段就报错

	// 如果给的值是空的，不需要校验
	// 这里有可能是checkBasic进来的
	// 如果是 说明字段校验过了，但是还得验证结构体的规则
	// 就比如以下情况，b字段是满足规则的
	// 但是b是结构体，有可能B结构体里面也有v规则字段，
	// 需要递归校验
	// type Example struct {
	//	a *A `v:"required-without:b"`
	//	b *B `v:"required-without:a"`
	// }
	// foo := &Example{
	//	A: &A{aa:1,bb:2},
	// }
	// checkStruct(ctx,foo)
	// 如果通过校验，（**struct | *struct | struct ） -> struct
	// 必须要解引用，不然panic
	structVal, ok := valueIsValid(structPtr)
	if !ok {
		return nil
	}

	errors := Errors{}

	for i := 0; i < len(f.RuleFields); i++ {
		field := f.RuleFields[i]

		fieldVal := structVal.FieldByName(field.FieldName)

		err := field.checkStructField(ctx, fieldVal, structVal)

		if err != nil {
			errors = append(errors, err)

			return errors

		}
	}

	return errors
}

// struct: field*
func (f *StructRule) checkStruct(ctx context.Context, a any) error {
	var structPtr reflect.Value
	val, ok := a.(reflect.Value)
	if ok {
		structPtr = val
	} else {
		structPtr = reflect.ValueOf(a)
	}

	// 如果给的值是空的，不需要校验
	// 这里有可能是checkBasic进来的
	// 如果是 说明字段校验过了，但是还得验证结构体的规则
	// 就比如以下情况，b字段是满足规则的
	// 但是b是结构体，有可能B结构体里面也有v规则字段，
	// 需要递归校验
	/*
		type StructA struct {
			aa int `v:"required"`
			bb int `v:"required"`
		}
		type Example struct {
			A StructA
			B string `v:"required-without:a"`
		}
		foo := &Example{
			A: &StructA{aa:1,bb:2},
		}
	*/
	// 如果通过校验，（**struct | *struct | struct ） -> struct
	// 是不是空指针
	structVal, ok := valueIsValid(structPtr)
	if !ok {
		return nil
	}

	errs := Errors{}

	for i := 0; i < len(f.RuleFields); i++ {
		field := f.RuleFields[i]

		fieldVal := structVal.Field(field.FieldIndex)

		err := field.checkStructField(ctx, fieldVal, structVal)

		if err != nil {
			// TODO：后续可以为error加上 当前结构体的名字
			// 比如 struct.field.error
			errs = append(errs, err)

		}
	}

	return errs
}

// field: basic | struct | slice | map
func (f *FieldRules) checkStructField(ctx context.Context, fieldVal reflect.Value, structPtr reflect.Value) error {
	//type User struct {
	//	Name string `v:"required"`
	//  Address []string `v:"required"`
	//  Orders []*Orders `v:"required"`
	//}
	// required Address
	// 需要校验Orders 是否满足规则，
	if len(f.RuleArray) != 0 {
		//var requiredData map[string]any
		// 设置关联校验的字段
		//if len(f.requiredFieldsIndex) != 0 {
		//	requiredData = make(map[string]any, len(f.requiredFieldsIndex))
		//	for Name, index := range f.requiredFieldsIndex {
		//		requiredFieldVal := structPtr.Field(index)
		//		requiredData[Name] = requiredFieldVal.Interface()
		//	}
		//}

		err := f.checkBasic(ctx, fieldVal, structPtr)
		if err != nil {
			return err
		}

	}

	switch f.kind {
	case BasicFiled, TimeField:
		return nil
	case SliceFiled:

		return f.checkSlice(ctx, fieldVal)
	case MapField:

		return f.checkMap(ctx, fieldVal)
	case StructrField:
		/*
			type Pass struct {
					Pass1 string `v:"required|same:Pass2"`
					Pass2 string `v:"required|same:Pass1"`
				}
				type User struct {
					Id   int
					Name string `v:"required"`
					Pass Pass
				}
				user := &User{
					Name: "",
					Pass: Pass{
						Pass1: "1",
						Pass2: "2",
					},
				}
			不合并错误会变成如下
			User error {
				Name：[The Name field is required]

					Pass struct error:
						[The Pass1 value `1` must be the same as field value `` The Pass2 value `2` must be the same as field value ``]

				Pass：[The Name field is required， The Pass1 value `1` must be the same as field value ``The Pass2 value `2` must be the same as field value ``]
			}
			可以看到最后两个合并被合并到一个错误上了
		*/
		// TODO: 增加一个错误合并的
		errs := f.StructRule.checkStruct(ctx, fieldVal)
		return errs
	}
	return nil
}

// slice: ( basic | struct)*
func (f *FieldRules) checkSlice(ctx context.Context, arr reflect.Value) error {
	// struct
	for i := 0; i < arr.Len(); i++ {
		item := arr.Index(i)
		err := f.StructRule.checkStruct(ctx, item)
		if err != nil {
			// bail
			return err
		}
	}

	return nil
}

// map[k]v: v=  ( basic | struct)*
func (f *FieldRules) checkMap(ctx context.Context, map_ reflect.Value) error {
	mapKeys := map_.MapKeys()
	// struct
	for _, key := range mapKeys {
		value := map_.MapIndex(key)
		err := f.StructRule.checkStruct(ctx, value)
		if err != nil {
			// bail
			return err
		}
	}
	return nil
}

// basic: int string bool float .... All golang built-in basic types
func (f *FieldRules) checkBasic(ctx context.Context, val reflect.Value, StructPtr reflect.Value) error {
	if len(f.Funcs) == 0 {
		return nil
	}
	var err error
	for ruleName, fn := range f.Funcs {
		// fieldVal := val.Interface()
		err = fn.Run(ctx, ruleimpl.RuleFuncInput{
			Value: val,

			StructPtr: StructPtr,
			Message:   f.MsgArray[ruleName],
		})
		// bail
		if err != nil {
			return err
		}
	}
	return nil
}
