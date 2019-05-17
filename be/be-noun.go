// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	cir "github.com/gocircuit/escher/circuit"
)

// Required matter: Noun
func materializeNoun(given Reflex, matter cir.Circuit) (residue interface{}) {
	noun := matter.At("Noun")
	for _, syn_ := range given {
		syn := syn_
		go func() {
			syn.Connect(DontCognize).ReCognize(noun)
		}()
	}
	if len(given) > 0 {
		return nil
	}
	return noun
}
