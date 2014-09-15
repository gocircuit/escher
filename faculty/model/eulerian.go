// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

// import (
// 	"container/list"
// 	// "fmt"
// 	"log"

// 	"github.com/gocircuit/escher/faculty"
// 	. "github.com/gocircuit/escher/circuit"
// 	"github.com/gocircuit/escher/be"
// )

/*
	Eulerian traverses the hierarchy of circuits induced by a given top-level/valveless circuit.

	Start = {
		Circuit Circuit
	}

	View = {
		Circuit Circuit // Current circuit in the exploration sequence
		Vector Vector
		Index int // Index of this circuit within exploration sequence, 0-based
		Depth int
		Dir string
		Path string // Loop
	}
*/
// type Eulerian struct{}

// func (Eulerian) Spark() {}

// func (Eulerian) CognizeView(*be.Eye, interface{}) {}

// func (Eulerian) CognizeStart(eye *be.Eye, dv interface{}) {
// 	ch := make(chan Circuit)
// 	var in = dv.(Circuit)
// 	??
// }

// func euler(eye *be.Eye, v view) {
// 	// eye.Show("View", v.Circuitize())
// 	for g, h := range x.Gates() {
// 		switch t := h.(type) {
// 		case Address:
// 			_, lookup := faculty.Root.LookupAddress(t.String())
// 			if lookup == nil {
// 				log.Fatalf("No circuit with address %s", t.String())
// 			}
// 			?? = lookup.(Circuit)
// 			??
// 		default: // skip non-address gates as data
// 		}
// 	}
// }
