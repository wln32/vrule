package vrule

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gvalid"
)

func Test_gf_Struct_Self_Reference(t *testing.T) {
	var ctx = context.TODO()
	type mainUser struct {
		UsersPtr *mainUser
		ID       *int   `v:"required" `
		Name     string `v:"required" `
	}
	ID := 1550
	var in1 = &mainUser{
		ID: &ID,
		//Name: "validSelf_Reference",
	}
	var arg = &mainUser{
		ID:       in1.ID,
		Name:     "validSelf_Reference",
		UsersPtr: in1,
	}
	gtest.C(t, func(t *gtest.T) {
		valid := gvalid.New().Bail()
		err := valid.Data(arg).Run(ctx)
		t.AssertNE(err, nil)
	})
}

func Test_my_Struct_Self_Reference(t *testing.T) {

	type mainUser struct {
		UsersPtr *mainUser
		ID       *int   `v:"required"`
		Name     string `v:"required" `
	}

	var in1 = &mainUser{
		Name: "in1",
	}
	var arg = &mainUser{
		Name:     "arg",
		UsersPtr: in1,
	}

	gtest.C(t, func(t *gtest.T) {
		err := StructNotCache(arg)

		t.AssertNE(err, nil)
	})
}

func Test_validator_Struct_Self_Reference(t *testing.T) {

	type mainUser struct {
		UsersPtr *mainUser
		ID       *int   `validate:"required"`
		Name     string `validate:"required" `
	}

	var in1 = &mainUser{
		Name: "in1",
	}
	var arg = &mainUser{
		Name:     "arg",
		UsersPtr: in1,
	}

	gtest.C(t, func(t *gtest.T) {
		err := StructNotCache(arg)
		_ = err
	})
}
