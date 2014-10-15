// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "log"
	"reflect"
)

// Value is one of: see.Address, string, int, float64, complex128, Circuit
type Value interface{}

func Copy(x Value) (y Value) {
	defer func() {
		if r := recover(); r != nil {
			y = x // If no Copy method, use Go copy semantic
		}
	}()
	return reflect.ValueOf(x).MethodByName("Copy").Call(nil)[0].Interface()
}

type sameness interface{
	Same(Value) bool
}

func Same(x, y Value) bool {
	if xx, ok := x.(sameness); ok {
		return xx.Same(y)
	}
	if yy, ok := y.(sameness); ok {
		return yy.Same(x)
	}
	return x == y
}
