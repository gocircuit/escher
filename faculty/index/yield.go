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

type Yield struct{}

func (Yield) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Yield) CognizeIndex(eye *be.Eye, value interface{}) {
	yieldIndex(eye, be.AsIndex(value), nil)
	eye.Show("End", "End")
}

func yieldIndex(eye *be.Eye, x be.Index, path []Name) {
	for _, n := range Circuit(x).SortedNames() {
		switch n.(type) {
		case int, string:
			switch t := Circuit(x).At(n).(type) {
			case Circuit:
				if be.IsIndex(t) {
					yieldIndex(eye, be.AsIndex(t), append(path, n))
				} else {
					eye.Show(DefaultValve, New().Grow("Value", t).Grow("Address", NewAddress(path...)))
				}
			default:
				eye.Show(DefaultValve, New().Grow("Value", t).Grow("Address", NewAddress(path...)))
			}
		default: // skip non-alphanemric names
		}
	}
}

func (Yield) Cognize(eye *be.Eye, value interface{}) {}

func (Yield) CognizeEnd(eye *be.Eye, value interface{}) {}
