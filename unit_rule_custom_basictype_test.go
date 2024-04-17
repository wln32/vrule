package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

// 仅在测试时使用
type vInt int
type vInt8 int8
type vInt16 int16
type vInt32 int32
type vInt64 int64

type vUint uint
type vUint8 uint8
type vUint16 uint16
type vUint32 uint32
type vUint64 uint64

type vFloat32 float32
type vFloat64 float64

type vBool bool
type vString string

func Test_CustomType_Int_Lte(t *testing.T) {

	type CustomBasicType struct {
		A vInt
		B int `v:"lte:A"`
	}

	err := getTestValid().StructNotCache(CustomBasicType{
		B: 10,
	})

	t.Log(err)

}

func Test_CustomType_String_RequiredIf(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type RequiredIfBasicStruct struct {
			Name   string
			String vString `v:"required-if:Name,world"`
		}
		obj := &RequiredIfBasicStruct{
			Name: "world",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		t.Log(err)
		t.Assert(err, "The String field is required")
	})

	gtest.C(t, func(t *gtest.T) {
		type RequiredIfBasicStruct struct {
			Name   vString
			String string `v:"required-if:Name,world"`
		}
		obj := &RequiredIfBasicStruct{
			Name: "world",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)
		t.Log(err)
		t.Assert(err, "The String field is required")

	})

	gtest.C(t, func(t *gtest.T) {
		type vvString vString
		type RequiredIfBasicStruct struct {
			Name   vString
			String vvString `v:"required-if:Name,world"`
		}
		obj := &RequiredIfBasicStruct{
			Name: "world",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)
		t.Log(err)
		t.Assert(err, "The String field is required")

	})
}

func Test_CustomType_String_RequiredUnless(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		type RequiredIfBasicStruct struct {
			Name   string
			String vString `v:"required-unless:Name,world"`
		}
		obj := &RequiredIfBasicStruct{
			Name: "world1",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		t.Log(err)
		t.Assert(err, "The String field is required")
	})

	gtest.C(t, func(t *gtest.T) {
		type RequiredIfBasicStruct struct {
			Name   vString
			String string `v:"required-unless:Name,world"`
		}
		obj := &RequiredIfBasicStruct{
			Name: "world1",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)
		t.Log(err)
		t.Assert(err, "The String field is required")

	})

	gtest.C(t, func(t *gtest.T) {
		type vvString vString
		type RequiredIfBasicStruct struct {
			Name   vString
			String vvString `v:"required-unless:Name,world"`
		}
		obj := &RequiredIfBasicStruct{
			Name: "world1",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)
		t.Log(err)
		t.Assert(err, "The String field is required")

	})
}
