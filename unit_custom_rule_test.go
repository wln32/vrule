package vrule

import (
	"context"
	"fmt"
	"github.com/wln32/vrule/ruleimpl"
	"strconv"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_CustomRule_Basic(t *testing.T) {

	type CustomStruct struct {
		Name string `v:"trim-length:6"`
	}

	fn := func(ctx context.Context, in *ruleimpl.CustomRuleInput) error {
		val := in.Value.(string)
		trimVal := strings.TrimSpace(val)
		trimLength, err := strconv.Atoi(in.Args)
		if err != nil {
			return err
		}

		if len(trimVal) != trimLength {
			return fmt.Errorf("the length of the string after removing spaces must be %d characters", trimLength)
		}
		return nil
	}

	gtest.C(t, func(t *gtest.T) {
		err := RegisterCustomRuleFunc(RegisterCustomRuleOption{
			RuleName: "trim-length",
			Fn:       fn,
		})
		t.AssertNil(err)
		var customStruct = CustomStruct{
			Name: "myn",
		}
		err = StructNotCache(customStruct)
		t.Assert(err, "the length of the string after removing spaces must be 6 characters")
	})

}

func Test_CustomRule_DuplicateDefinition(t *testing.T) {

	type CustomStruct struct {
		Min int `v:"min1:"`
	}

	fn := func(ctx context.Context, in *ruleimpl.CustomRuleInput) error {
		val := in.Value.(int)
		minValue, err := strconv.Atoi(in.Args)
		if err != nil {
			return err
		}

		if val < minValue {
			return fmt.Errorf("field:%s Does not meet the verification rules", in.FieldName)
		}
		return nil
	}

	gtest.C(t, func(t *gtest.T) {
		err := RegisterCustomRuleFunc(RegisterCustomRuleOption{
			RuleName: "min1",
			Fn:       fn,
		})
		t.AssertNil(err)
		var customStruct = CustomStruct{
			Min: 20,
		}
		err = StructNotCache(customStruct)
		t.Assert(err, `strconv.Atoi: parsing "": invalid syntax`)
	})

}

func Test_CustomRule_AssocField(t *testing.T) {

	type CustomStruct struct {
		Name string `v:""`
	}

	fn := func(ctx context.Context, in *ruleimpl.CustomRuleInput) error {
		val := in.Value.(string)
		trimVal := strings.TrimSpace(val)
		trimLength, err := strconv.Atoi(in.Args)
		if err != nil {
			return err
		}

		trimLen := len(trimVal)

		if trimLen != trimLength {
			return fmt.Errorf("the minimum length of the string after removing spaces should be %d characters", trimLength)
		}
		return nil
	}

	gtest.C(t, func(t *gtest.T) {
		err := RegisterCustomRuleFunc(RegisterCustomRuleOption{
			RuleName: "trim-min",
			Fn:       fn,
		})
		t.Assert(err, nil)
		var customStruct = CustomStruct{
			Name: "myn",
		}

		err = StructNotCache(customStruct)

		t.Assert(err, "current structure: vrule.CustomStruct has no rules to verify")
	})

}
