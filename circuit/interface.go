// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "log"
)

// Name is one of: int or string
type Name interface{}

// Value is one of: see.Address, string, int, float64, complex128, Circuit
type Value interface{}

// Irreducible
type Irreducible interface {
	Copy() Irreducible
	Same(Irreducible) bool
}

func Copy(x Value) Value {
	switch t := x.(type) {
	case Irreducible:
		return t.Copy()
	}
	return x
}

func Same(x, y Value) bool {
	xr, x_ := x.(Irreducible)
	yr, y_ := y.(Irreducible)
	if x_ && y_ {
		return xr.Same(yr)
	}
	return x == y
}
