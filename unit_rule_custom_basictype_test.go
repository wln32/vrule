package vrule

import "testing"

func Test_Custom_BasicType(t *testing.T) {
	type MyInt int

	type CustomBasicType struct {
		A MyInt
		B int `v:"lte:A"`
	}

	err := StructNotCache(CustomBasicType{
		B: 10,
	})

	t.Log(err)

}
