// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"
	"sync"
	// . "github.com/hoijui/escher/circuit"
)

// Cognize routines are called when a change in value is to be delivered to a reflex.
type Cognize func(value interface{})

// Nop cognizer for the synapse interface
func DontCognize(interface{}) {}

// Synapse is the “wire” connecting two reflexes.
// It remembers the last value transmitted in order to stop propagation of same-value messages.
type Synapse struct {
	accept <-chan Cognize
	offer  chan<- Cognize
	sync.Mutex
	recog *ReCognizer
}

func NewSynapse() (x, y *Synapse) {
	xy, yx := make(chan Cognize, 1), make(chan Cognize, 1)
	x = &Synapse{
		accept: xy,
		offer:  yx,
	}
	y = &Synapse{
		accept: yx,
		offer:  xy,
	}
	return
}

func (m *Synapse) String() string {
	if m == nil {
		return "<nil>"
	}
	return "Synapse"
}

func (m *Synapse) Connect(cognize Cognize) *ReCognizer {
	m.offer <- cognize
	close(m.offer)
	cog := <-m.accept
	m.Lock()
	defer m.Unlock()
	m.recog = &ReCognizer{cog: cog}
	return m.recog
}

// Link attaches two synapse endpoints together in one direction.
func Link(x, y *Synapse) {
	x.offer <- <-y.accept
}

// The two endpoints of a Synapse are ReCognizer objects.
type ReCognizer struct {
	cog Cognize
	// sync.Mutex
	// memory interface{}
}

// ReCognize sends value to the reciprocal side of this synapse.
func (s *ReCognizer) ReCognize(value interface{}) {
	// s.Lock()
	// defer s.Unlock()
	// if Same(s.memory, value) {
	// 	return
	// }
	// s.memory = Copy(value)
	s.cog(value)
}
