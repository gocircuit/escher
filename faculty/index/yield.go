// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Yield struct{ be.Sparkless }

func (Yield) CognizeIndex(eye *be.Eye, value interface{}) {
	yieldIndex(eye, value.(Circuit), nil)
	eye.Show("End", value)
}

func yieldIndex(eye *be.Eye, x Circuit, path []Name) {
	for _, n := range x.SortedNames() {
		switch t := x.At(n).(type) {
		case Circuit:
			if t.Vol() == 0 { // circuits without flow are treated as indices and recursed into
				yieldIndex(eye, t, append(path, n))
			} else {
				eye.Show(DefaultValve, New().Grow("Value", t).Grow("Address", NewAddress(path...)))
			}
		default:
			eye.Show(DefaultValve, New().Grow("Value", t).Grow("Address", NewAddress(path...)))
		}
	}
}

func (Yield) Cognize(*be.Eye, interface{}) {}

func (Yield) CognizeEnd(*be.Eye, interface{}) {}
