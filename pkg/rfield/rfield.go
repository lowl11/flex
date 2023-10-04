package rfield

import "reflect"

type RField struct {
	field reflect.StructField
}

func New(field reflect.StructField) *RField {
	return &RField{
		field: field,
	}
}
