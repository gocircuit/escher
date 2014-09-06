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
	"github.com/gocircuit/escher/understand"
)

func init() {
	ns := faculty.Root.Refine("model")
	ns.AddTerminal("ExploreOnStrobe", ExploreOnStrobe{})
}

// ExploreOnStrobe traverses the hierarchy of circuits induced by a given top-level/valveless circuit.
//
//	Strobe = {
//		When interface{}
//		Charge {
//			Circuit *understand.Circuit
//			Peer interface{} // Start peer name
//			Valve string // Start valve name
//		}
//	}
//
// 	Sequence = {
//		When interface{} // When value that sparked this sequence
//		Index int // Index of this circuit within exploration sequence, 0-based
//		Charge {
//			Circuit *understand.Circuit // Current circuit in the exploration sequence
//			Peer interface{} // Point-of-view peer
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
	img := dvalue.(Image)
	charge := img.Image("Charge")
	//
	var start = view{
		Circuit: charge.Interface("Circuit").(*understand.Circuit),
		Peer: charge.Interface("Peer"),
		Valve: charge.String("Valve"),
	}
	var v = start
	var n int // Number of steps
	for {
		???
		eye.Show( // yield
			"Sequence", 
			Image{
				"When": img.Interface("When")),
				"Index": n,
				"Charge": Image{
					"Circuit": v.Circuit,
					"Peer": v.Peer,
					"Valve": v.Valve,
				},
			},
		)
		n++
		// transition
		designPath := v.Circuit.PeerByName(v.Peer).Design().(see.RootPath) // gates are not allowed
		_, recall := faculty.Root.Walk(designPath)
		v.Circuit = recall.(*understand.Circuit) // cannot jump to gates
		v.Peer = v.Circuit.PeerByName("").ValveByName(v.Valve).Matching.Of.Name()
		v.Valve = ??


		if v == start {
			break
		}
	}
}

// view ...
type view struct {
	Circuit *understand.Circuit // Ambient circuit
	Peer interface{} // Focus peer
	Valve string // Focus valve
}
