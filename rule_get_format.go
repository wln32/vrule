package vrule

import (
	"github.com/wln32/vrule/ruleimpl"
)

func getFormatJsonRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {

	return ruleimpl.ValidFuncImpl(ruleimpl.JsonFormat)

}

func getFormatBooleanRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {

	return ruleimpl.ValidFuncImpl(ruleimpl.BooleanFormat)

}

func getFormatIntegerRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {

	return ruleimpl.ValidFuncImpl(ruleimpl.IntegerFormat)

}

func getFormatFloatRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {

	return ruleimpl.ValidFuncImpl(ruleimpl.FloatFormat)

}
