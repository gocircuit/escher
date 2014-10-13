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

func (m *Memorize) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	m.back.Init()
	go func() {
		eye.Show("Lookup", m)
	}()
	return nil
}

// Circuit: { * }
func (m *Memorize) CognizeCircuit(eye *be.Eye, v interface{}) {
	m.back.Fix(memory.Memory(v.(Circuit)))
}

func (m *Memorize) CognizeLookup(eye *be.Eye, v interface{}) {}

func (m *Memorize) Lookup(addr Address) Value {
	return m.back.Use().(memory.Memory).Lookup(addr)
}
