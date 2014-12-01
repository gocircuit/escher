// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "log"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// The index gate is a design for a source reflex that returns a copy of the
// index contextual to the materialization of this gate.
type Index struct{}

func (n Index) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	index := matter.CircuitAt("Index")
	go func() {
		for vlv, _ := range matter.CircuitAt("View").Gate {
			eye.Show(vlv, index)
		}
	}()
	if matter.CircuitAt("View").Len() == 0 {
		return Circuit(n.Index)
	}
	return nil
}

func (n Index) OverCognize(*be.Eye, Name, interface{}) {}
