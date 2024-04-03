package vrule

import (
	"fmt"
	"github.com/wln32/vrule/ruleimpl"
	"strings"
	"sync"
)

// 校验时rule
type RegisterCustomRuleOption struct {
	// 规则名称
	RuleName string
	// 自定义规则的验证函数
	Fn ruleimpl.CustomValidRuleFunc
}

var (
	customFuncMutex      sync.Mutex
	defaultCustomRuleMsg = "{field} does not satisfy the condition {customrule}"
)

var customRuleMsgMap = map[string]string{}

var customValidRuleFuncMap = map[string]ruleimpl.CustomValidRuleFunc{}

var registerCustomRuleInvalidOptionError = fmt.Errorf("invalid custom rule option")

var registerCustomRuleInvalidParameterError = fmt.Errorf("registered verification function or rule name cannot be empty")

var registerCustomRuleExistsError = "rule:`%s` you registered already exists and cannot be registered again"

func RegisterCustomRuleFunc(option RegisterCustomRuleOption) error {

	ruleName := strings.TrimSpace(option.RuleName)
	// 无效的
	if option.Fn == nil || ruleName == "" {
		panic(registerCustomRuleInvalidParameterError)
	}
	// 不能和内置的规则重名
	if _, ok := isBuiltinRule(ruleName); ok {
		panic(fmt.Errorf(registerCustomRuleExistsError, ruleName))
	}
	// 不能和自定义规则重名
	if _, ok := isCustomRuleFunc(ruleName); ok {
		panic(fmt.Errorf(registerCustomRuleExistsError, ruleName))
	}

	customFuncMutex.Lock()
	defer customFuncMutex.Unlock()

	customValidRuleFuncMap[ruleName] = option.Fn
	// customRuleMsgMap[ruleName] = defaultCustomRuleMsg

	return nil
}

// 让用户自定义规则的解析函数
func RegisterCustomRuleFuncWithParse(ruleName string, fn func(structRule *StructRule, fieldRule *FieldRules, args []string) ruleimpl.ValidFunc) error {

	ruleName = strings.TrimSpace(ruleName)
	// 无效的
	if fn == nil || ruleName == "" {
		panic(registerCustomRuleInvalidParameterError)
	}
	// 不能和内置的规则重名
	if _, ok := isBuiltinRule(ruleName); ok {
		panic(fmt.Errorf(registerCustomRuleExistsError, ruleName))
	}
	// 不能和自定义规则重名
	if _, ok := isCustomRuleFunc(ruleName); ok {
		panic(fmt.Errorf(registerCustomRuleExistsError, ruleName))
	}

	customFuncMutex.Lock()
	defer customFuncMutex.Unlock()

	builtinRulesMapToFunc[ruleName] = fn

	return nil
}

func isCustomRuleFunc(name string) (ruleimpl.CustomValidRuleFunc, bool) {
	rule, ok := customValidRuleFuncMap[name]
	return rule, ok
}
