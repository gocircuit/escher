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
	Flows emits one of the following two types of values to its default valve:
	{
		Is "Frame"
		Frame {
			0 { 
				Name *Name
				Value *Value
				Valve *Name
			}
			1 { 
				Name *Name
				Value *Value
				Valve *Name
			}
		}
	}

	{
		Is "Control"
		Control "End"
	}
*/
type Flows struct{}

func (Flows) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Flows) CognizeOfCircuit(eye *be.Eye, value interface{}) {
	u := value.(Circuit)
	for xname, xview := range u.Flow {
		for xvalve, xvec := range xview {
			yname, yvalve := xvec.Gate, xvec.Valve
			//
			r := New()
			r.Include("Is", "Frame")
			xy := New().Grow("Name", xname).Grow("Value", u.Gate[xname]).Grow("Valve", xvalve)
			yx := New().Grow("Name", xname).Grow("Value", u.Gate[yname]).Grow("Valve", yvalve)
			r.Include("Frame", New().Grow(0, xy).Grow(1, yx))
			eye.Show(DefaultValve, r)
		}
	}
	eye.Show(DefaultValve, New().Grow("Is", "Control").Grow("Control", "End"))
}

func (Flows) Cognize(eye *be.Eye, value interface{}) {}
