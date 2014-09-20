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
func MaterializeForkStart() be.Reflex {
	return basic.MaterializeUnion("_", "Circuit", "Vector")
}

// ForkView
func MaterializeForkView() be.Reflex {
	return basic.MaterializeUnion("_", "Circuit", "Vector", "Index", "Depth", "Dir", "Path",)
}

// ForkVector
func MaterializeForkVector() be.Reflex {
	return basic.MaterializeUnion("_", "Gate", "Valve")
}

// ForkMix
func MaterializeForkMix() be.Reflex {
	return basic.MaterializeUnion("_", "Positive", "Negative")
}

// ForkRange
func MaterializeForkRange() be.Reflex {
	return basic.MaterializeUnion("_", "Over", "With")
}

// ForkRangeView
func MaterializeForkRangeView() be.Reflex {
	return basic.MaterializeUnion("_", "Name", "Value", "Count", "Index")
}
