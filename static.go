package flex

import (
	"flex/internal/th"
	"reflect"
)

func GetType[T any]() reflect.Type {
	return th.UnwrapN(reflect.TypeOf(new(T)), 1)
}
