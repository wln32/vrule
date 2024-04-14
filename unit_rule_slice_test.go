package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Slice_Rec(t *testing.T) {
	type Slice_RecUserItem struct {
		A1 string `v:"required#请输入a1字段"`
		A2 int
	}

	type Slice_RecUser struct {
		Name  string              `v:"required|length:4,8#|长度必须在4-8之间"`
		Items []Slice_RecUserItem `v:"required"`
	}

	gtest.C(t, func(t *gtest.T) {
		data := &Slice_RecUser{
			Name: "1236",
			Items: []Slice_RecUserItem{
				{},
			},
		}

		err := getTestValid().StructNotCache(data).(*ValidationError)
		wants := map[string]string{
			"Items-1": "请输入a1字段",
		}
		itemErr := err.GetSliceFieldError("Items").GetError(0).GetFieldError("A1")
		t.Assert(itemErr, wants["Items-1"])
	})

}
