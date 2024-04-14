// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

type Foo struct {
	Bar *Bar `p:"bar" v:"required-without:Baz"`
	Baz *Baz `p:"baz" v:"required-without:Bar"`
}
type Bar struct {
	BarKey string `p:"bar_key" v:"required"`
}
type Baz struct {
	BazKey string `p:"baz_key" v:"required"`
}

// https://github.com/gogf/gf/issues/2503
func Test_Issue2503(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		foo := &Foo{
			Bar: &Bar{BarKey: "value"},
		}
		err := getTestValid().StructNotCache(foo)
		t.Assert(err, nil)
	})
}

// https://github.com/gogf/gf/issues/1983
func Test_Issue1983_1(t *testing.T) {
	// RecRequiredError as the attribute Student in Teacher is an initialized struct, which has default value.
	gtest.C(t, func(t *gtest.T) {
		type Student56 struct {
			Name string `v:"required"`
			Age  int
		}
		type Teacher60 struct {
			Students Student56
		}
		var (
			teacher = Teacher60{}
		)
		err := getTestValid().StructNotCache(teacher)
		t.AssertNE(err, nil)
	})

}
func TestIssue_1983_2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Student74 struct {
			Name string `v:"required"`
			Age  int
		}
		type Teacher78 struct {
			Students *Student74
		}
		var (
			teacher = Teacher78{}
		)
		err := getTestValid().StructNotCache(teacher)
		t.Assert(err, nil)
	})
}

func TestIssue_1983_3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Student92 struct {
			Name string `v:"required"`
			Age  int
		}
		type Teacher96 struct {
			Students Student92
		}
		var (
			teacher = Teacher96{
				Students: Student92{
					Name: "john",
				},
			}
			//data    = g.Map{
			//	"LongName":     "john",
			//	"students": nil,
			//}
		)
		err := getTestValid().StructNotCache(teacher)
		t.Assert(err, nil)
	})
}

// https://github.com/gogf/gf/issues/1921
func Test_Issue1921(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type SearchOption119 struct {
			Size int `v:"max:100"`
		}
		type SearchReq122 struct {
			Option *SearchOption119 `json:"parseRuleOption,omitempty"`
		}

		var (
			req = SearchReq122{
				Option: &SearchOption119{
					Size: 10000,
				},
			}
		)
		err := getTestValid().StructNotCache(req)
		t.Assert(err, "The Size value `10000` must be equal or lesser than 100")
	})
}

// https://github.com/gogf/gf/issues/2011
func Test_Issue2011(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Student142 struct {
			Name string `v:"required|min-length:6"`
			Age  int
		}
		type Teacher146 struct {
			Student *Student142
		}
		var (
			teacher = Teacher146{
				Student: &Student142{
					Name: "john",
				},
			}
		)
		err := getTestValid().StructNotCache(teacher)
		t.Assert(err, "The Name value `john` length must be equal or greater than 6")
	})
}
