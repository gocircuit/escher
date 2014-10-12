// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/circuit"
)

// TODO: Add test

/*
	Reservoir has valves Ground, Change, Lookup and the empty-string.

	Reservoir receives a ground circuit on valve Ground as a prerequisite.
	Valve Lookup must be connected to a lookup server, like escher.Memory.

	Reservoir accepts a stream of updates on valve Change and applies them
	to the ground circuit. After a “flush” update, the ground circuit is sent out
	the default valve.

	Updates are of the form:

		{
			Op "Gate"			// Indicates gate update
			Path { "a", "b", "c" }	// Path to gate in expanded circuit
			Value 3.14			// Value of gate
		}

	Or,

		{
			Op "Flush"
		}

*/
type Reservoir struct {
	ground plumb.Given
	lookup plumb.Client
}

func (rsv *Reservoir) Spark(eye *be.Eye, _ *be.Matter, _ ...interface{}) Value {
	rsv.ground.Init()
	rsv.lookup.Init(
		func (v interface{}) {
			eye.Show("Lookup", v)
		},
	)
	return Reservoir{}
}

func (rsv *Reservoir) CognizeGround(eye *be.Eye, v interface{}) {
	rsv.ground.Fix(v.(Circuit).Clone())
}

func (rsv *Reservoir) CognizeChange(eye *be.Eye, v interface{}) {
	switch v.(Circuit).StringAt("Op") {
	case "Flush":
		rsv.cognizeFlushChange(eye, v)
	case "Gate":
		rsv.cognizeGateChange(eye, v)
	default:
		panic(1)
	}
}

func (rsv *Reservoir) cognizeFlushChange(eye *be.Eye, v interface{}) {
	ground := rsv.ground.Use().(Circuit)
	rsv.ground.Flush()
	eye.Show(DefaultValve, ground)
}

func (rsv *Reservoir) cognizeGateChange(eye *be.Eye, v interface{}) {
	p := v.(Circuit).CircuitAt("Path")
	ground := rsv.ground.Use().(Circuit)
	view := ground
	numbers := p.SortedNumbers()
	for i := 0; i + 1 < len(numbers); i++ {
		name := p.Gate[numbers[i]] // name of next gate in the walk
		switch t := view.Gate[name].(type) {
		case Circuit:
			view = t
		default:
			y := New()
			view.Include(name, y)
			view = y
		}
	}
	view.Include(p.Gate[numbers[len(numbers)-1]], v.(Circuit).At("Value"))
}

func (rsv *Reservoir) CognizeLookup(eye *be.Eye, v interface{}) {
	rsv.lookup.Cognize(v)
}

func (rsv *Reservoir) Cognize(eye *be.Eye, v interface{}) {} // output valve
