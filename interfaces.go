package flex

import "reflect"

type RType interface {
	IsPrimitive() bool
	IsReference() bool
	IsFunc() bool
	IsTime() bool
	IsStruct() bool
	IsPtr() bool
	Unwrap() reflect.Type
}

type RValue interface {
	//
}

type RObject interface {
	Type() reflect.Type
	Value() reflect.Value
	IsFunc() bool
	IsPointer() bool
	IsStruct() bool
}

type RFunc interface {
	ArgumentsCount() int
	Arguments() []reflect.Type
	ArgumentsUnwrapped() []reflect.Type
	ReturnCount() int
	Returns() []reflect.Type
	Call(arguments ...any) (returnValues []any, err error)
}

type RStruct interface {
	FieldCount() int
	Fields() []reflect.StructField
	FieldsRow() []reflect.StructField
	MethodCount() int
	Methods() []reflect.Method
}

type RField interface {
	Tag(tagName string) []string
}
