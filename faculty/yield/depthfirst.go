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

type DepthFirst struct{}

func (DepthFirst) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (DepthFirst) Cognize(eye *be.Eye, v interface{}) {
	depthFirst(eye, nil, v)
}

func (DepthFirst) CognizeFrame(eye *be.Eye, v interface{}) {}

func (DepthFirst) CognizeControl(eye *be.Eye, v interface{}) {}

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
		Grow("Path", (Address{walk}).Circuit()).
		Grow("Address", Address{walk}).
		Grow("Name", nm).
		Grow("View", x)

	eye.Show("Frame", frame)
	if len(walk) == 0 {
		eye.Show("Control", "End")
	}
}
