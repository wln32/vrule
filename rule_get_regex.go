package vrule

import (
	"github.com/wln32/vrule/ruleimpl"
	"regexp"
)

// regex:pattern
func getRegexRuleMatchFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	re := &ruleimpl.RegexMatch{
		Pattern: regexp.MustCompile(vals[0]),
	}
	return ruleimpl.ValidFuncImpl(re.RegexRuleMatch)
}

// not-regex:pattern
func getRegexRuleNoMatchFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	re := &ruleimpl.RegexMatch{
		Pattern: regexp.MustCompile(vals[0]),
	}
	return ruleimpl.ValidFuncImpl(re.RegexRuleNoMatch)
}
