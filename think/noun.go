// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	. "github.com/gocircuit/escher/image"
)

func DontCognize(Image) {}

func NewNounReflex(v Image) Reflex {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize).ReCognize(v)
	}()
	return Reflex{"": t}
}
