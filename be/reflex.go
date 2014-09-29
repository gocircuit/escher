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

//
type Materializer interface {
	Materialize() (Reflex, Value)
}

//
type MaterializerWithMatter interface {
	Materialize(*Matter) (Reflex, Value)
}

//
type Gate interface {
	Spark() Value // Initializer
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

func (m *Matter) Circuit() Circuit {
	return New().
		Grow("Address", m.Address).
		Grow("Design", m.Design).
		Grow("Valve", valveSetCircuit(m.Valve)).
		Grow("Path", pathCircuit(m.Path))
}

func valveSetCircuit(s map[Name]struct{}) Circuit {
	r := New()
	var i int
	for v, _ := range s {
		r.Gate[i] = v
		i++
	}
	return r
}

func pathCircuit(p []Name) Circuit {
	r := New()
	for i, n := range p {
		r.Grow(i, n)
	}
	return r
}
