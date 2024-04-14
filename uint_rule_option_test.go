package vrule

import (
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Option_FieldName(t *testing.T) {
	type OptionFieldName struct {
		Name string `json:"Name" v:"required"`
	}

	gtest.C(t, func(t *gtest.T) {
		valid := New()
		valid.SetFieldNameFunc(func(_ reflect.Type, field reflect.StructField) string {

			name := field.Tag.Get("json")
			if name == "" {
				name = field.Name
			}
			return name
		})

		obj := &OptionFieldName{}
		err := valid.StructNotCache(obj)
		t.Assert(err, `The Name field is required`)
	})

}

func Test_Option_FilterField(t *testing.T) {
	type OptionFilterField struct {
		Id     uint   `json:"id"      v:"min:1"`
		PicUrl string `json:"pic_url" v:"required"`
		Link   string `json:"link"    v:"required"`
		Sort   int    `json:"sort"    v:"min:10" d:"1"`
	}

	gtest.C(t, func(t *gtest.T) {
		valid := New()
		valid.SetFilterFieldFunc(func(structType reflect.Type, field reflect.StructField) bool {
			tag := field.Tag.Get("d")
			if tag != "" {
				return true
			}
			return false
		})

		obj := &OptionFilterField{
			Id:     1,
			PicUrl: "https://example.com",
			Link:   "https://example.com",
			Sort:   0,
		}
		err := valid.StructNotCache(obj)
		t.Assert(err, nil)
	})
}

func Test_Option_StopOnFirstError(t *testing.T) {
	type OptionStopOnFirstError struct {
		Id     uint   `json:"id"      v:"min:1" `
		PicUrl string `json:"pic_url" v:"required" `
		Link   string `json:"link"    v:"required"`
		Sort   int    `json:"sort"    v:"min:10" d:"1"`
	}

	gtest.C(t, func(t *gtest.T) {
		valid := New()
		valid.StopOnFirstError(false)

		obj := &OptionStopOnFirstError{}
		err := valid.StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(len(err), 4)
	})
	gtest.C(t, func(t *gtest.T) {
		valid := New()
		valid.StopOnFirstError(true)

		obj := &OptionStopOnFirstError{}
		err := valid.StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(len(err), 1)
	})
}
