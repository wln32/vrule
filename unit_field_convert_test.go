package vrule

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func Test_convertNumber(t *testing.T) {

	//v1 := convertNumber[int8, int64]()(int8(18))
	//v2 := convertNumber[int16, int64]()(int16(18))
	//v3 := convertNumber[int32, int64]()(int32(18))
	//v4 := convertNumber[int, int64]()((18))
	//
	//v5 := convertNumber[uint8, uint64]()(uint8(18))
	//v6 := convertNumber[uint16, uint64]()(uint16(18))
	//v7 := convertNumber[uint32, uint64]()(uint32(18))
	//v8 := convertNumber[uint, uint64]()(uint(18))
	//
	//v9 := convertNumber[float32, float64]()(float32(18))

}

func Benchmark_my_convert(b *testing.B) {
	cvt := convertNumber[int, int64]()
	min := int64(20)
	max := int64(100)
	var a = reflect.ValueOf(88)

	fn := func(a any) error {

		val := a.(int64)
		if val < min || val > max {
			return fmt.Errorf("")
		}
		return nil
	}

	for i := 0; i < b.N; i++ {

		val := cvt(a.Interface())
		fn(val)
	}
}

func Benchmark_reflect_convert(b *testing.B) {

	min := int64(20)
	max := int64(100)

	fn := func(a reflect.Value) error {

		val := a.Int()
		if val < min || val > max {
			return fmt.Errorf("")
		}
		return nil
	}

	for i := 0; i < b.N; i++ {
		var a = reflect.ValueOf(88)
		fn(a)
	}
}

func Benchmark_convert_float64(b *testing.B) {

	cvt := convertNumber[int, float64]()

	min := float64(20)
	max := float64(100)
	fn := func(a any) error {

		val := a.(float64)
		if val < min || val > max {
			return fmt.Errorf("")
		}
		return nil
	}
	for i := 0; i < b.N; i++ {
		fn(cvt(98))
	}

}

type TestStructImplInterface struct{}

func (t *TestStructImplInterface) Test(ctx context.Context, in TestInput) error {
	ref := reflect.ValueOf(in.Value)
	if ref.IsValid() {
		return fmt.Errorf("1")
	}
	return nil
}

type TestInput struct {
	Value     any
	AssocData map[string]any
	Message   string
}

type TestInter interface {
	Test(ctx context.Context, in TestInput) error
}

type FuncInterface func(ctx context.Context, in TestInput) error

func (f FuncInterface) Test(ctx context.Context, in TestInput) error {
	return f(ctx, in)
}

func BenchmarkTest_Func_Interface(b *testing.B) {

	fn := func(ctx context.Context, in TestInput) error {
		ref := reflect.ValueOf(in.Value)
		if ref.IsValid() {
			return fmt.Errorf("1")
		}
		return nil
	}

	for i := 0; i < b.N; i++ {
		FuncInterface(fn).Test(context.Background(), TestInput{
			Value: 98,
		})
	}
}
func BenchmarkTest_Struct_Interface(b *testing.B) {

	var s TestStructImplInterface

	for i := 0; i < b.N; i++ {
		s.Test(context.Background(), TestInput{
			Value: 98,
		})
	}
}
