package vrule

import (
	"context"

	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/grand"
)

var testValid *Validator

func getTestValid() *Validator {
	if testValid == nil {
		testValid = New()
		testValid.StopOnFirstError(false)
	}
	return testValid
}

func Test_Struct_MultipleLevels(t *testing.T) {

	type TestErrorStruct struct {
		Args  string `v:"required" `
		Args2 string `v:"required" `
		Args3 string `v:"required" `
	}

	type mainUser struct {
		SliceStruct []TestErrorStruct `v:"required" `
		ID          string            `v:"required" `
		//Name   string            `v:"required" `
		//Score  string            `v:"required" `
		Scores []int `v:"required" `

		MapStruct map[string]TestErrorStruct `v:"required" `
		Struct    TestErrorStruct
	}

	var arg = &mainUser{
		ID: "",
		//Name:  "test",
		//Score: "score",
		SliceStruct: []TestErrorStruct{
			{Args: "1-args"},
		},
		MapStruct: map[string]TestErrorStruct{
			"key": {Args: "2-args"},
		},
		Struct: TestErrorStruct{
			Args3: "3-args",
		},
	}
	gtest.C(t, func(t *gtest.T) {

		for i := 0; i < 1000; i++ {
			err := getTestValid().StructNotCache(arg).(*ValidationError)
			idError := err.GetFieldError("ID")
			t.Assert(idError, "The ID field is required")
			scoresError := err.GetFieldError("Scores")
			t.Assert(scoresError, "The Scores field is required")

			index1 := err.GetSliceFieldError("SliceStruct").GetError(0)
			t.Assert(index1.GetFieldError("Args2"), "The Args2 field is required")
			t.Assert(index1.GetFieldError("Args3"), "The Args3 field is required")

			mapKey := err.GetMapFieldError("MapStruct").GetError("key")
			t.Assert(mapKey.GetFieldError("Args2"), "The Args2 field is required")
			t.Assert(mapKey.GetFieldError("Args3"), "The Args3 field is required")

			structError := err.GetStructFieldError("Struct")
			t.Assert(structError.GetFieldError("Args"), "The Args field is required")
			t.Assert(structError.GetFieldError("Args2"), "The Args2 field is required")
		}
	})
}
func Test_Struct_MultipleLevels_Bail(t *testing.T) {

	type TestErrorStruct struct {
		Args  string `v:"required" `
		Args2 string `v:"required" `
		Args3 string `v:"required" `
	}

	type mainUser struct {
		Struct      TestErrorStruct
		MapStruct   map[string]TestErrorStruct `v:"required" `
		SliceStruct []TestErrorStruct          `v:"required" `
		Scores      []int                      `v:"required" `
		ID          string                     `v:"required" `
	}

	var arg = &mainUser{
		ID: "",
		//Name:  "test",
		//Score: "score",
		SliceStruct: []TestErrorStruct{
			{Args: "1-args"},
		},
		MapStruct: map[string]TestErrorStruct{
			"key": {Args: "2-args"},
		},
		Struct: TestErrorStruct{
			Args3: "3-args",
		},
	}
	gtest.C(t, func(t *gtest.T) {
		valid := New()
		valid.StopOnFirstError(true)
		for i := 0; i < 1000; i++ {
			err := valid.StructNotCache(arg)

			t.Assert(err, `The Args field is required`)
		}
	})
}

func Benchmark_Struct_Test(b *testing.B) {
	type BenchUser1 struct {
		NickName string `v:"required"`
		Age1     int    `v:"required"`
	}
	type BenchUser2 struct {
		NickName string `v:"required"`
		Age2     int    `v:"required"`
	}
	type BenchUser struct {
		ID     *int                  `v:"required" `
		Name   string                `v:"required" `
		User1  []*BenchUser1         `v:"required" `
		Users2 map[string]BenchUser2 `v:"required" `
		//UsersPtr *BenchUser            //  `v:"required" `
	}
	ID := 1510
	var req2 = &BenchUser{
		ID:   &ID,
		Name: "wln",
		User1: []*BenchUser1{&BenchUser1{
			NickName: "gjfnj",
		}},
		Users2: map[string]BenchUser2{
			"1": BenchUser2{
				NickName: "kfgkfd",
			},
		},
	}

	b.ResetTimer()

	b.Run("直接验证", func(b *testing.B) {
		valid := New()
		rule := valid.ParseStruct(BenchUser{}, nil)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = rule.Valid(context.TODO(), reflect.ValueOf(req2), ValidRuleOption{})
		}
	})

	b.Run("通用化接口", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			getTestValid().StructNotCache(req2)
		}
	})
}

func Test_Concurrent_Register(t *testing.T) {
	type TestErrorStruct struct {
		Args  string `v:"required" `
		Args2 string `v:"required" `
		Args3 string `v:"required" `
	}
	var wg sync.WaitGroup
	reg := func() {
		defer wg.Add(-1)
		time.Sleep(time.Duration(grand.N(100, 1000)) * time.Millisecond)
		New().ParseStruct(&TestErrorStruct{}, nil)
	}
	const N = 10
	wg.Add(N)
	for i := 0; i < N; i++ {
		go reg()
	}
	wg.Wait()
}

func Benchmark_reflect_MapIter_map(b *testing.B) {
	m := map[string]string{
		"1": "1",
		"2": "2",
		"3": "4",
	}

	b.Run("map-iter", func(b *testing.B) {
		mapVal := reflect.ValueOf(m)
		for i := 0; i < b.N; i++ {
			iter := mapVal.MapRange()
			for iter.Next() {
				_ = iter.Value()
				_ = iter.Key()
			}
		}
	})

	b.Run("map-keys", func(b *testing.B) {
		mapVal := reflect.ValueOf(m)

		for i := 0; i < b.N; i++ {
			keys := mapVal.MapKeys()

			for _, key := range keys {
				mapVal.MapIndex(key)
			}
		}
	})

}

func Benchmark_AssocField_StructPtr(b *testing.B) {
	type AssocField_StructPtr struct {
		Num1 int
		Num2 int `v:"eq:Num1"`
	}

	obj := &AssocField_StructPtr{
		Num1: 1,
		Num2: 2,
	}

	valid := New().ParseStruct(obj, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		valid.Valid(context.TODO(), reflect.ValueOf(obj), ValidRuleOption{})
	}

}
