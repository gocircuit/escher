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
	Reservoir is a gate with valves: Ground, Gate, Lookup and the empty-string.

	Reservoir accepts path-value pairs on valve Gate, of the form
	{
		Path { "a", "b", "c" }
		Value 3.14
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
	rsv.ground.Fix(v)
}

// { Path { "a", "b", "c"}; Value * }
func (rsv *Reservoir) CognizeGate(eye *be.Eye, v interface{}) {
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
	eye.Show(DefaultValve, ground)
}

func (rsv *Reservoir) CognizeLookup(eye *be.Eye, v interface{}) {
	rsv.lookup.Cognize(v)
}

func (rsv *Reservoir) Cognize(eye *be.Eye, v interface{}) {} // output valve
