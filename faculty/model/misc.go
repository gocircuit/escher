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

// ForkStart
type ForkStart struct{}

func (ForkStart) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Circuit", "Vector")
}

// ForkView
type ForkView struct{}

func (ForkView) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Circuit", "Vector", "Index", "Depth", "Dir", "Series",)
}

// ForkVector
type ForkVector struct{}

func (ForkVector) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Gate", "Valve")
}

// ForkRange
type ForkRange struct{}

func (ForkRange) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Over", "With")
}

// ForkMix
type ForkMix struct{}

func (ForkMix) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Positive", "Negative")
}
