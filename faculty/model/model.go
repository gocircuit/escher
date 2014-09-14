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
	ns.AddTerminal("Orbit", Orbit{})
	ns.AddTerminal("OrbitStart", OrbitStart{})
	ns.AddTerminal("OrbitView", OrbitView{})
	ns.AddTerminal("Vector", Vector_{})
	//
	ns.AddTerminal("Reservoir", Reservoir{})
}
