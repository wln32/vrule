package vrule

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gogf/gf/v2/util/gvalid"
)

type BenchUser1 struct {
	NickName string `v:"required"`
	Age1     int    `v:"required"`
}
type BenchUser2 struct {
	NickName string `v:"required"`
	Age2     int    `v:"required"`

	//NickName3 string
	//Age4      int
	//
	//NickName5 string
	//Age6      int
	//NickName7 string
	//Age8      int
}
type BenchUser struct {
	ID     *int                  `v:"required" `
	Name   string                `v:"required" `
	User1  []*BenchUser1         `v:"required" `
	Users2 map[string]BenchUser2 `v:"required" `
	//UsersPtr *BenchUser            //  `v:"required" `
}

func Benchmark_gf_Validate(b *testing.B) {

	id := 312
	var req = &BenchUser{

		ID:   &id,
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
	_ = req

	valid := gvalid.New().Bail()

	var ctx = context.TODO()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		valid.Data(req).Run(ctx)
	}
}

func Benchmark_reflect_GetByNameField_t(b *testing.B) {
	a := reflect.ValueOf(BenchUser2{})

	for i := 0; i < b.N; i++ {
		a.FieldByName("NickName5")
	}
}
func Benchmark_reflect_GetByIndexField_t(b *testing.B) {
	a := reflect.ValueOf(BenchUser2{})

	for i := 0; i < b.N; i++ {
		a.Field(5)
	}
}

func Benchmark_validator_Validate(b *testing.B) {
	validate := validator.New()
	type BenchUser1 struct {
		NickName string `validate:"required"`
		Age1     int    `validate:"required"`
	}
	type BenchUser2 struct {
		NickName string `validate:"required"`
		Age2     int    `validate:"required"`

		//NickName3 string
		//Age4      int
		//
		//NickName5 string
		//Age6      int
		//NickName7 string
		//Age8      int
	}
	type BenchUser struct {
		A      int                   `json:"-"`
		ID     *int                  `validate:"required" `
		Name   string                `validate:"required" `
		User1  []*BenchUser1         `validate:"dive" `
		Users2 map[string]BenchUser2 `validate:"dive" `
		//UsersPtr *BenchUser            //  `v:"required" `
	}

	ID := 1550
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
	for i := 0; i < b.N; i++ {
		validate.Struct(req2)
	}
}

func Benchmark_my_fieldByIndex_return_error_Validate(b *testing.B) {

	ID := 1550
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
	for i := 0; i < b.N; i++ {
		StructNotCache(req2)

	}
}

func Benchmark_my_fieldByIndex_args_error_Validate(b *testing.B) {

	ID := 1550
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
	for i := 0; i < b.N; i++ {
		StructNotCache(req2)
	}
}

func Benchmark_Struct_fields1_my_validate_func(b *testing.B) {

	type fields1Struct424 struct {
		Name string `v:"required"`
	}
	req := &fields1Struct424{
		Name: "",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StructNotCache(req)
	}
}

func Benchmark_Struct_fields4_slice_map_my_validate_func(b *testing.B) {

	type SliceFields2Struct1 struct {
		Name1 string `v:"required"`
		Name2 string `v:"required"`
	}
	type MapFields2Struct2 struct {
		Name1 string `v:"required"`
		Name2 string `v:"required"`
	}

	type fields4Struct441 struct {
		Name1 string                       `v:"required"`
		Name2 string                       `v:"required"`
		Name3 []SliceFields2Struct1        `v:"required"`
		Name4 map[string]MapFields2Struct2 `v:"required"`
	}
	req := &fields4Struct441{
		Name1: "54654",
		Name2: "gftryrtgfd",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StructNotCache(req)
	}
}
