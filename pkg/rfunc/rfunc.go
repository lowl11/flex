package rfunc

import (
	"errors"
	"github.com/lowl11/flex/pkg/rtype"
	"reflect"
)

type RFunc struct {
	v reflect.Value
	t reflect.Type
}

func New(f any) *RFunc {
	inst, err := NewErr(f)
	if err != nil {
		panic(err)
	}

	return inst
}

func NewMethod(f reflect.Value) *RFunc {
	inst, err := NewErr(f)
	if err != nil {
		panic(err)
	}

	return inst
}

func NewErr(f any) (*RFunc, error) {
	if v, isValue := f.(reflect.Value); isValue {
		return &RFunc{
			v: v,
			t: v.Type(),
		}, nil
	}

	if !rtype.New(reflect.TypeOf(f)).IsFunc() {
		return nil, errors.New("Given object is not func")
	}

	return &RFunc{
		v: reflect.ValueOf(f),
		t: reflect.TypeOf(f),
	}, nil
}
