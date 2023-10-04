package rstruct

import (
	"errors"
	"flex/internal/th"
	"flex/internal/vh"
	"flex/pkg/rtype"
	"reflect"
)

type RStruct struct {
	v reflect.Value
	t reflect.Type
}

func New(reflectStruct any) *RStruct {
	s, err := NewErr(reflectStruct)
	if err != nil {
		panic(err)
	}

	return s
}

func NewErr(reflectStruct any) (*RStruct, error) {
	structType := reflect.TypeOf(reflectStruct)
	t := rtype.New(structType)
	if t.IsPtr() {
		t.Reset(t.Unwrap())
	}

	if !t.IsStruct() {
		return nil, errors.New("Given object is not struct")
	}

	return &RStruct{
		v: vh.Unwrap(reflect.ValueOf(reflectStruct)),
		t: th.Unwrap(reflect.TypeOf(reflectStruct)),
	}, nil
}
