// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
)

// The index gate is a design for a source reflex that returns a copy of the
// index contextual to the materialization of this gate.
type Index struct{}

func (Index) Spark(eye *be.Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	index, view := matter.CircuitAt("Index"), matter.CircuitAt("View")
	go func() {
		for vlv := range view.Gate {
			eye.Show(vlv, index)
		}
	}()
	if view.Len() == 0 {
		return index
	}
	return nil
}

func (Index) OverCognize(*be.Eye, cir.Name, interface{}) {}
