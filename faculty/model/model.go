// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package model provides a basis of gates for circuit transformations.
package model

import (
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("model.Hamiltonian", &Hamiltonian{})
	faculty.Register("model.Eulerian", &Eulerian{})
	faculty.Register("model.ForkStart", ForkStart{})
	faculty.Register("model.ForkView", ForkView{})
	faculty.Register("model.ForkVector", ForkVector{})
	//
	faculty.Register("model.Reservoir", &Reservoir{})
	//
	faculty.Register("model.Form", &Form{})
}
