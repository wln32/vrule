package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Date_Basic(t *testing.T) {
	type DateStruct struct {
		Name1 string `v:"date"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &DateStruct{
			Name1: "2006*/*454",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `2006*/*454` is not a valid date",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &DateStruct{
			Name1: `2006/01/02`,
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_DateTime_Basic(t *testing.T) {
	type DateTimeStruct struct {
		Name1 string `v:"datetime"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &DateTimeStruct{
			Name1: "2006*/*454",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `2006*/*454` is not a valid datetime",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &DateTimeStruct{
			Name1: `2006-01-02 15:04:05`,
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_DateFormat_Basic(t *testing.T) {
	type DateFormatStruct struct {
		Name1 string `v:"date-format:Y/m/d H:i:s"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &DateFormatStruct{
			Name1: "2006*/*454",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `2006*/*454` does not match the format: Y/m/d H:i:s",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := &DateFormatStruct{
			Name1: `2021/11/01 23:00:00`,
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Before_Basic(t *testing.T) {
	type BeforeStruct struct {
		Time1 string `v:"before:Time2"`
		Time2 string
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &BeforeStruct{
			Time1: "2022-09-04",
			Time2: "2022-09-03",
		}
		wants := map[string]string{
			"Time1": "The Time1 value `2022-09-04 00:00:00 +0800 CST` must be before field Time2 value `2022-09-03 00:00:00 +0800 CST`",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &BeforeStruct{
			Time1: "2022-09-01",
			Time2: "2022-09-03",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}
func Test_BeforeEqual_Basic(t *testing.T) {
	type BeforeEqualStruct struct {
		Time1 string `v:"before-equal:Time2"`
		Time2 string
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &BeforeEqualStruct{
			Time1: "2022-09-04",
			Time2: "2022-09-03",
		}
		wants := map[string]string{
			"Time1": "The Time1 value `2022-09-04 00:00:00 +0800 CST` must be before or equal to field Time2 value `2022-09-03 00:00:00 +0800 CST`",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &BeforeEqualStruct{
			Time1: "2022-09-03",
			Time2: "2022-09-03",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_After_Basic(t *testing.T) {
	type AfterStruct struct {
		Time1 string `v:"after:Time2"`
		Time2 string
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &AfterStruct{
			Time1: "2022-09-01",
			Time2: "2022-09-03",
		}
		wants := map[string]string{
			"Time1": "The Time1 value `2022-09-01 00:00:00 +0800 CST` must be after field Time2 value `2022-09-03 00:00:00 +0800 CST`",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &AfterStruct{
			Time1: "2022-09-04",
			Time2: "2022-09-03",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}
func Test_AfterEqual_Basic(t *testing.T) {
	type AfterEqualStruct struct {
		Time1 string `v:"after-equal:Time2"`
		Time2 string
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &AfterEqualStruct{
			Time1: "2022-09-01",
			Time2: "2022-09-03",
		}
		wants := map[string]string{
			"Time1": "The Time1 value `2022-09-01 00:00:00 +0800 CST` must be after or equal to field Time2 value `2022-09-03 00:00:00 +0800 CST`",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &AfterEqualStruct{
			Time1: "2022-09-03",
			Time2: "2022-09-03",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}
