package structuremap

import (
	"fmt"
	"reflect"
	"strings"
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

		tag := parseStructTag(fieldType)
		if tag.ignore {
			continue
		}

		if tag.omitEmpty && fieldValue.IsZero() {
			continue
		}
		result[tag.key] = fieldValue.Interface()
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
