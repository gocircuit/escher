// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/be"
)

// ForkCharge
type ForkCharge struct{}

func (ForkCharge) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Circuit", "Image", "Valve")
}

// ForkSequence
type ForkSequence struct{}

func (ForkSequence) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "When", "Index", "Charge")
}
