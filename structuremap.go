package structuremap

import "reflect"

func Encode(val any) (map[string]any, error) {
	v := reflect.ValueOf(val)
	if v.IsNil() {
		return nil, nil
	}

	// TODO: implement it
	return nil, nil
}
