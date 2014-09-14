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

// OrbitStart
type OrbitStart struct{}

func (OrbitStart) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Circuit", "Vector")
}

// OrbitView
type OrbitView struct{}

func (OrbitView) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Circuit", "Vector", "Index", "Depth", "Dir", "Series",)
}

// Vector_
type Vector_ struct{}

func (Vector_) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Gate", "Valve")
}
