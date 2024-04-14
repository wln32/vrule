package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Format_Boolean_Basic(t *testing.T) {
	type FormatBooleanStruct struct {
		Name1 string `v:"boolean"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatBooleanStruct{
			Name1: "hel",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `hel` field must be true or false",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatBooleanStruct{
			Name1: "false",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Format_Float_Basic(t *testing.T) {
	type FormatFloatStruct struct {
		Name1 string `v:"float"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatFloatStruct{
			Name1: "1gfdkg",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `1gfdkg` is not of valid float type",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatFloatStruct{
			Name1: "2.34",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Format_Integer_Basic(t *testing.T) {
	type FormatIntegerStruct struct {
		Name1 string `v:"integer"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatIntegerStruct{
			Name1: "1gfdkg",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `1gfdkg` is not an integer",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatIntegerStruct{
			Name1: "2",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Format_Json_Basic(t *testing.T) {
	type FormatJsonStruct struct {
		Name1 string `v:"json"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatJsonStruct{
			Name1: "1gfdkg",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `1gfdkg` is not a valid JSON string",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatJsonStruct{
			Name1: `{"longName":1}`,
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
