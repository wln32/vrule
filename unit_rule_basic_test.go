package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Required_Basic(t *testing.T) {
	type RequiredBasicStruct struct {
		Int8Ptr  *int8          `v:"required"`
		String   string         `v:"required"`
		Int      int            `v:"required"`
		Bool     bool           `v:"required"`
		IntArray []int          `v:"required"`
		IntMap   map[int]string `v:"required"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredBasicStruct{
			String: "hello world",
		}

		// validator := ParseStruct(obj, nil)
		wants := map[string]string{
			"Int8Ptr": "The Int8Ptr field is required",

			"IntArray": "The IntArray field is required",
			"IntMap":   "The IntMap field is required",
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
		obj := &RequiredBasicStruct{
			String:  "hello world",
			Int8Ptr: new(int8),
			IntArray: []int{
				1, 2, 3,
			},
			IntMap: map[int]string{
				1: "hello",
				2: "world",
			},
		}
		err := StructNotCache(obj).(*ValidationError).Errors()

		t.Assert(err, nil)

	})

}

func Test_RequiredIf_Basic(t *testing.T) {
	type RequiredIfBasicStruct struct {
		Int8    string `v:"required"`
		String2 string `v:"required-if:Int8,96"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredIfBasicStruct{
			Int8: "96",
		}

		wants := map[string]string{
			"String2": "The String2 field is required",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredIfBasicStruct{
			Int8:    "96",
			String2: "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredIfBasicStruct{
			Int8: "9",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})

}

func Test_RequiredUnless_Basic(t *testing.T) {
	type RequiredUnlessBasicStruct struct {
		Int8    string `v:"required"`
		String2 string `v:"required-unless:Int8,96"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredUnlessBasicStruct{
			Int8: "9",
		}

		wants := map[string]string{
			"String2": "The String2 field is required",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredUnlessBasicStruct{
			Int8:    "9",
			String2: "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredUnlessBasicStruct{
			Int8: "96",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})

}

func Test_RequiredWith_Basic(t *testing.T) {
	type RequiredWith_BasicStruct struct {
		Int8 int8   `v:"required"`
		With string `v:"required-with:Int8"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWith_BasicStruct{}

		wants := map[string]string{
			"With": "The With field is required",
		}
		_ = wants
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWith_BasicStruct{
			With: "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		type Test_Struct_Required_GTime2 struct {
			Uid       int64       `json:"uid"`
			Nickname  string      `json:"nickname" v:"required-with:Uid"`
			StartTime *gtime.Time `json:"start_time" v:"required-with:EndTime"`
			EndTime   *gtime.Time `json:"end_time" v:"required-with:StartTime"`
		}
		obj := Test_Struct_Required_GTime2{
			StartTime: nil,
			EndTime:   nil,
		}

		wants := map[string]string{
			"Nickname": `The Nickname field is required`,
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		type Test_Struct_Required_GTime1 struct {
			Uid       int64       `json:"uid"`
			Nickname  string      `json:"nickname" v:"required-with:Uid"`
			StartTime *gtime.Time `json:"start_time" v:"required-with:EndTime"`
			EndTime   *gtime.Time `json:"end_time" v:"required-with:StartTime"`
		}
		data := Test_Struct_Required_GTime1{
			StartTime: gtime.Now(),
			EndTime:   nil,
		}
		verr := g.Validator().Data(data).Run(ctx)
		t.AssertNE(verr, nil)

		wants := map[string]string{
			"Nickname": `The Nickname field is required`,
			"EndTime":  `The EndTime field is required`,
		}

		err := StructNotCache(data).(*ValidationError)
		for rule, msg := range wants {
			t.Assert(err.GetFieldError(rule), msg)
		}

	})

	gtest.C(t, func(t *gtest.T) {
		type UserApiSearch271 struct {
			Uid      int64  `json:"uid" v:"required"`
			Nickname string `json:"nickname" v:"required-with:Uid"`
		}
		data := UserApiSearch271{
			Uid: 1,
		}
		err := StructNotCache(data)

		t.Assert(err, `The Nickname field is required`)

	})

}

func Test_RequiredWithAll_Basic(t *testing.T) {
	type RequiredWithAll_BasicStruct struct {
		Int8   int8 `v:"required"`
		Bool   bool `v:"required"`
		IntPtr *int8
		With   string `v:"required-with-all:Int8,Bool,IntPtr"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWithAll_BasicStruct{
			IntPtr: new(int8),
		}

		wants := map[string]string{
			"With": "The With field is required",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWithAll_BasicStruct{
			IntPtr: new(int8),
			With:   "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_RequiredWithout_Basic(t *testing.T) {
	type RequiredWithout_BasicStruct struct {
		Int8    int8
		String  string
		Without string `v:"required-without:Int8,String"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWithout_BasicStruct{}

		wants := map[string]string{
			"Without": `The Without field is required`,
		}

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWithout_BasicStruct{
			String:  "1",
			Without: "with-out",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWithout_BasicStruct{
			String: "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_RequiredWithoutAll_Basic(t *testing.T) {
	type RequiredWithoutAll_BasicStruct struct {
		String     string
		WithoutAll string `v:"required-without-all:String"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWithoutAll_BasicStruct{}

		wants := map[string]string{
			"WithoutAll": `The WithoutAll field is required`,
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredWithoutAll_BasicStruct{
			String: "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

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
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatBooleanStruct{
			Name1: "false",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatFloatStruct{
			Name1: "2.34",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatIntegerStruct{
			Name1: "2",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &FormatJsonStruct{
			Name1: `{"longName":1}`,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

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
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &DateStruct{
			Name1: `2006/01/02`,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &DateTimeStruct{
			Name1: `2006-01-02 15:04:05`,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := &DateFormatStruct{
			Name1: `2021/11/01 23:00:00`,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

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

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

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
		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

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

		err := StructNotCache(obj).(*ValidationError).Errors()
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
		err := StructNotCache(obj).(*ValidationError)

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
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_Regex_Domain_Basic(t *testing.T) {
	type RegeDomainStruct struct {
		Domain  string `v:"domain"`
		Domain1 string `v:"domain"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegeDomainStruct{
			Domain:  "goframe#org",
			Domain1: "1a.2b",
		}
		wants := map[string]string{
			"Domain":  "The Domain value `goframe#org` is not a valid domain format",
			"Domain1": "The Domain1 value `1a.2b` is not a valid domain format",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegeDomainStruct{
			Domain:  "goframe.org",
			Domain1: "99designs.com",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_Regex_Url_Basic(t *testing.T) {
	type RegexUrlStruct struct {
		URL string `v:"url"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexUrlStruct{
			URL: "ws://goframe.org",
		}
		wants := map[string]string{
			"URL": "The URL value `ws://goframe.org` is not a valid URL address",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexUrlStruct{
			URL: "https://www.bilibili.com/",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Mac_Basic(t *testing.T) {
	type RegexMacStruct struct {
		MAC string `v:"mac"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexMacStruct{
			MAC: "Z0-CC-6A-D6-B1-1A",
		}
		wants := map[string]string{
			"MAC": "The MAC value `Z0-CC-6A-D6-B1-1A` is not a valid MAC address",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexMacStruct{
			MAC: "4C-CC-6A-D6-B1-1A",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Ip_Basic(t *testing.T) {
	type RegexIpStruct struct {
		Ip   string `v:"ip"`
		Ipv4 string `v:"ipv4"`
		Ipv6 string `v:"ipv6"`
		Ip2  string `v:"ip"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexIpStruct{
			Ip:   "127.0.0.1",
			Ipv6: "fe80::812b:1158:1f43:f0d1",
			Ipv4: "520.255.255.255", // error >= 10000
			Ip2:  "ze80::812b:1158:1f43:f0d1",
		}
		wants := map[string]string{
			"Ipv4": "The Ipv4 value `520.255.255.255` is not a valid IPv4 address",
			"Ip2":  "The Ip2 value `ze80::812b:1158:1f43:f0d1` is not a valid IP address",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexIpStruct{
			Ip:   "127.0.0.1",
			Ipv6: "fe80::812b:1158:1f43:f0d1",
			Ipv4: "193.255.255.255", // error >= 10000
			Ip2:  "2001:0da8:0207:0000:0000:0000:0000:8207",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_QQ_Basic(t *testing.T) {
	type RegexQQStruct struct {
		QQ string `v:"qq"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexQQStruct{
			QQ: "9999",
		}
		wants := map[string]string{
			"QQ": "The QQ value `9999` is not a valid QQ number",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexQQStruct{
			QQ: "123456789",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Password_Basic(t *testing.T) {
	type RegexPasswordStruct struct {
		Password  string `v:"password"`
		Password2 string `v:"password2"`
		Password3 string `v:"password3"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexPasswordStruct{
			Password:  "9999999",
			Password2: "159789523a",
			Password3: "9851af47953a",
		}
		wants := map[string]string{

			"Password2": "The Password2 value `159789523a` is not a valid password format",
			"Password3": "The Password3 value `9851af47953a` is not a valid password format",
		}

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexPasswordStruct{
			Password:  "qwerg_0213",
			Password2: "Aqwehbj13142",
			Password3: "gAtrgfhg1454#0",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_BankCard_Basic(t *testing.T) {
	type Regex_BankCard_Struct struct {
		BankCard1 string `v:"bank-card"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_BankCard_Struct{
			BankCard1: "6225760079930218",
		}
		wants := map[string]string{
			"BankCard1": "The BankCard1 value `6225760079930218` is not a valid bank card number",
		}

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_BankCard_Struct{
			BankCard1: "6259650871772098",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_ResidentId_Basic(t *testing.T) {
	type Regex_ResidentId_Struct struct {
		ResidentId string `v:"resident-id"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_ResidentId_Struct{
			ResidentId: "320107199506285482",
		}
		wants := map[string]string{
			"ResidentId": "The ResidentId value `320107199506285482` is not a valid resident id number",
		}

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

}

func Test_Regex_Postcode_Basic(t *testing.T) {
	type Regex_Postcode_Struct struct {
		Postcode string `v:"postcode"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Postcode_Struct{
			Postcode: "1000000",
		}
		wants := map[string]string{
			"Postcode": "The Postcode value `1000000` is not a valid postcode format",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Postcode_Struct{
			Postcode: "100000",
		}
		err := StructNotCache(obj).(*ValidationError)

		t.Assert(err, nil)
	})
}

func Test_Regex_Passport_Basic(t *testing.T) {
	type Regex_Passport_Struct struct {
		Passport string `v:"passport"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Passport_Struct{
			Passport: "123456",
		}
		wants := map[string]string{
			"Passport": "The Passport value `123456` is not a valid passport format",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Passport_Struct{
			Passport: "a100000",
		}
		err := StructNotCache(obj).(*ValidationError)
		t.Assert(err, nil)
	})
}

func Test_Regex_Telephone_Basic(t *testing.T) {
	type Regex_Telephone_Struct struct {
		Telephone string `v:"telephone"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Telephone_Struct{
			Telephone: "20-77542145",
		}
		wants := map[string]string{
			"Telephone": "The Telephone value `20-77542145` is not a valid telephone number",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Telephone_Struct{
			Telephone: "010-77542145",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
func Test_Regex_PhoneLoose_Basic(t *testing.T) {
	type Regex_PhoneLoose_Struct struct {
		Phone string `v:"phone-loose"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_PhoneLoose_Struct{
			Phone: "11578912345",
		}
		wants := map[string]string{
			"Phone": "The Phone value `11578912345` is not a valid phone number",
		}

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_PhoneLoose_Struct{
			Phone: "13578912345",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Email_Basic(t *testing.T) {
	type Regex_Email_Struct struct {
		Email string `v:"email"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Email_Struct{
			Email: "gf#goframe.org",
		}
		wants := map[string]string{
			"Email": "The Email value `gf#goframe.org` is not a valid email address",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Email_Struct{
			Email: "gf@goframe.org.cn",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
func Test_Regex_Phone_Basic(t *testing.T) {
	type Regex_Phone_Struct struct {
		Phone string `v:"phone"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Phone_Struct{
			Phone: "11578912345",
		}

		wants := map[string]string{
			"Phone": "The Phone value `11578912345` is not a valid phone number",
		}

		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Phone_Struct{
			Phone: "13578912345",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Regex_Basic(t *testing.T) {
	type Regex_Regex_Struct struct {
		Pattern  string `v:"regex:[1-9][0-9]{4,14}"`
		Pattern2 string `v:"regex:[1-9][0-9]{4,14}"`
		Pattern3 string `v:"regex:[1-9][0-9]{4,14}"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Regex_Struct{
			Pattern:  "1234",
			Pattern2: "01234",
			Pattern3: "10000",
		}

		wants := map[string]string{
			"Pattern":  "The Pattern value `1234` must be in regex of: [1-9][0-9]{4,14}",
			"Pattern2": "The Pattern2 value `01234` must be in regex of: [1-9][0-9]{4,14}",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Regex_Struct{
			Pattern:  "13578912345",
			Pattern2: "1101234",
			Pattern3: "10233000",
		}
		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_NotRegex_Basic(t *testing.T) {
	type Regex_NotRegex_Struct struct {
		Regex1 string `v:"regex:\\d{4}"`
		Regex2 string `v:"not-regex:\\d{4}"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_NotRegex_Struct{
			Regex1: "1234",
			Regex2: "1234",
		}

		wants := map[string]string{
			"Regex2": "The Regex2 value `1234` should not be in regex of: \\d{4}",
		}
		err := StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_NotRegex_Struct{
			Regex1: "1234",
			Regex2: "hghghghgh",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
