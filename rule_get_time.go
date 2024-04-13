package vrule

import (
	"github.com/wln32/vrule/ruleimpl"
)

// date
func getDateRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	return ruleimpl.ValidFuncImpl(ruleimpl.DateRuleFunc)
}

// datetime
func getDateTimeRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	return ruleimpl.ValidFuncImpl(ruleimpl.DateTimeRuleFunc)
}

// date-format:format
func getDateTimeFormat(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Pattern(f, ruleimpl.DataFormat, vals[0])
	// 校验format是否合法
	return &ruleimpl.DateFormatRule{
		Format: vals[0],
	}
}

// field可能是string 或者time
// before:field
func getBeforeTimeRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.Before, vals[0])
	ruleFunc := &ruleimpl.TimeRule{
		FieldName:       vals[0],
		AssocFieldIndex: f.requiredFieldsIndex[vals[0]],
	}
	return ruleimpl.ValidFuncImpl(ruleFunc.Before)
}

// field可能是string 或者time
// before-equal:field
func getBeforeEqualTimeRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.BeforeEqual, vals[0])
	ruleFunc := &ruleimpl.TimeRule{
		FieldName:       vals[0],
		AssocFieldIndex: f.requiredFieldsIndex[vals[0]],
	}
	return ruleimpl.ValidFuncImpl(ruleFunc.BeforeEqual)
}

// field可能是string 或者time
// after:field
func getAfterTimeRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.After, vals[0])
	ruleFunc := &ruleimpl.TimeRule{
		FieldName:       vals[0],
		AssocFieldIndex: f.requiredFieldsIndex[vals[0]],
	}
	return ruleimpl.ValidFuncImpl(ruleFunc.After)
}

// field可能是string 或者time
// after-equal:field
func getAfterEqualTimeRuleFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Field1(f, ruleimpl.AfterEqual, vals[0])

	ruleFunc := &ruleimpl.TimeRule{
		FieldName:       vals[0],
		AssocFieldIndex: f.requiredFieldsIndex[vals[0]],
	}
	return ruleimpl.ValidFuncImpl(ruleFunc.AfterEqual)
}
