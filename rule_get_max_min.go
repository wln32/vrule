package vrule

import (
	"reflect"

	"github.com/wln32/vrule/ruleimpl"

	"github.com/gogf/gf/v2/util/gconv"
)

// 格式: max:max
// 说明：字段的值最大为max(支持整形和浮点类型参数)。value<=max
// max必须是一个数字，比如1 2 3，不能是一个变量
func getMaxRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Max(f, ruleimpl.Max, ruleVals[0])
	max := ruleVals[0]

	switch f.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.MaxRuleNumber[int64]{
			Max:              gconv.Int64(max),
			FieldConvertFunc: getFieldReflectConvert[int64](int64(1)),
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.MaxRuleNumber[uint64]{
			Max:              gconv.Uint64(max),
			FieldConvertFunc: getFieldReflectConvert[uint64](uint64(1)),
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.MaxRuleNumber[float64]{
			Max:              gconv.Float64(max),
			FieldConvertFunc: getFieldReflectConvert[float64](float64(1)),
		}
	default:
		panicUnsupportedTypeError("get max rule func", f.Type)
	}

	return nil
}

// 格式: min:min
// 说明：字段的值最小为min(支持整形和浮点类型参数)， value>=min
// min必须是一个数字，比如1 2 3，不能是一个变量
func getMinRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Min(f, ruleimpl.Min, ruleVals[0])
	min := ruleVals[0]

	switch f.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return &ruleimpl.MinRuleNumber[int64]{
			Min:              gconv.Int64(min),
			FieldConvertFunc: getFieldReflectConvert[int64](int64(1)),
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.MinRuleNumber[uint64]{
			Min:              gconv.Uint64(min),
			FieldConvertFunc: getFieldReflectConvert[uint64](uint64(1)),
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.MinRuleNumber[float64]{
			Min:              gconv.Float64(min),
			FieldConvertFunc: getFieldReflectConvert[float64](float64(1)),
		}
	default:
		panicUnsupportedTypeError("get min rule func", f.Type)
	}

	return nil
}
