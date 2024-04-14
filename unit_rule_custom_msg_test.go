package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Struct_Required_CustomMsg(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type Struct_Required32 struct {
			Name string  `v:"Name@required|length:6,16#名称不能为空|名称长度为{min}到{max}个字符"`
			Age  float32 `v:"Age@between:18,30#年龄为18到30周岁"`
		}

		obj := &Struct_Required32{"john", 16}

		var wants = map[string]string{
			"Name": "名称长度为6到16个字符",
			"Age":  "年龄为18到30周岁",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

}

func Test_Struct_MaxMin_CustomMsg(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type MaxMin struct {
			Max int     `v:"max:5#最大值不能超过5"`
			Min float32 `v:"min:5#最小值不能小于-5"`
		}

		obj := &MaxMin{
			Max: 11,
			Min: 1,
		}

		var wants = map[string]string{
			"Max": "最大值不能超过5",
			"Min": "最小值不能小于-5",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

}

func Test_Struct_Required_CustomMsg_2(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type Struct_Required52 struct {
			Username string `json:"username" valid:"required#用户名不能为空"`
			Password string `json:"password" valid:"required#登录密码不能为空"`
			Id       int    `valid:"required|min:10#|ID不能为空"`
			Age      int    `valid:"required|min:1#|年龄不能为空"`
		}
		var obj Struct_Required52

		var wants = map[string]string{
			"Username": "用户名不能为空",
			"Password": "登录密码不能为空",
			"Id":       "ID不能为空",
			"Age":      "年龄不能为空",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})
}

func Test_Struct_With_EmbeddedObject(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type Embedded105 struct {
			Pass1 string `valid:"password1@required|same:Pass2"`
			Pass2 string `valid:"password2@required|same:Pass1"`
		}
		type EmbeddedObject109 struct {
			Id   int
			Name string `valid:"LongName@required#请输入您的姓名"`
			Embedded105
		}
		obj := &EmbeddedObject109{
			Name: "",
			Embedded105: Embedded105{
				Pass1: "1",
				Pass2: "2",
			},
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)
		wants := map[string]string{
			"Name":      "请输入您的姓名",
			"password1": "The password1 value `1` must be the same as field Pass2 value `2`",
			"password2": "The password2 value `2` must be the same as field Pass1 value `1`",
		}

		nameFieldError := err.GetFieldError("Name")
		t.Assert(nameFieldError.Error(), wants["Name"])

		pass1FieldError := err.GetStructFieldError("Embedded105").GetFieldError("Pass1")
		t.Assert(pass1FieldError.Error(), wants["password1"])

		pass2FieldError := err.GetStructFieldError("Embedded105").GetFieldError("Pass2")
		t.Assert(pass2FieldError.Error(), wants["password2"])

	})
}

func Test_Struct_Optional(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type CheckStruct_Optional149 struct {
			Page int `v:"required|min:1         # page is required"`
			Size int `v:"required|between:1,100 # size is required"`
		}
		obj := &CheckStruct_Optional149{
			Page: 1,
			Size: 10,
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		type CheckStruct_Optional164 struct {
			Page      int `v:"required|min:1         # page is required"`
			Size      int `v:"required|between:1,100 # size is required"`
			ProjectId int `v:"between:1,10000        # project id must between {min}, {max}"`
		}
		obj := &CheckStruct_Optional164{
			Page: 1,
			Size: 10,
		}
		wants := map[string]string{
			"ProjectId": "project id must between 1, 10000",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
}

func Test_Struct_NoValidTag_Ptr(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type NoValidTag_PtrStruct372 struct {
			Age uint `v:"min:18"`
		}
		type NoValidTagStruct375 struct {
			Name   string
			Params *NoValidTag_PtrStruct372
		}
		obj := &NoValidTagStruct375{
			Name: "john",
			Params: &NoValidTag_PtrStruct372{
				Age: 0,
			},
		}
		wants := map[string]string{
			"Age": "The Age value `0` must be equal or greater than 18",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		fieldError := err.GetStructFieldError("Params").GetFieldError("Age")
		t.Assert(fieldError, wants["Age"])
	})
}
