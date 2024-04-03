package vrule

import (
	"fmt"
	"reflect"
)

// 解析规则时，参数数量不对
func panicRuleParameterError(structName, fieldName, rule string, want int, but int) {
	panic(
		fmt.Errorf("%s.%s Rule `%s` parameter error, wanted %d but only had %d",
			structName, fieldName, rule, want, but))
}

// 解析规则时，参数数量不对
func panicRuleParameterWithMsgError(structName, fieldName, rule string, want, but string) {
	panic(
		fmt.Errorf("%s.%s Rule `%s` parameter error, wanted %s but only had %s",
			structName, fieldName, rule, want, but))
}

// 无效的规则
func panicInvalidRuleError(structName, fieldName, rule string) {
	panic(
		fmt.Errorf("%s.%s  Invalid rule `%s`", structName, fieldName, rule))
}

// 不支持的类型
func panicUnsupportedTypeError(sth string, typ reflect.Type) {
	panic(fmt.Errorf("%s，Unsupported type:%v", sth, typ))
}
