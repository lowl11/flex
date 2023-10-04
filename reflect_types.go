package flex

import (
	"flex/pkg/rfield"
	"flex/pkg/rfunc"
	"flex/pkg/robject"
	"flex/pkg/rstruct"
	"flex/pkg/rtype"
	"flex/pkg/rvalue"
	"reflect"
)

func Type(reflectType reflect.Type) RType {
	return rtype.New(reflectType)
}

func Value(reflectValue reflect.Value) RValue {
	return rvalue.New(reflectValue)
}

func Object(object any) RObject {
	return robject.New(object)
}

func Func(reflectFunc any) RFunc {
	return rfunc.New(reflectFunc)
}

func Method(reflectFunc reflect.Value) RFunc {
	return rfunc.NewMethod(reflectFunc)
}

func Struct(reflectStruct any) RStruct {
	return rstruct.New(reflectStruct)
}

func Field(field reflect.StructField) RField {
	return rfield.New(field)
}
