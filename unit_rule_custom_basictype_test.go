package vrule

import (
	"testing"
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

func Test_Custom_BasicType(t *testing.T) {

	type CustomBasicType struct {
		A vInt
		B int `v:"lte:A"`
	}

	err := getTestValid().StructNotCache(CustomBasicType{
		B: 10,
	})

	t.Log(err)

}
