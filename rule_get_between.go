package vrule

import (
	"fmt"
	"reflect"

	"github.com/wln32/vrule/ruleimpl"

	"github.com/gogf/gf/v2/util/gconv"
)

// 格式: between:min,max
// 说明：参数大小为min到max之间(支持整形和浮点类型参数)。
// min和max必须是具体的数字，比如1 2 3 4 ，不能是一个变量
func getBetweenRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	min := ruleVals[0]
	max := ruleVals[1]

	switch f.typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &ruleimpl.BetweenRuleNumber[int64]{
			Min:              gconv.Int64(min),
			Max:              gconv.Int64(max),
			FieldConvertFunc: getFieldReflectConvert[int64](int64(1)),
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ruleimpl.BetweenRuleNumber[uint64]{
			Min:              gconv.Uint64(min),
			Max:              gconv.Uint64(max),
			FieldConvertFunc: getFieldReflectConvert[uint64](uint64(1)),
		}
	case reflect.Float32, reflect.Float64:
		return &ruleimpl.BetweenRuleNumber[float64]{
			Min:              gconv.Float64(min),
			Max:              gconv.Float64(max),
			FieldConvertFunc: getFieldReflectConvert[float64](float64(1)),
		}

	default:
		panicUnsupportedTypeError("get between rule func", f.typ)
	}

	return nil
}

// 格式: in:value1,value2,...
// 说明：字段值应该在value1,value2,...中（字符串匹配）
// in:value1,value2,value3
// 参数类型支持数字类型和字符串，如果字符串带有空格，需要用引号包裹起来
// TODO 支持字符串类型
func getInRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	// 获取字段类型，把val转成和字段一样的类型
	kind := f.typ.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		arr := getNumberArray[int64](f.typ.Kind(), ruleVals)
		ruleFunc := &ruleimpl.RangeRule[int64]{
			Values:           arr,
			FieldConvertFunc: getFieldReflectConvert[int64](int64(1)),
		}
		return ruleimpl.ValidFuncImpl(ruleFunc.In)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		arr := getNumberArray[uint64](f.typ.Kind(), ruleVals)
		ruleFunc := &ruleimpl.RangeRule[uint64]{
			Values:           arr,
			FieldConvertFunc: getFieldReflectConvert[uint64](uint64(1)),
		}
		return ruleimpl.ValidFuncImpl(ruleFunc.In)
	case reflect.Float32, reflect.Float64:
		arr := getNumberArray[float64](f.typ.Kind(), ruleVals)
		ruleFunc := &ruleimpl.RangeRule[float64]{
			Values:           arr,
			FieldConvertFunc: getFieldReflectConvert[float64](float64(1)),
		}
		return ruleimpl.ValidFuncImpl(ruleFunc.In)
	default:
		panicUnsupportedTypeError("get in rule func", f.typ)
	}
	return nil
}

// 格式: not-in:value1,value2,...
// 说明：字段值不应该在value1,value2,...中（字符串匹配）。
// not-in:value1,value2,value3
// 参数类型支持数字类型和字符串，如果字符串带有空格，需要用引号包裹起来
// TODO 支持字符串类型
func getNotInRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	kind := f.typ.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		arr := getNumberArray[int64](f.typ.Kind(), ruleVals)
		ruleFunc := &ruleimpl.RangeRule[int64]{
			Values:           arr,
			FieldConvertFunc: getFieldReflectConvert[int64](int64(1)),
		}
		return ruleimpl.ValidFuncImpl(ruleFunc.NotIn)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		arr := getNumberArray[uint64](f.typ.Kind(), ruleVals)
		ruleFunc := &ruleimpl.RangeRule[uint64]{
			Values:           arr,
			FieldConvertFunc: getFieldReflectConvert[uint64](uint64(1)),
		}
		return ruleimpl.ValidFuncImpl(ruleFunc.NotIn)
	case reflect.Float32, reflect.Float64:
		arr := getNumberArray[float64](f.typ.Kind(), ruleVals)
		ruleFunc := &ruleimpl.RangeRule[float64]{
			Values:           arr,
			FieldConvertFunc: getFieldReflectConvert[float64](float64(1)),
		}
		return ruleimpl.ValidFuncImpl(ruleFunc.NotIn)

	default:
		panicUnsupportedTypeError("get not-in rule func", f.typ)
	}
	return nil
}

func getNumberArray[T ruleimpl.Number](kind reflect.Kind, vals []string) (a []T) {

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		arrs := gconv.Int64s(vals)
		arr := make([]T, len(vals))
		for i := 0; i < len(arrs); i++ {
			arr[i] = T(arrs[i])
		}
		return arr

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		arrs := gconv.Uint64s(vals)
		arr := make([]T, len(vals))
		for i := 0; i < len(arrs); i++ {
			arr[i] = T(arrs[i])
		}
		return arr
	case reflect.Float32, reflect.Float64:
		arrs := gconv.Float64s(vals)
		arr := make([]T, len(vals))
		for i := 0; i < len(arrs); i++ {
			arr[i] = T(arrs[i])
		}
		return arr
	default:

		panic(fmt.Errorf("getNumberArray: Unsupported parameter type: %s", kind))
	}

}
