package rfunc

import (
	"errors"
	"flex/internal/th"
	"reflect"
)

func (f *RFunc) ArgumentsCount() int {
	return f.t.NumIn()
}

func (f *RFunc) Arguments() []reflect.Type {
	argCount := f.ArgumentsCount()
	arguments := make([]reflect.Type, 0, argCount)
	for i := 0; i < argCount; i++ {
		arguments = append(arguments, f.t.In(i))
	}
	return arguments
}

func (f *RFunc) ArgumentsUnwrapped() []reflect.Type {
	argCount := f.ArgumentsCount()
	arguments := make([]reflect.Type, 0, argCount)
	for i := 0; i < argCount; i++ {
		arguments = append(arguments, th.Unwrap(f.t.In(i)))
	}
	return arguments
}

func (f *RFunc) ReturnCount() int {
	return f.t.NumOut()
}

func (f *RFunc) Returns() []reflect.Type {
	returnCount := f.ReturnCount()
	returns := make([]reflect.Type, 0, returnCount)
	for i := 0; i < returnCount; i++ {
		returns = append(returns, f.t.Out(i))
	}
	return returns
}

func (f *RFunc) Call(arguments ...any) (returnValues []any, err error) {
	argumentTypes := f.Arguments()

	// check arguments count
	if len(arguments) != len(argumentTypes) {
		return nil, errors.New("Arguments count does not match")
	}

	// check arguments types
	for index, argType := range argumentTypes {
		if reflect.TypeOf(arguments[index]) != argType {
			return nil, errors.New("Argument types does not match")
		}
	}

	defer func() { // catch panic
		if recoverError := recover(); recoverError != nil {
			errorMessage, isString := recoverError.(string)
			if isString {
				err = errors.New(errorMessage)
			} else {
				inErr, isError := recoverError.(error)
				if isError {
					err = inErr
				}
			}
		}
	}()

	// prepare in arguments
	in := make([]reflect.Value, 0, len(arguments))
	for _, arg := range arguments {
		in = append(in, reflect.ValueOf(arg))
	}

	// call function
	values := f.v.Call(in)

	// convert result to any
	returnValues = make([]any, 0, len(values))
	for _, v := range values {
		returnValues = append(returnValues, v.Interface())
	}

	return
}
