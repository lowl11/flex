package robject

import (
	"reflect"
)

func (object RObject) Type() reflect.Type {
	return reflect.TypeOf(object.object)
}

func (object RObject) Value() reflect.Value {
	return reflect.ValueOf(object.object)
}

func (object RObject) IsFunc() bool {
	return object.Type().Kind() == reflect.Func
}

func (object RObject) IsPointer() bool {
	return object.Type().Kind() == reflect.Ptr
}

func (object RObject) IsStruct() bool {
	return object.Type().Kind() == reflect.Struct
}
