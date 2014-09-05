// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package model provides a basis of gates for circuit transformations.
package model

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	ns := faculty.Root.Refine("model")
	ns.AddTerminal("ExploreInDepthOnStrobe", ExploreInDepthOnStrobe{})
	// ns.AddTerminal("ExploreInBreadthOnStrobe", ExploreInBreadthOnStrobe{})
}

// ExploreInDepthOnStrobe traverses the hierarchy of circuits induced by a given top-level/valveless circuit.
//
//	Strobe = {
//		When interface{}
//		Charge *understand.Circuit
//	}
// 	Sequence = {
//		When interface{} // When value that sparked this sequence
//		Index int // Index of this circuit within exploration sequence, 0-based
//		Charge *understand.Circuit // Current circuit in the exploration sequence
//	}
//
type ExploreInDepthOnStrobe struct{}

func (ExploreInDepthOnStrobe) Materialize() be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(CognizeExploreInDepthOnStrobe, "Strobe", "Sequence")
	return reflex
}

func CognizeExploreInDepthOnStrobe(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	if dvalve != "Strobe" {
		return
	}
	cir := dvalue.(Image).Interface("Charge")
	for ??
}
