package ruleimpl

import (
	"context"
	"reflect"
)

type RuleFuncInput struct {
	// this value
	Value reflect.Value

	Message string
	// 关联字段的值
	// AssocFieldValues map[string]any
	// 结构体指针，用来实现关联字段的验证
	StructPtr reflect.Value
}

type ValidFunc interface {
	Run(ctx context.Context, input RuleFuncInput) error
}

type ValidFuncImpl func(ctx context.Context, input RuleFuncInput) error

func (f ValidFuncImpl) Run(ctx context.Context, input RuleFuncInput) error {
	return f(ctx, input)
}
