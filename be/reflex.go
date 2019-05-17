// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	cir "github.com/gocircuit/escher/circuit"
)

// Reflex is a bundle of not yet attached sense endpoints (synapses).
type Reflex map[cir.Name]*Synapse

//
type Materializer func(Reflex, cir.Circuit) interface{}

// Material represents a materializable object implemented as a Go type.
type Material interface {
	Spark(eye *Eye, matter cir.Circuit, aux ...interface{}) cir.Value // Initializer
}

type Sparkless struct{}

func (Sparkless) Spark(eye *Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	return nil
}
