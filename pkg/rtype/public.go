package rtype

import (
	"github.com/lowl11/flex/internal/th"
	"reflect"
)

func (t *RType) Reset(newType reflect.Type) {
	t.reflectType = newType
}

func (t *RType) Type() reflect.Type {
	return t.reflectType
}

func (t *RType) IsPrimitive() bool {
	nativeType := th.Unwrap(t.reflectType)
	switch nativeType.Kind() {
	case reflect.Ptr, reflect.Struct, reflect.Interface,
		reflect.Slice, reflect.Array, reflect.Map:
		return false
	}

	return true
}

func (t *RType) IsReference() bool {
	return !t.IsPrimitive()
}

func (t *RType) IsFunc() bool {
	return t.reflectType.Kind() == reflect.Func
}

func (t *RType) IsStruct() bool {
	return t.reflectType.Kind() == reflect.Struct
}

func (t *RType) IsTime() bool {
	return t.reflectType.String() == "time.Time"
}

func (t *RType) IsUUID() bool {
	return t.reflectType.String() == "uuid.UUID"
}

func (t *RType) IsPtr() bool {
	return t.reflectType.Kind() == reflect.Ptr
}

func (t *RType) Unwrap() reflect.Type {
	return th.Unwrap(t.reflectType)
}

func (t *RType) IsInterface() bool {
	return t.reflectType.Kind() == reflect.Interface
}
