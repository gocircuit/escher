// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"fmt"
)

// Vector ...
type Vector Circuit

func NewVector(gate, valve Name) Vector {
	return Vector(New().Grow("Gate", gate).Grow("Valve", valve))
}

func (v Vector) Reduce() (gate, valve Name) {
	return Circuit(v).At("Gate"), Circuit(v).At("Valve")
}

func (v Vector) Gate() Name {
	return Circuit(v).At("Gate")
}

func (v Vector) Valve() Name {
	return Circuit(v).At("Valve")
}

func (v Vector) Copy() Reducible {
	return Vector(Circuit(v).Clone())
}

func (v Vector) String() string {
	g, u := v.Reduce()
	return fmt.Sprintf("%v:%v", g, u)
}

func (v Vector) Same(x Reducible) bool {
	w, ok := x.(Vector)
	if !ok {
		return false
	}
	return Same(Circuit(v), Circuit(w))
}
