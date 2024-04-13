package vrule

import (
	"sync"

	ruleimpl "github.com/wln32/vrule/ruleimpl"
)

var once sync.Once

type registerRuleFunc func(structRule *StructRule, fieldRule *FieldRules, s []string) ruleimpl.ValidFunc

var builtinRulesMapToFunc = map[string]registerRuleFunc{
	ruleimpl.Required: getRequiredRuleFunc,
	//=================================================
	ruleimpl.RequiredIf:     getRequiredIfRuleFunc,
	ruleimpl.RequiredUnless: getRequiredUnlessRuleFunc,
	//=================================================
	ruleimpl.RequiredWith:       getRequiredWithRuleFunc,
	ruleimpl.RequiredWithAll:    getRequiredWithAllRuleFunc,
	ruleimpl.RequiredWithout:    getRequiredWithoutRuleFunc,
	ruleimpl.RequiredWithoutAll: getRequiredWithoutAllRuleFunc,
	//==============string====================================
	ruleimpl.Size:      getStringSizeRuleFunc,
	ruleimpl.Length:    getStringLengthRuleFunc,
	ruleimpl.MinLength: getStringMinLengthRuleFunc,
	ruleimpl.MaxLength: getStringMaxLengthRuleFunc,
	//================cmp====================================
	ruleimpl.Between: getBetweenRuleFunc,
	ruleimpl.Max:     getMaxRuleFunc,
	ruleimpl.Min:     getMinRuleFunc,
	ruleimpl.In:      getInRuleFunc,
	ruleimpl.NotIn:   getNotInRuleFunc,
	//================cmp===============================
	ruleimpl.Lte:       getLteRuleFunc,
	ruleimpl.Lt:        getLtRuleFunc,
	ruleimpl.Gte:       getGteRuleFunc,
	ruleimpl.Gt:        getGtRuleFunc,
	ruleimpl.NotEq:     getNotEqRuleFunc,
	ruleimpl.Eq:        getEqRuleFunc,
	ruleimpl.Different: getDifferentRuleFunc,
	ruleimpl.Same:      getSameRuleFunc,
	//==================regex===============================
	ruleimpl.RegexRuleName:    getRegexRuleMatchFunc,
	ruleimpl.NotRegexRuleName: getRegexRuleNoMatchFunc,
	//===============date==================================
	ruleimpl.Date:     getDateRuleFunc,
	ruleimpl.DateTime: getDateTimeRuleFunc,

	ruleimpl.Before:      getBeforeTimeRuleFunc,
	ruleimpl.BeforeEqual: getBeforeEqualTimeRuleFunc,
	ruleimpl.After:       getAfterTimeRuleFunc,
	ruleimpl.AfterEqual:  getAfterEqualTimeRuleFunc,

	ruleimpl.DataFormat: getDateTimeFormat,
	//=================format===================================
	ruleimpl.JsonRuleName:    getFormatJsonRuleFunc,
	ruleimpl.BooleanRuleName: getFormatBooleanRuleFunc,
	ruleimpl.IntegerRuleName: getFormatIntegerRuleFunc,
	ruleimpl.FloatRuleName:   getFormatFloatRuleFunc,
	//=================================================
	//ruleimpl.RegexQQ:         getRegexRuleMatchQQFunc,
	//ruleimpl.RegexDomain:     getRegexRuleMatchDomainFunc,
	//ruleimpl.RegexUrl:        getRegexRuleMatchUrlFunc,
	//ruleimpl.RegexMAC:        getRegexRuleMatchMacFunc,
	//ruleimpl.RegexIP:         getRegexRuleMatchIpFunc,
	//ruleimpl.RegexIPV4:       getRegexRuleMatchIpv4Func,
	//ruleimpl.RegexIPV6:       getRegexRuleMatchIpv6Func,
	//ruleimpl.RegexBankCard:   getRegexRuleMatchBankCardFunc,
	//ruleimpl.RegexResidentID: getRegexRuleMatchResidentIdFunc,
	//ruleimpl.RegexPostCode:   getRegexRuleMatchPostCodeFunc,
	//ruleimpl.RegexPassword:   getRegexRuleMatchPasswordFunc,
	//ruleimpl.RegexPassword2:  getRegexRuleMatchPassword2Func,
	//ruleimpl.RegexPassword3:  getRegexRuleMatchPassword3Func,
	//ruleimpl.RegexPassport:   getRegexRuleMatchPassportFunc,
	//ruleimpl.RegexPhone:      getRegexRuleMatchPhoneFunc,
	//ruleimpl.RegexTelephone:  getRegexRuleMatchTelephoneFunc,
	//ruleimpl.RegexPhoneLoos:  getRegexRuleMatchPhoneLooseFunc,
	//ruleimpl.RegexEmail:      getRegexRuleMatchEmailFunc,

}

func isBuiltinRule(ruleName string) (registerRuleFunc, bool) {
	if rulefunc, ok := builtinRulesMapToFunc[ruleName]; ok {
		return rulefunc, true
	}
	return nil, false
}

var _ = registerFormatAndRegexRuleFunc()

// 注册
func registerFormatAndRegexRuleFunc() error {
	once.Do(func() {
		// regex
		for ruleName, impl := range ruleimpl.RegexRuleMapToFunc {
			builtinRulesMapToFunc[ruleName] = func(structRule *StructRule, fieldRule *FieldRules, s []string) ruleimpl.ValidFunc {
				return impl
			}
		}
	})
	return nil
}
