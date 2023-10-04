package vh

import "reflect"

func Unwrap(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}

	elem := v.Elem()
	for {
		if elem.Kind() != reflect.Ptr {
			break
		}

		elem = elem.Elem()
	}

	return elem
}
