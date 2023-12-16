package structuremap

import (
	"fmt"
	"reflect"
	"strings"
)

func Encode(val any) (map[string]any, error) {
	v := reflect.ValueOf(val)

	// recursively dereference
	v1, ok := asStructValue(v)
	if !ok {
		if !v1.IsValid() {
			return nil, fmt.Errorf("cannot encode nil value")
		}
		if v1.Kind() != reflect.Struct {
			return nil, fmt.Errorf("not a struct type (%T)", val)
		}
	}
	return encode(v)
}

// recursively dereference
func asStructValue(v reflect.Value) (reflect.Value, bool) {
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if !v.IsValid() {
		return reflect.Value{}, false
	}
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	return v, true
}

func encode(v reflect.Value) (map[string]any, error) {
	ty := v.Type()
	result := make(map[string]any)
	numField := v.NumField()
	for i := 0; i < numField; i++ {
		fieldValue := v.Field(i)
		fieldType := ty.Field(i)
		if !fieldType.IsExported() {
			continue
		}

		tag := parseStructTag(fieldType)
		if tag.ignore {
			continue
		}

		if tag.omitEmpty && fieldValue.IsZero() {
			continue
		}
		v2, ok := asStructValue(fieldValue)
		if !ok {
			result[tag.key] = fieldValue.Interface()
		} else {
			v2Encoded, err := encode(v2)
			if err != nil {
				return nil, fmt.Errorf("failed to encode value of field %s: %w", fieldType.Name, err)
			}
			result[tag.key] = v2Encoded
		}
	}

	return result, nil
}

type structTag struct {
	key       string
	omitEmpty bool

	ignore bool
}

func parseStructTag(field reflect.StructField) *structTag {
	tag := field.Tag.Get("structuremap")
	if tag == "" {
		return &structTag{
			key: field.Name,
		}
	}
	if tag == "-" {
		return &structTag{
			ignore: true,
		}
	}

	splitted := strings.Split(tag, ",")
	key := field.Name
	if len(splitted) > 0 {
		maybeKey := splitted[0]
		if maybeKey != "" {
			key = maybeKey
		}
	}

	isOmitEmpty := false
	if len(splitted) > 1 {
		maybeOmitempty := splitted[1]
		isOmitEmpty = maybeOmitempty == "omitempty"
	}
	return &structTag{
		key:       key,
		omitEmpty: isOmitEmpty,
	}
}
