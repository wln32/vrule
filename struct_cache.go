package vrule

import (
	"reflect"
	"sync"
)

type StructCache struct {
	cache map[string]*StructRule
	mu    sync.Mutex
}

func (s *StructCache) AddStructRule(v *StructRule) {

	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.cache[v.LongName]
	if !ok {
		s.cache[v.LongName] = v
	}
}
func (s *StructCache) GetStructRule(typ reflect.Type) *StructRule {

	name := getStructName(typ)
	rule, ok := s.cache[name]
	if ok {
		return rule
	}
	return rule
}

func (s *StructCache) GetStructRuleOrCreate(typ reflect.Type, v *Validator) *StructRule {

	name := getStructName(typ)
	rule, ok := s.cache[name]
	if ok {
		return rule
	}
	// 如果没有，就创建一个
	rule = s.createStructRule(typ, v)
	return rule
}

func getStructName(typ reflect.Type) string {

	name := typ.Name()
	if typ.PkgPath() != "" {
		name = typ.PkgPath() + "." + name
	} else {
		name = typ.String()
	}
	return name
}

func (s *StructCache) createStructRule(typ reflect.Type, v *Validator) *StructRule {

	structRule := v.ParseStruct(typ, v.cache)
	return structRule
}
