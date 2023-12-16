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

	ty := v.Type()
	result := make(map[string]any)
	numField := v.NumField()
	for i := 0; i < numField; i++ {
		fieldValue := v.Field(i)
		fieldType := ty.Field(i)
		if !fieldType.IsExported() {
			continue
		}

		key := fieldType.Name
		result[key] = fieldValue.Interface()
	}

	return result, nil
}
