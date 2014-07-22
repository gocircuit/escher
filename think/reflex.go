// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	"github.com/gocircuit/escher/tree"
)

// Reflex is a bundle of un-attached sense endpoints
type Reflex map[tree.Name]*Synapse

type Gate interface {
	Materialize() Reflex
}

// Ignore gates ignore their empty-string valve
type Ignore struct{}

func (Ignore) Materialize() Reflex {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize)
	}()
	return Reflex{"": t}
}
