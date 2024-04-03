package vrule

import (
	"github.com/wln32/vrule/ruleimpl"
	"reflect"

	"github.com/gogf/gf/v2/util/gconv"
)

// 格式: max:max
// 说明：参数大小最大为max(支持整形和浮点类型参数)。
func getMaxRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {

	max := ruleVals[0]

	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.MaxRuleNumber[int64]{
			Max:              gconv.Int64(max),
			FieldConvertFunc: getFieldTypeConvert[int64](f),
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.MaxRuleNumber[uint64]{
			Max:              gconv.Uint64(max),
			FieldConvertFunc: getFieldTypeConvert[uint64](f),
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.MaxRuleNumber[float64]{
			Max:              gconv.Float64(max),
			FieldConvertFunc: getFieldTypeConvert[float64](f),
		}
	default:
		panicUnsupportedTypeError("get max rule func", f.typ)
	}

	return nil
}

// 格式: min:min
// 说明：参数大小最小为min(支持整形和浮点类型参数)。
func getMinRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {

	min := ruleVals[0]

	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.MinRuleNumber[int64]{
			Min:              gconv.Int64(min),
			FieldConvertFunc: getFieldTypeConvert[int64](f),
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.MinRuleNumber[uint64]{
			Min:              gconv.Uint64(min),
			FieldConvertFunc: getFieldTypeConvert[uint64](f),
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.MinRuleNumber[float64]{
			Min:              gconv.Float64(min),
			FieldConvertFunc: getFieldTypeConvert[float64](f),
		}
	default:
		panicUnsupportedTypeError("get min rule func", f.typ)
	}

	return nil
}
