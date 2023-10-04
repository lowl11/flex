package rtype

import "reflect"

type RType struct {
	reflectType reflect.Type
}

func New(reflectType reflect.Type) *RType {
	return &RType{
		reflectType: reflectType,
	}
}
