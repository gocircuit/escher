// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type DepthFirst struct{}

func (DepthFirst) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return DepthFirst{}
}

func (DepthFirst) CognizeCircuit(eye *be.Eye, v interface{}) {
	depthFirst(eye, nil, v)
}

func depthFirst(eye *be.Eye, walk []Name, v interface{}) {
	x, ok := v.(Circuit)
	if !ok {
		return
	}
	for n, v := range x.Gate {
		depthFirst(eye, append(walk, n), v)
	}
	eye.Show(
		DefaultValve, 
		New().
			Grow("Path", (Address{walk}).Circuit()).
			Grow("View", x),
	)
	if len(walk) == 0 {
		eye.Show(DefaultValve, nil) // EOF
	}
}

func (DepthFirst) Cognize(eye *be.Eye, v interface{}) {}
