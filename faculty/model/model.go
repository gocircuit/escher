// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package model provides a basis of gates for circuit transformations.
package model

import (
	"container/list"
	// "fmt"
	"log"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/faculty/basic"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	ns := faculty.Root.Refine("model")
	ns.AddTerminal("ExploreOnStrobe", ExploreOnStrobe{})
	ns.AddTerminal("ForkCharge", ForkCharge{})
	ns.AddTerminal("ForkSequence", ForkSequence{})
}

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

// ExploreOnStrobe traverses the hierarchy of circuits induced by a given top-level/valveless circuit.
//
//	Strobe = {
//		When *
//		Charge {
//			Circuit Circuit
//			Image string // Start peer name
//			Valve string // Start valve name
//		}
//	}
//
// 	Sequence = {
//		When * // When value that sparked this sequence
//		Index int // Index of this circuit within exploration sequence, 0-based
//		Charge {
//			Circuit Circuit // Current circuit in the exploration sequence
//			Image string // Point-of-view peer
//			Valve string // Point-of-view valve of pov peer
//		}
//	}
//
type ExploreOnStrobe struct{}

func (ExploreOnStrobe) Materialize() be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(CognizeExploreOnStrobe, "Strobe", "Sequence")
	return reflex
}

func CognizeExploreOnStrobe(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	if dvalve != "Strobe" {
		return
	}
	strobe := dvalue.(Circuit)
	charge := strobe.CircuitAt("Charge")
	//
	var start = view{
		Circuit: charge.CircuitAt("Circuit"),
		Image: charge.StringAt("Image"),
		Valve: charge.StringAt("Valve"),
	}
	var v = start
	var n int // Number of steps
	var memory list.List
	for {
		u := New().
			Grow("When", strobe.AtNil("When")).
			Grow("Index", n).
			Grow(
				"Charge", 
				New().Grow("Circuit", v.Circuit).Grow("Image", v.Image).Grow("Valve", v.Valve),
			)
		//
		eye.Show("Sequence", u) // yield current view
		//
		if memory.Len() > 1e9 {
			log.Printf("memory overload")
			memory.Remove(memory.Front())
		}
		memory.PushFront(v) // remember
		n++

		// transition
		entering := v.Circuit.AtNil(v.Image) // address of next image
		switch t := entering.(type) {
		case Address:
			_, lookup := faculty.Root.Lookup(t.Walk()...)
			v.Circuit = lookup.(Circuit) // transition to next circuit
			toImg, toValve := v.Circuit.Follow(t.Name(), v.Valve)
			v.Image, v.Valve = toImg.(string), toValve.(string)

		case Super:
			e := memory.Front() // backtrack
			if e == nil {
				log.Fatalf("insufficient memory")
			}
			u := e.Value.(view)
			memory.Remove(e)
			v.Circuit = u.Circuit
			toImg, toValve := v.Circuit.Follow(u.Image, v.Image)
			v.Image, v.Valve = toImg.(string), toValve.(string)

		default:
			panic("unknown image meaning")
		}
		if v == start {
			break
		}
	}
}

// view ...
type view struct {
	Circuit Circuit // Ambient circuit
	Image string // Focus peer
	Valve string // Focus valve
}
