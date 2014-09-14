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
	ns := faculty.Root.Refine("model")
	//
	ns.AddTerminal("Hamiltonian", Hamiltonian{})
	ns.AddTerminal("ForkStart", ForkStart{})
	ns.AddTerminal("ForkView", ForkView{})
	ns.AddTerminal("ForkVector", ForkVector{})
	//
	ns.AddTerminal("Reservoir", Reservoir{})
}
