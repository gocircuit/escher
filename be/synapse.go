// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

// Cognize routines are called when a change in value is to be delivered to a reflex.
type Cognize func(value interface{})

// Synapse is the “wire” connecting two reflexes.
// It remembers the last value transmitted in order to stop propagation of same-value messages.
type Synapse struct {
	learn <-chan Cognize
	teach chan<- Cognize
	sync.Mutex
	q *ReCognizer
}

func NewSynapse() (x, y *Synapse) {
	xy, yx := make(chan Cognize, 1), make(chan Cognize, 1)
	x = &Synapse{
		learn: xy,
		teach: yx,
	}
	y = &Synapse{
		learn: yx,
		teach: xy,
	}
	return
}

func (m *Synapse) String() string {
	if m == nil {
		return "<nil>"
	}
	return "Synapse"
}

func (m *Synapse) Focus(cognize Cognize) *ReCognizer {
	m.teach <- cognize
	q := <-m.learn
	m.Lock()
	defer m.Unlock()
	m.q = &ReCognizer{q: q}
	return m.q
}

// Link attaches two synapse endpoints together.
func Link(m1, m2 *Synapse) {
	m1.teach <- <-m2.learn
}

// The two endpoints of a Synapse are ReCognizer objects.
type ReCognizer struct {
	q Cognize
	sync.Mutex
	memory interface{}
}

// ReCognize sends value to the reciprocal side of this synapse.
func (s *ReCognizer) ReCognize(value interface{}) {
	s.Lock()
	defer s.Unlock()
	if SameMeaning(s.memory, value) {
		return
	}
	s.memory = CopyMeaning(value)
	s.q(value)
}
