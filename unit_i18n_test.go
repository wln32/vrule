package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_i18n(t *testing.T) {
	type TestI18nStruct struct {
		Args  string `v:"required" `
		Args2 int64  `v:"between:2,15" `
		Args3 int    `v:"max:60" `
	}

	gtest.C(t, func(t *gtest.T) {
		valid := New()
		valid.I18nSetLanguage("zh-CN")

		obj := &TestI18nStruct{
			Args:  "",
			Args2: 60,
			Args3: 61,
		}

		err := valid.StructNotCache(obj).(*ValidationError)
		t.Assert(err.GetFieldError("Args"), "Args字段不能为空")
		t.Assert(err.GetFieldError("Args2"), "Args2字段值`60`必须介于 2和15之间")
		t.Assert(err.GetFieldError("Args3"), "Args3字段值`61`字段最大值应当为60")

	})

}
