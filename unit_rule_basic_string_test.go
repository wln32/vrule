package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_MaxLength_Basic(t *testing.T) {
	type MaxLengthStruct struct {
		Name1 string `v:"max-length:5"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &MaxLengthStruct{
			Name1: "fsjdjb",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `fsjdjb` length must be equal or lesser than 5",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &MaxLengthStruct{
			Name1: "hel",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_MinLength_Basic(t *testing.T) {
	type MinLengthStruct struct {
		Name1 string `v:"min-length:5"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &MinLengthStruct{
			Name1: "hel",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `hel` length must be equal or greater than 5",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &MinLengthStruct{
			Name1: "fsjdjb",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Length_Basic(t *testing.T) {
	type LengthStruct struct {
		Name1 string `v:"length:4,6"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &LengthStruct{
			Name1: "hel",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `hel` length must be between 4 and 6",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &LengthStruct{
			Name1: "fsjdjb",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Size_Basic(t *testing.T) {
	type SizeStruct struct {
		Name1 string `v:"size:4"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &SizeStruct{
			Name1: "hel",
		}
		wants := map[string]string{
			"Name1": "The Name1 value `hel` length must be 4",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &SizeStruct{
			Name1: "fsjd",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
