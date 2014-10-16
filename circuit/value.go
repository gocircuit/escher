// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "log"
)

// Value is one of: see.Address, string, int, float64, complex128, Circuit
type Value interface{}

func DeepCopy(x Value) (y Value) {
	switch t := x.(type) {
	case Circuit:
		return t.DeepCopy()
	case Address:
		return t.Copy()
	}
	return x
}

func Copy(x Value) (y Value) {
	switch t := x.(type) {
	case Circuit:
		return t.Copy()
	case Address:
		return t.Copy()
	}
	return x
}

func Same(x, y Value) bool {
	switch t := x.(type) {
	case Circuit:
		return t.Same(y)
	case Address:
		return t.Same(y)
	}
	return x == y
}
