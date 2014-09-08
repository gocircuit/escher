// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"github.com/gocircuit/escher/see"
)

// Ignore gates ignore their empty-string valve
type Ignore struct{}

func (Ignore) Materialize(*Matter) Reflex {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize)
	}()
	return Reflex{see.DefaultValve: t}
}

func DontCognize(interface{}) {}

func NewNounReflex(v interface{}) Reflex {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize).ReCognize(v)
	}()
	return Reflex{see.DefaultValve: t}
}
