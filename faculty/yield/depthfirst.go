// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package yield

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type DepthFirst struct{ be.Sparkless }

func (DepthFirst) Cognize(eye *be.Eye, v interface{}) {
	depthFirst(eye, nil, v)
}

func (DepthFirst) CognizeFrame(eye *be.Eye, v interface{}) {}

func (DepthFirst) CognizeEnd(eye *be.Eye, v interface{}) {}

func depthFirst(eye *be.Eye, walk []Name, v interface{}) {
	x, ok := v.(Circuit)
	if !ok {
		return
	}
	for _, n := range x.SortedNames() {
		switch n.(type) { // skip non alpha-numeric names
		case int, string:
			v := x.At(n)
			depthFirst(eye, append(walk, n), v)
		}
	}

	var nm Name = "" // The root circuit is shown with the empty name
	if len(walk) > 0 {
		nm = walk[len(walk)-1]
	}

	frame := New().
		Grow("Address", Circuit(NewAddress(walk...))).
		Grow("Name", nm).
		Grow("View", x)

	eye.Show("Frame", frame)
	if len(walk) == 0 {
		eye.Show("End", v)
	}
}
