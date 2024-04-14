package vrule

import (
	"reflect"
	"testing"
)

func Test_getStructFields(t *testing.T) {
	type Embedded105Test_getStructFields struct {
		Id6   int
		Pass3 string `valid:"password1@required|same:password2#请输入您的密码|您两次输入的密码不一致"`
		Pass4 string `valid:"password2@required|same:password1#请再次输入您的密码|您两次输入的密码不一致"`
	}

	type Embedded105 struct {
		Id8   int
		Pass1 string `valid:"password1@required|same:password2#请输入您的密码|您两次输入的密码不一致"`
		Pass2 string `valid:"password2@required|same:password1#请再次输入您的密码|您两次输入的密码不一致"`
		E     Embedded105Test_getStructFields
	}
	type EmbeddedObject109 struct {
		Id   int
		Name string `valid:"LongName@required#请输入您的姓名"`
		E    *Embedded105
		Arr  []int
		Arr2 []Embedded105Test_getStructFields
	}
	user := &EmbeddedObject109{
		Name: "",
		E: &Embedded105{
			Pass1: "1",
			Pass2: "2",
		},
	}
	_ = user
	fields := getStructFields(reflect.TypeOf(EmbeddedObject109{}))

	var printStructfn func(typ reflect.Type)
	var printFieldFn func(i int, field reflect.StructField)
	printStructfn = func(typ reflect.Type) {
		// t.Log("struct==", Type)
		for i := 0; i < typ.NumField(); i++ {
			printFieldFn(i, typ.Field(i))
		}

	}

	printFieldFn = func(i int, field reflect.StructField) {
		// t.Log(i, field)
		switch field.Type.Kind() {
		case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Array:
			sty, ok := isStructType(field.Type)
			if ok {
				printStructfn(sty)
			}
		case reflect.Struct:
			printStructfn(field.Type)

		}
	}

	for i, field := range fields {
		printFieldFn(i, field.Field)
	}

}

func Test_ParseErrorMsg(t *testing.T) {
	msg1 := "The {field} value `{value}` must be after field {field1} value `{value1}`"
	handler := parseErrorMsg(msg1)
	t.Logf("%T\n", handler)

	msg2 := "The {field} value `{value}` must be after field {field1} value "
	handler = parseErrorMsg(msg2)
	t.Logf("%T\n", handler)

	msg3 := "The {field} value must be after field {field1} value `{value1}`"
	handler = parseErrorMsg(msg3)
	t.Logf("%T\n", handler)

	msg4 := "The {field} value  must be after field {field1} value "
	handler = parseErrorMsg(msg4)
	t.Logf("%T\n", handler)
}
