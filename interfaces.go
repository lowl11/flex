package flex

import "reflect"

type RType interface {
	Reset(t reflect.Type)
	Unwrap() reflect.Type
	Type() reflect.Type
	IsPrimitive() bool
	IsReference() bool
	IsFunc() bool
	IsTime() bool
	IsStruct() bool
	IsPtr() bool
}

type RValue interface {
	Unwrap() reflect.Value
	Reset(newValue reflect.Value)
	Value() any
	UnwrappedValue() any
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
	FieldByType(t reflect.Type) any
	MethodCount() int
	Methods() []reflect.Method
	FieldValueByTag(tagName, tagValue string) any
}

type RField interface {
	IsPublic() bool
	Type() reflect.Type
	Tag(tagName string) []string
}
