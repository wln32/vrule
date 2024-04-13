package vrule

import (
	"context"
	"strings"

	ruleimpl "github.com/wln32/vrule/ruleimpl"

	"github.com/gogf/gf/v2/text/gstr"
)

var defaultCtx = context.TODO()

// return index
// -1 =not found
func (f *FieldRules) hasRule(rule string) (vals []string, ok bool) {
	vals, ok = f.ruleArray[rule]
	return
}

func (f *FieldRules) hasMsg(rule string) (msg string, ok bool) {
	msg, ok = f.msgArray[rule]
	return
}

// [属性别名@]校验规则[#错误提示]
// 校验规则则为当前属性的校验规则，多个校验规则请使用|符号组合，例如：required|between:1,100
func (f *FieldRules) parseRuleAndMsg(ruleString string, msg string, v *Validator) {
	if f.ruleArray == nil {
		f.ruleArray = make(map[string][]string)
	}
	if f.msgArray == nil {
		f.msgArray = make(map[string]string)
	}

	rules := strings.Split(ruleString, "|")
	for i := 0; i < len(rules); i++ {
		if strings.TrimSpace(rules[i]) != "" {

			f.addRule(rules[i])
		}
	}
	// 如果没有rule，不需要下一步
	if len(f.ruleArray) == 0 {
		return
	}

	ruleToMsgMap := map[string]string{}

	// 设置默认的错误提示信息
	for rule, _ := range f.ruleArray {
		// 如果是内置规则，直接添加
		_, ok := isBuiltinRule(rule)
		if ok {
			errMsg := ""
			if v.translationOption != nil {
				// 如果用户设置了
				errMsg = v.translationOption.TranslateFunc(defaultCtx, rule)
				if errMsg == "" {
					// 如果没有，使用默认的错误提示
					errMsg = v.i18n.Translate(defaultCtx, rule)
				}
			}

			if errMsg == "" {
				// 如果没有，使用默认的错误提示
				errMsg = ruleimpl.RuleMsgMap[rule]
			}

			ruleToMsgMap[rule] = errMsg
			continue
		}
		// 自定义规则错误提示：使用自定义校验函数返回的error
		//_, ok = isCustomRuleFunc(rule)
		//if ok {
		//	ruleToMsgMap[rule] = customRuleMsgMap[rule]
		//}
	}
	// 去掉rule的冒号后面的内容，例如：max:10 -> max
	var getRuleName = func(rule string) string {
		index := strings.Index(rule, ":")
		if index == -1 {
			return rule
		}
		return rule[:index]
	}

	// msg没有值时，长度总是为1
	msgs := strings.Split(msg, "|")
	for i, msg := range msgs {
		// 如果用户设置了错误提示
		if msg != "" {
			ruleName := getRuleName(rules[i])
			// 直接替换
			ruleToMsgMap[ruleName] = msg
		}
	}

	// 最后一步，替换占位符中的值
	for ruleName, msg := range ruleToMsgMap {
		f.addMsg(ruleName, msg)
	}

	return
}

// length: 6,16
// required-if: field1,value1,field2,value2,...
func (f *FieldRules) addRule(rule string) {

	var vals []string

	index := strings.Index(rule, ":")

	if index == -1 { // required json integer float之类的 没有参数
		f.ruleArray[rule] = vals
		return
	}
	ruleName := rule[:index]
	// 如果是内置规则，直接添加
	_, ok := isBuiltinRule(ruleName)
	if ok {
		// TODO 参数的解析可以放到注册规则那里去
		ruleVal := rule[index+1:]
		// 如果是正则表达式，直接添加
		// regex : \\d{4,18} 正则表达式不需要split
		if ruleName == ruleimpl.RegexRuleName || ruleName == ruleimpl.NotRegexRuleName {
			f.ruleArray[ruleName] = []string{ruleVal}
			return
		}
		//  日期之类的date-format: Y-m-d H:i:s
		if strings.Contains(ruleName, "date") {
			f.ruleArray[ruleName] = []string{ruleVal}
			return
		}
		// 其他正常的参数
		vals = strings.Split(ruleVal, ",")
		f.ruleArray[ruleName] = vals
		return
	}
	// 如果是自定义规则
	_, ok = isCustomRuleFunc(ruleName)
	if ok {
		args := rule[index+1:]

		f.ruleArray[ruleName] = []string{args}

		return
	}
	panicInvalidRuleError("", f.fieldName, ruleName)

}

// `v:"uid      @integer|min:1# |请输入用户ID"`
// 需要注意的是，这种规则下，integer的错误信息使用默认的，
// min的错误信息使用自定义的
// `v:"between:1,10000 #project id must between {min}, {max}"`
// 需要把min，max 替换成1 10000这种
// TODO： i18n
func (f *FieldRules) addMsg(rule, msg string) {

	if gstr.Contains(msg, "{") {
		msg = gstr.ReplaceByMap(msg, map[string]string{
			"{field}": f.name, // Field longName of the `value`.
			//"{value}":    需要再验证的时候再替换
			//"{pattern}":    strings.Join(f.ruleArray[rule], ","), // 正则和not-in， in
			//
			//"{customrule}": rule, // 用户自定义的规则名

		})
		// msg = f.replaceMinMaxMsg(msg)
	}
	ruleName := strings.Split(rule, ":")[0]
	f.msgArray[ruleName] = msg
}

func (f *FieldRules) hasMinAndMaxMsgWithRule() (int, bool) {
	const (
		Invalid = iota - 1
		Between
		Length
		Max
		Min
	)
	var hasMinAndMax = func(ruleName string) bool {
		_, ok := f.hasRule(ruleName)
		return ok
	}
	if hasMinAndMax(ruleimpl.Between) {
		return Between, true
	}
	if hasMinAndMax(ruleimpl.Length) {
		return Length, true
	}
	if hasMinAndMax(ruleimpl.Max) {
		return Max, true
	}
	if hasMinAndMax(ruleimpl.Min) {
		return Min, true
	}
	return 0, false
}

// 替换错误提示信息里面有{min}和{max}的信息
func (f *FieldRules) replaceMinMaxMsg(msg string) string {
	kind, ok := f.hasMinAndMaxMsgWithRule()
	if !ok {
		return msg
	}
	// n = 0 max min
	const (
		Between = iota
		Length
		Max
		Min
	)
	switch kind {
	case Max:
		vals, _ := f.hasRule(ruleimpl.Max)
		msg = gstr.ReplaceByMap(msg, map[string]string{
			"{max}": vals[0],
		})
	case Min:
		vals, _ := f.hasRule(ruleimpl.Min)
		msg = gstr.ReplaceByMap(msg, map[string]string{
			"{min}": vals[0],
		})
	case Between:
		vals, _ := f.hasRule(ruleimpl.Between)
		if len(vals) != 2 {
			panicRuleParameterError("", f.name, ruleimpl.Between, 2, len(vals))
		}
		msg = gstr.ReplaceByMap(msg, map[string]string{
			"{min}": vals[0],
			"{max}": vals[1],
		})
	case Length:
		vals, _ := f.hasRule(ruleimpl.Length)
		if len(vals) != 2 {
			panicRuleParameterError("", f.name, ruleimpl.Length, 2, len(vals))
		}
		msg = gstr.ReplaceByMap(msg, map[string]string{
			"{min}": vals[0],
			"{max}": vals[1],
		})
	}

	return msg
}
