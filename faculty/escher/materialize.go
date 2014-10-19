// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// Materialize
type Materialize struct{}

func (Materialize) Spark(eye *be.Eye, _ *be.Matter, _ ...interface{}) Value {
	return nil
}

func (Materialize) CognizeBefore(eye *be.Eye, value interface{}) {
	v := value.(Circuit)
	mem := v.CircuitAt("Memory")
	op := v.At("Op")
	residual := be.Materialize(mem, op)
	after :=  New().
		Grow("Memory", mem).
		Grow("Op", op).
		Grow("Residual", residual)
	// if len(reflex) > 0 {
	// 	after.Grow("Unconnected", reflex).Grow("u", reflex)
	// }
	eye.Show("After", after)
}

func (Materialize) CognizeAfter(eye *be.Eye, v interface{}) {
	panic("time goes forward")
}
