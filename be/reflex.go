// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
	// . "github.com/gocircuit/escher/faculty"
)

// Reflex is a bundle of not yet attached sense endpoints (synapses).
type Reflex map[Name]*Synapse

type MaterializerFunc func() Reflex

type Materializer interface {
	Materialize() Reflex
}

type MaterializerWithMatter interface {
	Materialize(*Matter) Reflex
}

type Gate interface {
	Spark() // Initializer
}

// Matter describes the circuit context that commissioned the present materialization.
type Matter struct {
	Address Address // Address of the materialized design in memory
	Design interface{} // Design
	Valve map[Name]struct{} // Valves connected to this design in the enclosing circuit
	//
	Path []Name // Materialization path of this reflex, recursively following gate names
	Super *Matter // Matter of the circuit that recalled this reflex as a peer
}
