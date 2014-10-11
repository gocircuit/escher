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
}

func (rsv *Reservoir) Spark(*be.Matter, ...interface{}) Value {
	rsv.Circuit = New()
	return Reservoir{}
}

func (rsv *Reservoir) CognizeGround(eye *be.Eye, v interface{}) {
	rsv.ground.Fix(v)
}

// { Path { "a", "b", "c"}; Value * }
func (rsv *Reservoir) CognizeGate(eye *be.Eye, v interface{}) {
	p := v.(Circuit).CircuitAt("Path")
	ground := rsv.ground.Use().(Circuit)
	u := ground
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
	eye.Show(DefaultValue, ground)
}

func (rsv *Reservoir) CognizeLookup(eye *be.Eye, v interface{}) {} ??

func (rsv *Reservoir) Cognize(eye *be.Eye, v interface{}) {} // output valve
