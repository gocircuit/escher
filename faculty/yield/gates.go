// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package yield

import (
	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
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
	u := value.(cir.Circuit)
	for name := range u.SortedNames() {
		frame := cir.New()
		frame.Grow("Name", name).Grow("Value", u.At(name))
		eye.Show("Frame", frame)
	}
	eye.Show("End", value)
}

func (Gates) CognizeFrame(eye *be.Eye, value interface{}) {}

func (Gates) CognizeEnd(eye *be.Eye, value interface{}) {}
