package vrule

import (
	"regexp"

	"github.com/wln32/vrule/ruleimpl"
)

// regex:pattern
func getRegexRuleMatchFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Pattern(f, ruleimpl.RegexRuleName, vals[0])
	re := &ruleimpl.RegexMatch{
		Pattern: regexp.MustCompile(vals[0]),
	}
	return ruleimpl.ValidFuncImpl(re.RegexRuleMatch)
}

// not-regex:pattern
func getRegexRuleNoMatchFunc(s *StructRule, f *FieldRules, vals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Pattern(f, ruleimpl.NotRegexRuleName, vals[0])
	re := &ruleimpl.RegexMatch{
		Pattern: regexp.MustCompile(vals[0]),
	}
	return ruleimpl.ValidFuncImpl(re.RegexRuleNoMatch)
}
