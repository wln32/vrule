package ruleimpl

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// boolMap defines the boolean values.
var boolMap = map[string]struct{}{
	"1":     {},
	"0":     {},
	"true":  {},
	"false": {},

	//"on":    {},
	//"yes":   {},
	//"":      {},
	//"off":   {},
	//"no":    {},
}

func JsonFormat(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.(string)
	if json.Valid([]byte(val)) {
		return nil
	}
	errMsg := strings.Replace(input.Message, "{value}", val, 1)
	return errors.New(errMsg)
}
func BooleanFormat(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.(string)
	if _, ok := boolMap[strings.ToLower(val)]; ok {
		return nil
	}
	errMsg := strings.Replace(input.Message, "{value}", val, 1)
	return errors.New(errMsg)
}
func IntegerFormat(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.(string)
	if _, err := strconv.Atoi(val); err == nil {
		return nil
	}
	errMsg := strings.Replace(input.Message, "{value}", val, 1)
	return errors.New(errMsg)
}
func FloatFormat(ctx context.Context, input RuleFuncInput) error {
	val := input.Value.(string)
	if _, err := strconv.ParseFloat(val, 10); err == nil {
		return nil
	}
	errMsg := strings.Replace(input.Message, "{value}", val, 1)
	return errors.New(errMsg)
}
