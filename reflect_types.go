package flex

import (
	"github.com/lowl11/flex/pkg/rfield"
	"github.com/lowl11/flex/pkg/rfunc"
	"github.com/lowl11/flex/pkg/robject"
	"github.com/lowl11/flex/pkg/rstruct"
	"github.com/lowl11/flex/pkg/rtype"
	"github.com/lowl11/flex/pkg/rvalue"
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
