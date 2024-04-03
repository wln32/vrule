package ruleimpl

const (
	Required           = "required"
	RequiredIf         = "required-if"
	RequiredUnless     = "required-unless"
	RequiredWith       = "required-with"
	RequiredWithAll    = "required-with-all"
	RequiredWithout    = "required-without"
	RequiredWithoutAll = "required-without-all"
	Between            = "between"
	Size               = "size"
	Length             = "length"
	MaxLength          = "max-length"
	MinLength          = "min-length"
	Lte                = "lte"
	Lt                 = "lt"
	Gte                = "gte"
	Gt                 = "gt"
	NotEq              = "not-eq"
	Eq                 = "eq"
	Different          = "different"
	Same               = "same"
	Max                = "max"
	Min                = "min"
	In                 = "in"
	NotIn              = "not-in"

	//::::::::::::::::::::::::::::::::::::::::::::::::::
	Date       = "date"
	DateTime   = "datetime"
	DataFormat = "date-format"

	Before      = "before"
	BeforeEqual = "before-equal"
	After       = "after"
	AfterEqual  = "after-equal"

	//=============================================
	Array = "#array"
	Enums = "#enums"

	// ==============================================
	RegexDomain = `domain`
	RegexEmail  = `email`
	//=========================
	RegexUrl  = `url`
	RegexIP   = "ip"
	RegexIPV4 = "ipv4"
	RegexIPV6 = "ipv6"
	RegexMAC  = `mac`
	//=============================================
	RegexPassport  = `passport`
	RegexPassword  = `password`
	RegexPassword2 = "password2"
	RegexPassword3 = "password3"
	//==============================================
	RegexPhone     = `phone`
	RegexPhoneLoos = `phone-loose`
	RegexPostCode  = `postcode`
	//=============================================
	RegexQQ         = `qq`
	RegexResidentID = "resident-id"
	RegexTelephone  = `telephone`
	RegexBankCard   = "bank-card"

	// ===============================================
	JsonRuleName    = "json"
	IntegerRuleName = "integer"
	FloatRuleName   = "float"
	BooleanRuleName = "boolean"
	// ==================================================
	RegexRuleName    = "regex"
	NotRegexRuleName = "not-regex"
)
