package flex

import (
	"github.com/lowl11/flex/internal/th"
	"reflect"
)

func GetType[T any]() reflect.Type {
	return th.UnwrapN(reflect.TypeOf(new(T)), 1)
}
