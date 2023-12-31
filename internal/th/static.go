package th

import "reflect"

func Unwrap(t reflect.Type) reflect.Type {
	if t.Kind() != reflect.Ptr {
		return t
	}

	elem := t.Elem()
	for {
		if elem.Kind() != reflect.Ptr {
			break
		}

		elem = elem.Elem()
	}

	return elem
}

func UnwrapN(t reflect.Type, n int) reflect.Type {
	if t.Kind() != reflect.Ptr {
		return t
	}

	elem := t.Elem()
	for n > 0 {
		n--
		if elem.Kind() != reflect.Ptr {
			break
		}

		elem = elem.Elem()
	}

	return elem
}

func IsPtr(object any) bool {
	return reflect.TypeOf(object).Kind() == reflect.Ptr
}
