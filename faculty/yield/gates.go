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

/*
	:Frame = {
		Name Name
		Value Value
	}

	:End = frame
*/
type Gates struct{ be.Sparkless }

func (Gates) Cognize(eye *be.Eye, value interface{}) {
	u := value.(Circuit)
	for name, _ := range u.SortedNames() {
		frame := New()
		frame.Grow("Name", name).Grow("Value", u.At(name))
		eye.Show("Frame", frame)
	}
	eye.Show("End", value)
}

func (Gates) CognizeFrame(eye *be.Eye, value interface{}) {}

func (Gates) CognizeEnd(eye *be.Eye, value interface{}) {}
