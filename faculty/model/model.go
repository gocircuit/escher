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
	faculty.Register("model.ForkStart", MaterializeForkStart)
	faculty.Register("model.ForkView", MaterializeForkView)
	faculty.Register("model.ForkVector", MaterializeForkVector)
	//
	faculty.Register("model.Reservoir", &Reservoir{})
	//
	faculty.Register("model.Mix", &Mix{})
	faculty.Register("model.ForkMix", MaterializeForkMix)
	//
	faculty.Register("model.Range_", &Range{})
	faculty.Register("model.ForkRange", MaterializeForkRange)
	faculty.Register("model.ForkRangeView", MaterializeForkRangeView)
	//
	faculty.Register("model.IO", IO{})
}
