package vrule

import (
	ruleimpl "github.com/wln32/vrule/ruleimpl"

	"github.com/gogf/gf/v2/util/gconv"
)

// 格式: size:size
// 说明：参数长度为 size (长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func getStringSizeRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Size(f, ruleimpl.Size, ruleVals[0])
	return &ruleimpl.StringSizeRule{
		Size: gconv.Int(ruleVals[0]),
	}
}

// 格式: length:min,max
// 说明：参数长度为min到max(长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func getStringLengthRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Min(f, ruleimpl.Length, ruleVals[0])
	replaceRuleMsg_Max(f, ruleimpl.Length, ruleVals[1])
	min := ruleVals[0]
	max := ruleVals[1]
	return &ruleimpl.StringLengthRule{
		Min: gconv.Int(min),
		Max: gconv.Int(max),
	}
}

// 格式: min-length:min
// 说明：参数长度最小为min(长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func getStringMinLengthRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Min(f, ruleimpl.MinLength, ruleVals[0])
	min := ruleVals[0]
	return &ruleimpl.StringMinLengthRule{
		Min: gconv.Int(min),
	}

}

// 格式: max-length:max
// 说明：参数长度最大为max(长度参数为整形)，注意底层使用Unicode计算长度，因此中文一个汉字占1个长度单位。
func getStringMaxLengthRuleFunc(_ *StructRule, f *FieldRules, ruleVals []string) ruleimpl.ValidFunc {
	replaceRuleMsg_Max(f, ruleimpl.MaxLength, ruleVals[0])
	max := ruleVals[0]
	return &ruleimpl.StringMaxLengthRule{
		Max: gconv.Int(max),
	}

}
