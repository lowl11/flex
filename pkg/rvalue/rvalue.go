package rvalue

import "reflect"

type RValue struct {
	reflectValue reflect.Value
}

func New(reflectValue reflect.Value) *RValue {
	return &RValue{
		reflectValue: reflectValue,
	}
}
