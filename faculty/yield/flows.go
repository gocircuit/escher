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
		0 { 
			Name Name
			Value Value
			Valve Valve
		}
		1 { 
			Name Name
			Value Value
			Valve Valve
		}
	}

	:Control = "End"
*/
type Flows struct{}

func (Flows) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func sanitizeValue(u Circuit, name Name) Value {
	value, ok := u.Gate[name]
	if !ok {
		return NewAddress("missing")
	}
	if value == nil {
		return NewAddress("nil")
	}
	return value
}

func (Flows) Cognize(eye *be.Eye, value interface{}) {
	u := value.(Circuit)
	for xname, xview := range u.Flow {
		for xvalve, xvec := range xview {
			yname, yvalve := xvec.Gate, xvec.Valve
			//
			xvalue, yvalue := sanitizeValue(u, xname), sanitizeValue(u, yname)
			//
			frame := New()
			xy := New().Grow("Name", xname).Grow("Value", xvalue).Grow("Valve", xvalve)
			yx := New().Grow("Name", xname).Grow("Value", yvalue).Grow("Valve", yvalve)
			frame.Grow(0, xy).Grow(1, yx)
			eye.Show("Frame", frame)
		}
	}
	eye.Show("Control", "End")
}

func (Flows) CognizeFrame(eye *be.Eye, value interface{}) {}

func (Flows) CognizeControl(eye *be.Eye, value interface{}) {}
