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
	"github.com/gocircuit/escher/kit/plumb"
	"github.com/gocircuit/escher/kit/memory"
)

// Memorize receives a circuit on valve Circuit and servers lookup requests against it on valve Lookup.
type Memorize struct {
	back plumb.Given
}

func (m *Memorize) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	m.back.Init()
	return &Memorize{}
}

// Circuit: { * }
func (m *Memorize) CognizeCircuit(eye *be.Eye, v interface{}) {
	m.back.Fix(v)
}

func (m *Memorize) CognizeLookup(eye *be.Eye, v interface{}) {
	eye.Show("Lookup", memory.Memory(m.back.Use().(Circuit)).Lookup(v.(Address)))
}
