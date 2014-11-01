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
	Gates emits one of the following two types of values to its default valve:
	{
		Is "Frame"
		Frame {
			Name *Name
			Value *Value
		}
	}

	{
		Is "Control"
		Control "End"
	}
*/
type Gates struct{}

func (Gates) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Gates) CognizeOfCircuit(eye *be.Eye, value interface{}) {
	u := value.(Circuit)
	for name, _ := range u.SortedNames() {
		r := New()
		r.Include("Is", "Frame")
		r.Include("Frame", New().Grow("Name", name).Grow("Value", u.At(name)))
		eye.Show(DefaultValve, r)
	}
	eye.Show(DefaultValve, New().Grow("Is", "Control").Grow("Control", "End"))
}

func (Gates) Cognize(eye *be.Eye, value interface{}) {}
