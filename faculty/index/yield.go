// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
)

type Yield struct{ be.Sparkless }

func (Yield) CognizeIndex(eye *be.Eye, value interface{}) {
	yieldIndex(eye, value.(cir.Circuit), nil)
	eye.Show("End", value)
}

func yieldIndex(eye *be.Eye, x cir.Circuit, path []cir.Name) {
	for _, n := range x.SortedNames() {
		switch t := x.At(n).(type) {
		case cir.Circuit:
			if t.Vol() == 0 { // circuits without flow are treated as indices and recursed into
				yieldIndex(eye, t, append(path, n))
			} else {
				eye.Show(cir.DefaultValve, cir.New().Grow("Value", t).Grow("Address", cir.NewAddress(path...)))
			}
		default:
			eye.Show(cir.DefaultValve, cir.New().Grow("Value", t).Grow("Address", cir.NewAddress(path...)))
		}
	}
}

func (Yield) Cognize(*be.Eye, interface{}) {}

func (Yield) CognizeEnd(*be.Eye, interface{}) {}
