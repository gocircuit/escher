// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"
	// "log"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// System
//
// On Before, System receives:
//	{
//		Index Index	// namespace of values
//		Value Value	// value to materialize
//	}
//
// On After, System sends:
//	{
//		Index Index
//		Value Value
//		Residue Value
//	}
//
type System struct {
	barrier Circuit
}

func (s *System) Spark(_ *be.Eye, matter Circuit, _ ...interface{}) Value {
	s.barrier = matter
	return nil
}

func (s *System) CognizeBefore(eye *be.Eye, value interface{}) {
	u := value.(Circuit)
	index, design := be.AsIndex(u.CircuitAt("Index")), u.At("Design")
	??
	residue := be.System(index, v, s.barrier)
	after := New().
		Grow("Index", Circuit(index)).
		Grow("Value", v).
		Grow("Residue", residue)
	// if len(reflex) > 0 {
	// 	after.Grow("Unconnected", reflex).Grow("u", reflex)
	// }
	eye.Show("After", after)
}

func (s *System) CognizeAfter(eye *be.Eye, v interface{}) {}
