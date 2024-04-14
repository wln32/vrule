package ruleimpl

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type TimeRule struct {
	FieldName string
	// 关联字段索引
	AssocFieldIndex int
}

// 格式: before:field
// 说明：判断给定的日期/时间是否在指定字段的日期/时间之前。
// 版本：框架版本>=v2.2.0
func (t *TimeRule) Before(ctx context.Context, input RuleFuncInput) error {
	valueTime := gconv.Time(input.Value)
	beforeTime := gconv.Time(input.StructPtr.Field(t.AssocFieldIndex).Interface())
	if valueTime.Before(beforeTime) {
		return nil

	}
	return errors.New(gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  valueTime.String(),
		"{field1}": t.FieldName,
		"{value1}": beforeTime.String(),
	}))
}

// 格式: before-equal:field
// 说明：判断给定的日期/时间是否在指定字段的日期/时间之前，或者与指定字段的日期/时间相等。
// 版本：框架版本>=v2.2.0
func (t *TimeRule) BeforeEqual(ctx context.Context, input RuleFuncInput) error {
	valueTime := gconv.Time(input.Value)
	beforeTime := gconv.Time(input.StructPtr.Field(t.AssocFieldIndex).Interface())
	if valueTime.Before(beforeTime) || valueTime.Equal(beforeTime) {
		return nil

	}
	return errors.New(gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  valueTime.String(),
		"{field1}": t.FieldName,
		"{value1}": beforeTime.String(),
	}))

}

// 格式: after:field
// 说明：判断给定的日期/时间是否在指定字段的日期/时间之后。
// 版本：框架版本>=v2.2.0
func (t *TimeRule) After(ctx context.Context, input RuleFuncInput) error {
	valueTime := gconv.Time(input.Value)
	afterTime := gconv.Time(input.StructPtr.Field(t.AssocFieldIndex).Interface())
	if valueTime.After(afterTime) {
		return nil
	}

	return errors.New(gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  valueTime.String(),
		"{field1}": t.FieldName,
		"{value1}": afterTime.String(),
	}))
}

// 格式: after-equal:field
// 说明：判断给定的日期/时间是否在指定字段的日期/时间之后，或者与指定字段的日期/时间相等。
// 版本：框架版本>=v2.2.0
func (t *TimeRule) AfterEqual(ctx context.Context, input RuleFuncInput) error {
	valueTime := gconv.Time(input.Value)
	afterTime := gconv.Time(input.StructPtr.Field(t.AssocFieldIndex).Interface())
	if valueTime.After(afterTime) || valueTime.Equal(afterTime) {
		return nil
	}

	return errors.New(gstr.ReplaceByMap(input.Message, map[string]string{
		"{value}":  valueTime.String(),
		"{value1}": afterTime.String(),
	}))
}

type DateFormatRule struct {
	Format string
}

type iTime interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// 格式: date-format:format
// 说明：判断日期是否为指定的日期/时间格式，format参数格式为gtime日期格式(可以包含日期及时间)，格式说明参考章节：gtime模块
// 示例：date-format:Y-m-d H:i:s
func (t *DateFormatRule) Run(ctx context.Context, in RuleFuncInput) error {

	if _, err := gtime.StrToTimeFormat(in.Value.String(), t.Format); err != nil {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{value}": gconv.String(in.Value),
		}))
	}
	return nil

}

// 支持不带连接符号的8位长度日期，格式如： 2006-01-02, 2006/01/02, 2006.01.02, 20060102
var dateFormatPattern = regexp.MustCompile(`\d{4}[.\-_/]?\d{2}[.\-_/]?\d{2}`)

// 格式: date
// 说明：参数为常用日期类型，日期之间支持的连接符号-或/或.，
// 也支持不带连接符号的8位长度日期，格式如： 2006-01-02, 2006/01/02, 2006.01.02, 20060102
func DateRuleFunc(ctx context.Context, in RuleFuncInput) error {

	if dateFormatPattern.MatchString(in.Value.String()) == false {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{value}": gconv.String(in.Value),
		}))
	}
	return nil
}

// 格式: datetime
// 说明：参数为常用日期时间类型，其中日期之间支持的连接符号只支持-，格式如： 2006-01-02 12:00:00
func DateTimeRuleFunc(ctx context.Context, in RuleFuncInput) error {

	if _, err := gtime.StrToTimeFormat(in.Value.String(), `Y-m-d H:i:s`); err != nil {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{value}": gconv.String(in.Value),
		}))
	}
	return nil

}
