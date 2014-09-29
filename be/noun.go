// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
)

// Ignore gates ignore their empty-string valve
type Ignore struct{}

func (Ignore) Materialize(*Matter) (Reflex, Value) {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize)
	}()
	return Reflex{DefaultValve: t}, nil
}

func DontCognize(interface{}) {}

func MaterializeNoun(v interface{}) (Reflex, Value) {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize).ReCognize(v)
	}()
	return Reflex{DefaultValve: t}, t
}
