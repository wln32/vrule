package ruleimpl

var RuleMsgMap = map[string]string{
	Required:           `The {field} field is required`,
	RequiredIf:         "The {field} field is required",
	RequiredUnless:     "The {field} field is required",
	RequiredWith:       "The {field} field is required",
	RequiredWithAll:    "The {field} field is required",
	RequiredWithout:    "The {field} field is required",
	RequiredWithoutAll: "The {field} field is required",
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	Between:   "The {field} value `{value}` must be between {min} and {max}",
	Size:      "The {field} value `{value}` length must be {size}",
	Length:    "The {field} value `{value}` length must be between {min} and {max}",
	MaxLength: "The {field} value `{value}` length must be equal or lesser than {max}",
	MinLength: "The {field} value `{value}` length must be equal or greater than {min}",
	//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	Lte:       "The {field} value `{value}` must be lesser than or equal to field {field1} value `{value1}`",
	Lt:        "The {field} value `{value}` must be lesser than field {field1} value `{value1}`",
	Gte:       "The {field} value `{value}` must be greater than or equal to field {field1} value `{value1}`",
	Gt:        "The {field} value `{value}` must be greater than field {field1} value `{value1}`",
	NotEq:     "The {field} value `{value}` must not be equal to field {field1} value `{value1}`",
	Eq:        "The {field} value `{value}` must be equal to field {field1} value `{value1}`",
	Different: "The {field} value `{value}` must be different from field {field1} value `{value1}`",
	Same:      "The {field} value `{value}` must be the same as field {field1} value `{value1}`",
	Max:       "The {field} value `{value}` must be equal or lesser than {max}",
	Min:       "The {field} value `{value}` must be equal or greater than {min}",
	In:        "The {field} value `{value}` is not in acceptable range: {pattern}",
	NotIn:     "The {field} value `{value}` must not be in range: {pattern}",

	// ::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	RegexDomain: "The {field} value `{value}` is not a valid domain format",
	RegexEmail:  "The {field} value `{value}` is not a valid email address",
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	RegexUrl:  "The {field} value `{value}` is not a valid URL address",
	RegexIP:   "The {field} value `{value}` is not a valid IP address",
	RegexIPV4: "The {field} value `{value}` is not a valid IPv4 address",
	RegexIPV6: "The {field} value `{value}` is not a valid IPv6 address",
	RegexMAC:  "The {field} value `{value}` is not a valid MAC address",
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	RegexPassport:  "The {field} value `{value}` is not a valid passport format",
	RegexPassword:  "The {field} value `{value}` is not a valid password format",
	RegexPassword2: "The {field} value `{value}` is not a valid password format",
	RegexPassword3: "The {field} value `{value}` is not a valid password format",
	//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	RegexPhone:     "The {field} value `{value}` is not a valid phone number",
	RegexPhoneLoos: "The {field} value `{value}` is not a valid phone number",
	RegexPostCode:  "The {field} value `{value}` is not a valid postcode format",
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	RegexQQ:         "The {field} value `{value}` is not a valid QQ number",
	RegexResidentID: "The {field} value `{value}` is not a valid resident id number",
	RegexTelephone:  "The {field} value `{value}` is not a valid telephone number",
	RegexBankCard:   "The {field} value `{value}` is not a valid bank card number",

	// ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	JsonRuleName:    "The {field} value `{value}` is not a valid JSON string",
	IntegerRuleName: "The {field} value `{value}` is not an integer",
	FloatRuleName:   "The {field} value `{value}` is not of valid float type",
	BooleanRuleName: "The {field} value `{value}` field must be true or false",
	// ::::::::::::::::::::::::::::::::::::::::::::::::::
	RegexRuleName:    "The {field} value `{value}` must be in regex of: {pattern}",
	NotRegexRuleName: "The {field} value `{value}` should not be in regex of: {pattern}",

	//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	Date:        "The {field} value `{value}` is not a valid date",
	DataFormat:  "The {field} value `{value}` does not match the format: {pattern}",
	DateTime:    "The {field} value `{value}` is not a valid datetime",
	After:       "The {field} value `{value}` must be after field {field1} value `{value1}`",
	AfterEqual:  "The {field} value `{value}` must be after or equal to field {field1} value `{value1}`",
	Before:      "The {field} value `{value}` must be before field {field1} value `{value1}`",
	BeforeEqual: "The {field} value `{value}` must be before or equal to field {field1} value `{value1}`",
}
