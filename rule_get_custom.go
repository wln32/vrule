package vrule

import (
	"github.com/wln32/vrule/ruleimpl"
)

// 支持参数： 比如自定义规则
func getCustomValidRuleFunc(s *StructRule, f *FieldRules, ruleName string, vals []string) ruleimpl.ValidFunc {
	fn, ok := isCustomRuleFunc(ruleName)
	if !ok {
		return nil
	}

	vfn := &ruleimpl.RegisterCustomRuleFunc{
		Args:      vals[0],
		RuleName:  ruleName,
		FieldName: f.fieldName,
		FieldType: f.typ,
		Fn:        fn,
	}

	return ruleimpl.ValidFuncImpl(vfn.RunWithError)
}
