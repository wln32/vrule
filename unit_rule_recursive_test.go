package vrule

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

var ctx = context.TODO()

func Test_CheckStruct_Recursive_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type CheckStruct_RecursivePass struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type CheckStruct_RecursiveUser struct {
			Id   int
			Name string `v:"required"`
			Pass CheckStruct_RecursivePass
		}
		obj := &CheckStruct_RecursiveUser{
			Name: "",
			Pass: CheckStruct_RecursivePass{
				Pass1: "1",
				Pass2: "2",
			},
		}

		wants := map[string]string{
			"Name":  "The Name field is required",
			"Pass1": "The Pass1 value `1` must be the same as field Pass2 value `2`",

			"Pass2": "The Pass2 value `2` must be the same as field Pass1 value `1`",
		}

		for i := 0; i < 100; i++ {
			err := StructNotCache(obj).(*ValidationError)

			t.Assert(err.GetFieldError("Name"), wants["Name"])
			t.Assert(err.GetStructFieldError("Pass").GetFieldError("Pass1"), wants["Pass1"])
			t.Assert(err.GetStructFieldError("Pass").GetFieldError("Pass2"), wants["Pass2"])
		}
	})
}

func Test_CheckStruct_Recursive_SliceStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Pass57 struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User61 struct {
			Id     int
			Name   string `v:"required"`
			Passes []Pass57
		}
		obj := &User61{
			Name: "",
			Passes: []Pass57{
				{
					Pass1: "1",
					Pass2: "2",
				},
				{
					Pass1: "3",
					Pass2: "4",
				},
			},
		}

		wants := map[string]string{
			"1-Pass1": "The Pass1 value `1` must be the same as field Pass2 value `2`",
			"1-Pass2": "The Pass2 value `2` must be the same as field Pass1 value `1`",

			"2-Pass1": "The Pass1 value `3` must be the same as field Pass2 value `4`",
			"2-Pass2": "The Pass2 value `4` must be the same as field Pass1 value `3`",
			"Name":    "The Name field is required",
		}

		for i := 0; i < 100; i++ {
			err := StructNotCache(obj).(*ValidationError)

			t.Assert(err.GetFieldError("Name"), wants["Name"])

			index1 := err.GetSliceFieldError("Passes").GetError(0)
			index2 := err.GetSliceFieldError("Passes").GetError(1)

			t.Assert(index1.GetFieldError("Pass1"), wants["1-Pass1"])
			t.Assert(index1.GetFieldError("Pass2"), wants["1-Pass2"])

			t.Assert(index2.GetFieldError("Pass1"), wants["2-Pass1"])
			t.Assert(index2.GetFieldError("Pass2"), wants["2-Pass2"])
		}

	})
}

func Test_CheckStruct_Recursive_MapStruct_Bail(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Pass102 struct {
			Pass1 string `v:"required|same:Pass2"`
			Pass2 string `v:"required|same:Pass1"`
		}
		type User106 struct {
			Id     int
			Name   string `v:"required"`
			Passes map[string]Pass102
		}
		obj := &User106{
			Name: "",
			Passes: map[string]Pass102{
				"1": {
					Pass1: "1",
					Pass2: "2",
				},
				"2": {
					Pass1: "3",
					Pass2: "4",
				},
			},
		}
		wants := map[string]string{
			"1-Pass1": "The Pass1 value `1` must be the same as field Pass2 value `2`",
			"1-Pass2": "The Pass2 value `2` must be the same as field Pass1 value `1`",

			"2-Pass1": "The Pass1 value `3` must be the same as field Pass2 value `4`",
			"2-Pass2": "The Pass2 value `4` must be the same as field Pass1 value `3`",
			"Name":    "The Name field is required",
		}

		for i := 0; i < 100; i++ {
			err := StructNotCache(obj).(*ValidationError)

			t.Assert(err.GetFieldError("Name"), wants["Name"])
			mapFieldError := err.GetMapFieldError("Passes")

			key1 := mapFieldError.GetError("1")
			key2 := mapFieldError.GetError("2")

			t.Assert(key1.GetFieldError("Pass1"), wants["1-Pass1"])
			t.Assert(key1.GetFieldError("Pass2"), wants["1-Pass2"])

			t.Assert(key2.GetFieldError("Pass1"), wants["2-Pass1"])
			t.Assert(key2.GetFieldError("Pass2"), wants["2-Pass2"])
		}
	})
}
