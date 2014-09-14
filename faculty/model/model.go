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
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

func init() {
	ns := faculty.Root.Refine("model")
	ns.AddTerminal("Orbit", Orbit{})
	ns.AddTerminal("ForkCharge", ForkCharge{})
	ns.AddTerminal("ForkSequence", ForkSequence{})
	ns.AddTerminal("Reservoir", Reservoir{})
}

/*
	Orbit traverses the hierarchy of circuits induced by a given top-level/valveless circuit.

	Spark = {
		Circuit Circuit
		Start Vector
	}

	Series = {
		Circuit Circuit // Current circuit in the exploration sequence
		Flow Flow
		Index int // Index of this circuit within exploration sequence, 0-based
		Depth int
		UpDown string
	}
*/
type Orbit struct{}

func CognizeSeries(*be.Eye, string, interface{}) {}

func CognizeSpark(eye *be.Eye, dvalve string, dvalue interface{}) {
	spark := dvalue.(Circuit)
	charge := spark.CircuitAt("Charge")
	//
	var start = view{
		Index: 0,
		Circuit: charge.CircuitAt("Circuit"),
		Gate: charge.StringAt("Gate"),
		Valve: charge.StringAt("Valve"),
	}
	var v = start
	var memory list.List
	for {
		// Current view
		eye.Show("Sequence", v.SequenceTerm(spark.At("When"))) // yield current view

		// transition
		entering := v.Circuit.At(v.Gate) // address of next image
		switch t := entering.(type) {
		case Address:
			//
			if memory.Len() > 100 {
				log.Fatalf("memory overload")
				// memory.Remove(memory.Front())
			}
			memory.PushFront(v) // remember
			//
			_, lookup := faculty.Root.LookupAddress(t.String())
			v.Circuit = lookup.(Circuit) // transition to next circuit
			toImg, toValve := v.Circuit.Follow(t.Name(), v.Valve)
			v.Gate, v.Valve = toImg.(string), toValve.(string)

		case Super:
			e := memory.Front() // backtrack
			if e == nil {
				log.Fatalf("insufficient memory")
			}
			u := e.Value.(view)
			memory.Remove(e)
			//
			v.Circuit = u.Circuit
			// log.Printf("Following %s:%s in %s", u.Gate, v.Valve, v.Circuit)
			toImg, toValve := v.Circuit.Follow(u.Gate, v.Valve)
			v.Gate, v.Valve = toImg.(string), toValve.(string)

		default:
			panic("unknown image meaning")
		}
		v.Index++
		//
		if v.Same(start) {
			x := v.SequenceTerm(spark.At("When"))
			x.Grow("Charge", "Loop")
			eye.Show("Sequence", x)
			break
		}
	}
}

// view ...
type view struct {
	Index int 
	Circuit Circuit // Ambient circuit
	Gate string // Focus peer
	Valve string // Focus valve
}

func (v view) UpDown() string {
	if _, ok := v.Circuit.At(v.Gate).(Super); ok {
		return "Up"
	}
	return "Down"
}

func (v view) SequenceTerm(when Meaning) Circuit {
	return New().
		Grow("When", when).
		Grow("Index", v.Index).
		Grow("Charge", 
			New().
			Grow("Circuit", v.Circuit).
			Grow("Gate", v.Gate).
			Grow("Valve", v.Valve).
			Grow("Dir", v.UpDown()),
		)
}

func (v view) Same(w view) bool {
	return Same(v.Circuit, w.Circuit) && v.Gate == w.Gate && v.Valve == w.Valve
}
