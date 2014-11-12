// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"
	// "log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// Materialize
//
// On Before, Materialize receives:
//	{
//		Index Index	// namespace of values
//		Value Value	// value to materialize
//	}
//
// On After, Materialize sends:
//	{
//		Index Index
//		Value Value
//		Residue Value
//	}
//
type Materialize struct{}

func (Materialize) Spark(eye *be.Eye, _ *be.Matter, _ ...interface{}) Value {
	return nil
}

func (Materialize) CognizeBefore(eye *be.Eye, value interface{}) {
	u := value.(Circuit)
	index := be.AsIndex(u.At("Index"))
	v := u.At("Value")
	residual := be.Materialize(index, v)
	after :=  New().
		Grow("Index", Circuit(index)).
		Grow("Value", v).
		Grow("Residue", residual)
	// if len(reflex) > 0 {
	// 	after.Grow("Unconnected", reflex).Grow("u", reflex)
	// }
	eye.Show("After", after)
}

func (Materialize) CognizeAfter(eye *be.Eye, v interface{}) {}
