package structuremap

import (
	"fmt"
	"reflect"
)

func Encode(val any) (map[string]any, error) {
	v := reflect.ValueOf(val)

	// recursively dereference
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if !v.IsValid() {
		return nil, nil
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct type (%T)", val)
	}

	// TODO: implement it
	return nil, nil
}
