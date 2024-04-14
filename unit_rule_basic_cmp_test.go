package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Between_Basic(t *testing.T) {
	type _Between_BasicStruct struct {
		Int   int   `v:"between:1,20"`
		Int8  int8  `v:"between:1,20"`
		Int16 int16 `v:"between:1,20"`
		Int32 int32 `v:"between:1,20"`
		Int64 int64 `v:"between:1,20"`

		Uint    uint    `v:"between:1,20"`
		Uint8   uint8   `v:"between:1,20"`
		Uint16  uint16  `v:"between:1,20"`
		Uint32  uint32  `v:"between:1,20"`
		Uint64  uint64  `v:"between:1,20"`
		Float32 float32 `v:"between:1,20"`
		Float64 float64 `v:"between:1,20"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &_Between_BasicStruct{
			Int:     -2,
			Int8:    -3,
			Int16:   -4,
			Int32:   -5,
			Int64:   -6,
			Uint:    70,
			Uint8:   80,
			Uint16:  0,
			Uint32:  100,
			Uint64:  110,
			Float32: 120,
			Float64: 130,
		}
		wants := map[string]string{
			"Int":     "The Int value `-2` must be between 1 and 20",
			"Int8":    "The Int8 value `-3` must be between 1 and 20",
			"Int16":   "The Int16 value `-4` must be between 1 and 20",
			"Int32":   "The Int32 value `-5` must be between 1 and 20",
			"Int64":   "The Int64 value `-6` must be between 1 and 20",
			"Uint":    "The Uint value `70` must be between 1 and 20",
			"Uint8":   "The Uint8 value `80` must be between 1 and 20",
			"Uint16":  "The Uint16 value `0` must be between 1 and 20",
			"Uint32":  "The Uint32 value `100` must be between 1 and 20",
			"Uint64":  "The Uint64 value `110` must be between 1 and 20",
			"Float32": "The Float32 value `120` must be between 1 and 20",
			"Float64": "The Float64 value `130` must be between 1 and 20",
		}
		err := StructNotCache(obj).(*ValidationError)

		for i := 0; i < 100; i++ {
			for rule, msg := range wants {
				fieldError := err.GetFieldError(rule)
				t.Assert(fieldError.Error(), msg)
			}
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &_Between_BasicStruct{
			Int:     2,
			Int8:    3,
			Int16:   4,
			Int32:   5,
			Int64:   6,
			Uint:    7,
			Uint8:   8,
			Uint16:  9,
			Uint32:  10,
			Uint64:  11,
			Float32: 12,
			Float64: 13,
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_NotIn_Basic(t *testing.T) {
	type NotIn_BasicStruct struct {
		Int     int     `v:"not-in:1,20"`
		Int64   int64   `v:"not-in:1,20"`
		Uint    uint    `v:"not-in:1,20"`
		Uint8   uint8   `v:"not-in:1,20"`
		Float32 float32 `v:"not-in:1,20"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &NotIn_BasicStruct{
			Int:     1,
			Int64:   -6,
			Uint:    70,
			Uint8:   20,
			Float32: 120,
		}
		wants := map[string]string{
			"Int":   "The Int value `1` must not be in range: 1,20",
			"Uint8": "The Uint8 value `20` must not be in range: 1,20",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &NotIn_BasicStruct{}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_In_Basic(t *testing.T) {
	type In_BasicStruct struct {
		Int     int     `v:"in:1,20"`
		Int64   int64   `v:"in:1,20"`
		Uint    uint    `v:"in:1,20"`
		Uint8   uint8   `v:"in:1,20"`
		Float32 float32 `v:"in:1,20"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &In_BasicStruct{
			Int:     1,
			Int64:   -6,
			Uint:    70,
			Uint8:   20,
			Float32: 120,
		}
		wants := map[string]string{
			"Int64":   "The Int64 value `-6` is not in acceptable range: 1,20",
			"Uint":    "The Uint value `70` is not in acceptable range: 1,20",
			"Float32": "The Float32 value `120` is not in acceptable range: 1,20",
		}

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &In_BasicStruct{
			Int:     1,
			Int64:   1,
			Uint:    20,
			Uint8:   20,
			Float32: 20,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_Cmp_lte_Basic(t *testing.T) {
	type _Cmp_lteStruct struct {
		Score1 int8
		Score2 int   `v:"lte:Score1"`
		Score3 int8  `v:"lte:Score1"`
		Score4 int16 `v:"lte:Score1"`
		Score5 int32 `v:"lte:Score1"`
		Score6 int64 `v:"lte:Score1"`

		Score7  uint    `v:"lte:Score1"`
		Score8  uint8   `v:"lte:Score1"`
		Score9  uint16  `v:"lte:Score1"`
		Score10 uint32  `v:"lte:Score1"`
		Score11 uint64  `v:"lte:Score1"`
		Score12 float32 `v:"lte:Score1"`
		Score13 float64 `v:"lte:Score1"`
		Score14 int64   `v:"between:20,30"`
	}
	obj := &_Cmp_lteStruct{
		Score1:  -18,
		Score14: 90,
	}

	wants := map[string]string{
		"Score2":  "The Score2 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score3":  "The Score3 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score4":  "The Score4 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score5":  "The Score5 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score6":  "The Score6 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score7":  "The Score7 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score8":  "The Score8 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score9":  "The Score9 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score10": "The Score10 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score11": "The Score11 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score12": "The Score12 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score13": "The Score13 value `0` must be lesser than or equal to field Score1 value `-18`",
		"Score14": "The Score14 value `90` must be between 20 and 30",
	}

	gtest.C(t, func(t *gtest.T) {
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
}

func Test_Cmp_lt_Basic(t *testing.T) {
	type _Cmp_lt_1Struct struct {
		Score1 int8
		Score2 int   `v:"lt:Score1"`
		Score3 int8  `v:"lt:Score1"`
		Score4 int16 `v:"lt:Score1"`
		Score5 int32 `v:"lt:Score1"`
		Score6 int64 `v:"lt:Score1"`

		Score7  uint    `v:"lt:Score1"`
		Score8  uint8   `v:"lt:Score1"`
		Score9  uint16  `v:"lt:Score1"`
		Score10 uint32  `v:"lt:Score1"`
		Score11 uint64  `v:"lt:Score1"`
		Score12 float32 `v:"lt:Score1"`
		Score13 float64 `v:"lt:Score1"`
	}
	obj := &_Cmp_lt_1Struct{
		Score1: -18,
	}

	wants := map[string]string{
		"Score2":  "The Score2 value `0` must be lesser than field Score1 value `-18`",
		"Score3":  "The Score3 value `0` must be lesser than field Score1 value `-18`",
		"Score4":  "The Score4 value `0` must be lesser than field Score1 value `-18`",
		"Score5":  "The Score5 value `0` must be lesser than field Score1 value `-18`",
		"Score6":  "The Score6 value `0` must be lesser than field Score1 value `-18`",
		"Score7":  "The Score7 value `0` must be lesser than field Score1 value `-18`",
		"Score8":  "The Score8 value `0` must be lesser than field Score1 value `-18`",
		"Score9":  "The Score9 value `0` must be lesser than field Score1 value `-18`",
		"Score10": "The Score10 value `0` must be lesser than field Score1 value `-18`",
		"Score11": "The Score11 value `0` must be lesser than field Score1 value `-18`",
		"Score12": "The Score12 value `0` must be lesser than field Score1 value `-18`",
		"Score13": "The Score13 value `0` must be lesser than field Score1 value `-18`",
	}

	gtest.C(t, func(t *gtest.T) {

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
}

func Test_Cmp_gt_Basic(t *testing.T) {
	type Cmp_gtStruct struct {
		Score1  int8
		Score4  int16   `v:"gt:Score1"`
		Score5  int32   `v:"gt:Score1"`
		Score11 uint64  `v:"gt:Score1"`
		Score12 float32 `v:"gt:Score1"`
	}
	obj := &Cmp_gtStruct{
		Score1: 18,
	}

	wants := map[string]string{
		"Score4":  "The Score4 value `0` must be greater than field Score1 value `18`",
		"Score5":  "The Score5 value `0` must be greater than field Score1 value `18`",
		"Score11": "The Score11 value `0` must be greater than field Score1 value `18`",
		"Score12": "The Score12 value `0` must be greater than field Score1 value `18`",
	}

	gtest.C(t, func(t *gtest.T) {

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
}

func Test_Cmp_gte_Basic(t *testing.T) {
	type Cmp_gteStruct struct {
		Score1  int8
		Score4  int16   `v:"gte:Score1"`
		Score5  int32   `v:"gte:Score1"`
		Score11 uint64  `v:"gte:Score1"`
		Score12 float32 `v:"gte:Score1"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &Cmp_gteStruct{
			Score1: 18,
		}
		wants := map[string]string{
			"Score4":  "The Score4 value `0` must be greater than or equal to field Score1 value `18`",
			"Score5":  "The Score5 value `0` must be greater than or equal to field Score1 value `18`",
			"Score11": "The Score11 value `0` must be greater than or equal to field Score1 value `18`",
			"Score12": "The Score12 value `0` must be greater than or equal to field Score1 value `18`",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
}

func Test_EqString_Basic(t *testing.T) {
	type EqStringStruct struct {
		Name1 string
		Name2 string `v:"eq:Name1"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &EqStringStruct{
			Name2: "wln32",
		}
		wants := map[string]string{
			"Name2": "The Name2 value `wln32` must be equal to field Name1 value ``",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &EqStringStruct{
			Name2: "wln32",
			Name1: "wln32",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_EqNumber_Basic(t *testing.T) {
	type EqNumberStruct struct {
		Name1 int
		Name2 int `v:"eq:Name1"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &EqNumberStruct{
			Name2: 32,
		}
		wants := map[string]string{
			"Name2": "The Name2 value `32` must be equal to field Name1 value `0`",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &EqNumberStruct{
			Name2: 32,
			Name1: 32,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_EqBool_Basic(t *testing.T) {
	type EqBoolStruct struct {
		Name1 bool
		Name2 bool `v:"eq:Name1"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &EqBoolStruct{
			Name2: true,
		}
		wants := map[string]string{
			"Name2": "The Name2 value `true` must be equal to field Name1 value `false`",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &EqBoolStruct{
			Name2: true,
			Name1: true,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_NotEqString_Basic(t *testing.T) {
	type NotEqStringStruct struct {
		Name1 string
		Name2 string `v:"not-eq:Name1"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &NotEqStringStruct{
			Name2: "wln32",
			Name1: "wln32",
		}
		wants := map[string]string{
			"Name2": "The Name2 value `wln32` must not be equal to field Name1 value `wln32`",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := &NotEqStringStruct{
			Name2: "wln32",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_NotEqNumber_Basic(t *testing.T) {
	type NotEqNumberStruct struct {
		Name1 int
		Name2 int `v:"not-eq:Name1"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &NotEqNumberStruct{
			Name2: 32,
			Name1: 32,
		}
		wants := map[string]string{
			"Name2": "The Name2 value `32` must not be equal to field Name1 value `32`",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &NotEqNumberStruct{
			Name2: 32,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_NotEqBool_Basic(t *testing.T) {
	type NotEqBoolStruct struct {
		Name1 bool
		Name2 bool `v:"not-eq:Name1"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &NotEqBoolStruct{
			Name2: true,
			Name1: true,
		}
		wants := map[string]string{
			"Name2": "The Name2 value `true` must not be equal to field Name1 value `true`",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &NotEqBoolStruct{
			Name2: true,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Max_Basic(t *testing.T) {
	type MaxNumberStruct struct {
		Name1 int `v:"max:32"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &MaxNumberStruct{
			Name1: 64,
		}
		wants := map[string]string{
			"Name1": "The Name1 value `64` must be equal or lesser than 32",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &MaxNumberStruct{
			Name1: 20,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Min_Basic(t *testing.T) {
	type MinNumberStruct struct {
		Name1 int `v:"min:32"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &MinNumberStruct{
			Name1: 20,
		}
		wants := map[string]string{
			"Name1": "The Name1 value `20` must be equal or greater than 32",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &MinNumberStruct{
			Name1: 64,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
