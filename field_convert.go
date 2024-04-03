package vrule

import (
	"fmt"
	"github.com/wln32/vrule/ruleimpl"
	"reflect"
)

// 输入一个FROM数字类型，转换到TO类型，奇技淫巧，不建议使用
func convertNumber[FROM, TO ruleimpl.Number]() func(any) TO {
	return func(a any) TO {
		return TO(a.(FROM))
	}
}

func defaultFieldValueConvert(a any) any {
	return a
}

// 转换到TO类型
func getFieldTypeConvert[TO ruleimpl.Number](f *FieldRules) func(any) TO {
	if f == nil {
		return nil
	}
	switch f.typ.Kind() {
	//==========================================================
	case reflect.Int:
		return convertNumber[int, TO]()

	case reflect.Int8:
		return convertNumber[int8, TO]()

	case reflect.Int16:
		return convertNumber[int16, TO]()

	case reflect.Int32:
		return convertNumber[int32, TO]()
	case reflect.Int64:
		return convertNumber[int64, TO]()
		//==========================================================
	case reflect.Uint:
		return convertNumber[uint, TO]()

	case reflect.Uint8:
		return convertNumber[uint8, TO]()

	case reflect.Uint16:
		return convertNumber[uint16, TO]()

	case reflect.Uint32:
		return convertNumber[uint32, TO]()
	case reflect.Uint64:
		return convertNumber[uint64, TO]()
		//==========================================================
	case reflect.Float32:
		return convertNumber[float32, TO]()
	case reflect.Float64:
		return convertNumber[float64, TO]()

	}
	panicUnsupportedTypeError("get field convert func", f.typ)
	return nil
}

// 获取结构体中fieldName的类型转换函数
func getAssocFieldTypeConvert[TO ruleimpl.Number](s *StructRule, fieldName string) func(any) TO {
	field, ok := s.typ.FieldByName(fieldName)
	if !ok {
		panic(fmt.Errorf("structure: %s has no fields: %s", s.longName, fieldName))
	}

	switch field.Type.Kind() {
	case reflect.Int:
		return convertNumber[int, TO]()
	case reflect.Int8:
		return convertNumber[int8, TO]()
	case reflect.Int16:
		return convertNumber[int16, TO]()
	case reflect.Int32:
		return convertNumber[int32, TO]()
	case reflect.Int64:
		return convertNumber[int64, TO]()
		//==========================================================
	case reflect.Uint:
		return convertNumber[uint, TO]()
	case reflect.Uint8:
		return convertNumber[uint8, TO]()
	case reflect.Uint16:
		return convertNumber[uint16, TO]()
	case reflect.Uint32:
		return convertNumber[uint32, TO]()
	case reflect.Uint64:
		return convertNumber[uint64, TO]()
		//==========================================================
	case reflect.Float32:
		return convertNumber[float32, TO]()
	case reflect.Float64:
		return convertNumber[float64, TO]()

	}
	panicUnsupportedTypeError("get assoc field convert func", field.Type)
	return nil
}
