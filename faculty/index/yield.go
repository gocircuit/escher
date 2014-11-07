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

/*
	:Frame = {
		Name Name
		Value Value
	}

	:Control = "End"
*/
type Yield struct{}

func (Yield) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Yield) CognizeIndex(eye *be.Eye, value interface{}) {
	yieldIndex(eye, value.(be.Index))
	eye.Show("Control", "End")
}

func yieldIndex(eye *be.Eye, x be.Index) {
	for _, n := range Circuit(x).SortedNames() {
		switch n.(type) {
		case int, string:
			switch t := Circuit(x).At(n).(type) {
			case be.Index:
				yieldIndex(eye, t)
			default:
				eye.Show(DefaultValve, t)
			}
		default: // skip non-alphanemric names
		}
	}
}

func (Yield) Cognize(eye *be.Eye, value interface{}) {}

func (Yield) CognizeControl(eye *be.Eye, value interface{}) {}
