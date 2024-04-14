package vrule

import (
	"strings"
)

func replaceRuleMsg_Field1(f *FieldRules, ruleName string, field1 string) {
	msg := f.MsgArray[ruleName]
	msg = strings.Replace(msg, "{field1}", field1, -1)
	f.MsgArray[ruleName] = msg
}
func replaceRuleMsg_Pattern(f *FieldRules, ruleName string, field1 string) {
	msg := f.MsgArray[ruleName]
	msg = strings.Replace(msg, "{pattern}", field1, -1)
	f.MsgArray[ruleName] = msg
}

func replaceRuleMsg_Max(f *FieldRules, ruleName string, field1 string) {
	msg := f.MsgArray[ruleName]
	msg = strings.Replace(msg, "{max}", field1, -1)
	f.MsgArray[ruleName] = msg
}

func replaceRuleMsg_Min(f *FieldRules, ruleName string, field1 string) {
	msg := f.MsgArray[ruleName]
	msg = strings.Replace(msg, "{min}", field1, -1)
	f.MsgArray[ruleName] = msg
}

func replaceRuleMsg_Size(f *FieldRules, ruleName string, field1 string) {
	msg := f.MsgArray[ruleName]
	msg = strings.Replace(msg, "{size}", field1, -1)
	f.MsgArray[ruleName] = msg
}

func replaceRuleMsg_CustomRule(f *FieldRules, ruleName string, field1 string) {
	msg := f.MsgArray[ruleName]
	msg = strings.Replace(msg, "{customrule}", field1, -1)
	f.MsgArray[ruleName] = msg
}
