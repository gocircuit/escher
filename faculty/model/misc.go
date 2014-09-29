// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// ForkStart
func MaterializeForkStart() (be.Reflex, Value) {
	return be.MaterializeUnion("Circuit", "Vector")
}

// ForkView
func MaterializeForkView() (be.Reflex, Value) {
	return be.MaterializeUnion("Circuit", "Vector", "Index", "Depth", "Dir", "Path",)
}

// ForkVector
func MaterializeForkVector() (be.Reflex, Value) {
	return be.MaterializeUnion("Gate", "Valve")
}

// ForkMix
func MaterializeForkMix() (be.Reflex, Value) {
	return be.MaterializeUnion("Positive", "Negative")
}

// ForkRange
func MaterializeForkRange() (be.Reflex, Value) {
	return be.MaterializeUnion("Over", "With")
}

// ForkRangeView
func MaterializeForkRangeView() (be.Reflex, Value) {
	return be.MaterializeUnion("Name", "Value", "Count", "Index")
}
