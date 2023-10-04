package rstruct

import (
	"flex/pkg/rtype"
	"reflect"
)

func (s RStruct) FieldCount() int {
	return s.t.NumField()
}

func (s RStruct) Fields() []reflect.StructField {
	fieldCount := s.FieldCount()

	fields := make([]reflect.StructField, 0, fieldCount)
	for i := 0; i < fieldCount; i++ {
		fields = append(fields, s.t.Field(i))
	}

	return fields
}

func (s RStruct) FieldsRow() []reflect.StructField {
	fieldCount := s.FieldCount()

	fields := make([]reflect.StructField, 0, fieldCount)
	for i := 0; i < fieldCount; i++ {
		field := s.t.Field(i)
		t := rtype.New(field.Type)
		if t.IsPtr() {
			t.Reset(t.Unwrap())
		}

		if t.IsStruct() && !t.IsTime() {
			fields = append(fields, New(reflect.New(t.Type()).Interface()).FieldsRow()...)
			continue
		}

		fields = append(fields, field)
	}

	return fields
}

func (s RStruct) MethodCount() int {
	return s.t.NumMethod()
}

func (s RStruct) Methods() []reflect.Method {
	methodCount := s.MethodCount()
	methods := make([]reflect.Method, 0, methodCount)
	for i := 0; i < methodCount; i++ {
		methods = append(methods, s.t.Method(i))
	}
	return methods
}
