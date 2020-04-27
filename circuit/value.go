// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
// "log"
)

// Value is one of: string, int, float64, complex128, Circuit
type Value interface{}

// DeepCopy creates a deep copy of the supplied value
func DeepCopy(x Value) (y Value) {
	switch t := x.(type) {
	case Circuit:
		return t.DeepCopy()
	}
	return x
}

// Copy creates a shallow copy of the supplied value
func Copy(x Value) (y Value) {
	switch t := x.(type) {
	case Circuit:
		return t.Copy()
	}
	return x
}

// Same checks whether the two supplied values are the same.
// See `Value.Same()`
func Same(x, y Value) bool {
	switch t := x.(type) {
	case Circuit:
		return t.Same(y)
	}
	return x == y
}
