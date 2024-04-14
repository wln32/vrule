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
		Int8    int8
		String2 string `v:"required-if:Int8,96"`

		Int16   vInt16
		String3 vString `v:"required-if:Int16,98"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredIfBasicStruct{
			Int8:  96,
			Int16: 97,
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
			Int8:    96,
			String2: "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredIfBasicStruct{
			Int8: 9,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})

}

func Test_RequiredUnless_Basic(t *testing.T) {
	type RequiredUnlessBasicStruct struct {
		Int8    int8
		String2 string `v:"required-unless:Int8,96"`

		Int16   vInt16
		String3 string `v:"required-unless:Int16,0"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredUnlessBasicStruct{
			Int8: 9,
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
			Int8:    9,
			String2: "1",
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RequiredUnlessBasicStruct{
			Int8: 96,
		}

		err := StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)

	})

}

func Test_RequiredWith_Basic(t *testing.T) {
	type RequiredWith_BasicStruct struct {
		Int8 int8
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
		Int8   int8
		Bool   bool
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
