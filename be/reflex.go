// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/faculty"
)

// Reflex is a bundle of not yet attached sense endpoints (synapses).
type Reflex map[Name]*Synapse

type Gate interface {
	Materialize() Reflex
}

type GateWithMatter interface {
	Materialize(*Matter) Reflex
}

// Matter describes the circuit context that commissioned the present materialization.
type Matter struct {
	Being string // Name of materializing being
	Address string // Address of this circuit design within the faculties namespace
	Design Circuit // Circuit design of this reflex
	Memory Faculty // Faculty within which this circuit design is implemented
	Super *Matter // Matter of the circuit that recalled this reflex as a peer
}

func (m *Matter) LastName() string {
	return m.Name[len(m.Name)-1]
}
