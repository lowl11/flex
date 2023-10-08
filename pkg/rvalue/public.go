package rvalue

import (
	"github.com/lowl11/flex/internal/vh"
	"reflect"
)

func (v *RValue) Value() any {
	return v.reflectValue.Interface()
}

func (v *RValue) UnwrappedValue() any {
	return v.Unwrap().Interface()
}

func (v *RValue) Unwrap() reflect.Value {
	return vh.Unwrap(v.reflectValue)
}

func (v *RValue) Reset(newValue reflect.Value) {
	v.reflectValue = newValue
}
