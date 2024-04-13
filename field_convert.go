package vrule

import (
	"fmt"
	"reflect"

	"github.com/wln32/vrule/ruleimpl"
)

// 输入一个FROM数字类型，转换到TO类型，奇技淫巧，不建议使用
func convertNumber[FROM, TO ruleimpl.Number]() func(any) TO {
	return func(a any) TO {
		return TO(a.(FROM))
	}
}

func getFieldReflectConvert[TO ruleimpl.Number](x any) func(reflect.Value) TO {

	switch x.(type) {
	case int64:
		return func(value reflect.Value) TO {
			return TO(value.Int())
		}
	case uint64:
		return func(value reflect.Value) TO {
			return TO(value.Uint())
		}
	case float64:
		return func(value reflect.Value) TO {
			return TO(value.Float())
		}
	default:
		panicUnsupportedTypeError("get field convert func", reflect.TypeOf(x))
	}
	return nil
}

// 获取结构体中fieldName的类型转换函数
// 主要是cmp系列的规则
func getAssocFieldTypeConvert[TO ruleimpl.Number](s *StructRule, fieldName string) func(value reflect.Value) TO {
	field, ok := s.typ.FieldByName(fieldName)
	if !ok {
		panic(fmt.Errorf("structure: %s has no fields: %s", s.longName, fieldName))
	}

	switch field.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(value reflect.Value) TO {
			return TO(value.Int())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(value reflect.Value) TO {
			return TO(value.Uint())
		}

	case reflect.Float32, reflect.Float64:
		return func(value reflect.Value) TO {
			return TO(value.Float())
		}

	}
	panicUnsupportedTypeError("get assoc field convert func", field.Type)
	return nil
}
