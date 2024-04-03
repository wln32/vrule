package vrule

import (
	"context"
)

func Struct(a any) error {
	return defaultValidator.Struct(a)
}

func StructCtx(ctx context.Context, a any) error {
	return defaultValidator.StructCtx(ctx, a)
}

func StructNotCache(a any) error {
	return defaultValidator.StructNotCache(a)
}

func ParseStruct(a any) *StructRule {
	rule := defaultValidator.ParseStruct(a, nil)
	return rule
}
