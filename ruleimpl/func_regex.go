package ruleimpl

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/text/gregex"
)

var (
	patternQQ     = regexp.MustCompile(REGEX_QQ)
	patternDomain = regexp.MustCompile(REGEX_DOMAIN)
	patternUrl    = regexp.MustCompile(REGEX_URL)
	patternEmail  = regexp.MustCompile(REGEX_EMAIL)
	patternMac    = regexp.MustCompile(REGEX_MAC)
	patternIp     = regexp.MustCompile(REGEX_IP)
	patternIpv4   = regexp.MustCompile(REGEX_IPV4)
	patternIpv6   = regexp.MustCompile(REGEX_IPV6)
	//:::::::::::::::::::::::::::::::::::::::::
	// patternBankcard   = regexp.MustCompile(RegexBankCard)
	patternResidentId = regexp.MustCompile(REGEX_RESIDENT_ID)
	patternPostcode   = regexp.MustCompile(REGEX_POST_CODE)
	patternPassword   = regexp.MustCompile(REGEX_PASSWORD)
	patternPassword2  = regexp.MustCompile(REGEX_PASSWORD2)
	patternPassword3  = regexp.MustCompile(REGEX_PASSWORD3)
	patternPassport   = regexp.MustCompile(REGEX_PASSPORT)
	patternTelephone  = regexp.MustCompile(REGEX_TELPHONE)
	patternPhoneLoose = regexp.MustCompile(REGEX_PHONE_LOOSE)
	patternPhone      = regexp.MustCompile(REGEX_PHONE)
)

var RegexRuleMapToFunc = map[string]ValidFuncImpl{
	RegexQQ:         RegexRuleMatchQQ,
	RegexDomain:     RegexRuleMatchDomain,
	RegexUrl:        RegexRuleMatchUrl,
	RegexMAC:        RegexRuleMatchMac,
	RegexIP:         RegexRuleMatchIp,
	RegexIPV4:       RegexRuleMatchIpv4,
	RegexIPV6:       RegexRuleMatchIpv6,
	RegexBankCard:   RegexRuleMatchBankCard,
	RegexResidentID: RegexRuleMatchResidentId,
	RegexPostCode:   RegexRuleMatchPostCode,
	RegexPassword:   RegexRuleMatchPassword,
	RegexPassword2:  RegexRuleMatchPassword2,
	RegexPassword3:  RegexRuleMatchPassword3,
	RegexPassport:   RegexRuleMatchPassport,
	RegexPhone:      RegexRuleMatchPhone,
	RegexTelephone:  RegexRuleMatchTelephone,
	RegexPhoneLoos:  RegexRuleMatchPhoneLoose,
	RegexEmail:      RegexRuleMatchEmail,
	//RegexRuleName:   RegexRuleMatch,
	//
	//NotRegexRuleName: RegexRuleNoMatch,
}

type RegexMatch struct {
	Pattern *regexp.Regexp
}

func (r *RegexMatch) RegexRuleMatch(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := r.Pattern.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func (r *RegexMatch) RegexRuleNoMatch(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := r.Pattern.MatchString(val)
	if !match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}

func RegexRuleMatchQQ(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternQQ.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}

func RegexRuleMatchDomain(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternDomain.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchUrl(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternUrl.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchMac(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternMac.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchIp(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()

	if ok := patternIpv4.MatchString(val) || patternIpv6.MatchString(val); !ok {
		return errors.New(strings.Replace(input.Message, "{value}", val, 1))
	}

	return nil
}
func RegexRuleMatchIpv4(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternIpv4.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchIpv6(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternIpv6.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchBankCard(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := checkLuHn(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchResidentId(ctx context.Context, input RuleFuncInput) error {
	id := input.Value.String()
	ok := checkResidentId(id)
	if ok {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", id, 1))
}
func RegexRuleMatchPostCode(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternPostcode.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchPassword(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternPassword.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}

// 格式: password2
// 说明：中等强度密码（在通用密码规则的基础上，要求密码必须包含大小写字母和数字）。
func RegexRuleMatchPassword2(ctx context.Context, input RuleFuncInput) error {
	value := input.Value.String()
	if gregex.IsMatchString(`^[\w\S]{6,18}$`, value) &&
		gregex.IsMatchString(`[a-z]+`, value) &&
		gregex.IsMatchString(`[A-Z]+`, value) &&
		gregex.IsMatchString(`\d+`, value) {
		return nil
	}

	return errors.New(strings.Replace(input.Message, "{value}", value, 1))
}

// 格式: password3
// 说明：强等强度密码（在通用密码规则的基础上，必须包含大小写字母、数字和特殊字符）。
func RegexRuleMatchPassword3(ctx context.Context, input RuleFuncInput) error {
	value := input.Value.String()
	if gregex.IsMatchString(`^[\w\S]{6,18}$`, value) &&
		gregex.IsMatchString(`[a-z]+`, value) &&
		gregex.IsMatchString(`[A-Z]+`, value) &&
		gregex.IsMatchString(`\d+`, value) &&
		gregex.IsMatchString(`[^a-zA-Z0-9]+`, value) {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", value, 1))
}
func RegexRuleMatchPassport(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternPassport.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}

func RegexRuleMatchPhone(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternPhone.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}

func RegexRuleMatchTelephone(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternTelephone.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchPhoneLoose(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternPhoneLoose.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}
func RegexRuleMatchEmail(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.String()
	match := patternEmail.MatchString(val)
	if match {
		return nil
	}
	return errors.New(strings.Replace(input.Message, "{value}", val, 1))
}

// checkResidentId checks whether given id a china resident id number.
//
// xxxxxx yyyy MM dd 375 0  十八位
// xxxxxx   yy MM dd  75 0  十五位
//
// 地区：     [1-9]\d{5}
// 年的前两位：(18|19|([23]\d))  1800-2399
// 年的后两位：\d{2}
// 月份：     ((0[1-9])|(10|11|12))
// 天数：     (([0-2][1-9])|10|20|30|31) 闰年不能禁止29+
//
// 三位顺序码：\d{3}
// 两位顺序码：\d{2}
// 校验码：   [0-9Xx]
//
// 十八位：^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$
// 十五位：^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$
//
// 总：
// (^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)
func checkResidentId(id string) bool {
	id = strings.ToUpper(strings.TrimSpace(id))
	if len(id) != 18 {
		return false
	}
	var (
		weightFactor = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		checkCode    = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
		last         = id[17]
		num          = 0
	)
	for i := 0; i < 17; i++ {
		tmp, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		num = num + tmp*weightFactor[i]
	}
	if checkCode[num%11] != last {
		return false
	}

	return patternResidentId.MatchString(`id	`)
}

func checkLuHn(value string) bool {
	var (
		sum     = 0
		nDigits = len(value)
		parity  = nDigits % 2
	)
	for i := 0; i < nDigits; i++ {
		var digit = int(value[i] - 48)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
