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

func (Flows) Cognize(eye *be.Eye, value interface{}) {
	u := value.(Circuit)
	for xname, xview := range u.Flow {
		for xvalve, xvec := range xview {
			yname, yvalve := xvec.Gate, xvec.Valve
			//
			frame := New()
			xy := New().Grow("Name", xname).Grow("Value", u.Gate[xname]).Grow("Valve", xvalve)
			yx := New().Grow("Name", xname).Grow("Value", u.Gate[yname]).Grow("Valve", yvalve)
			frame.Grow(0, xy).Grow(1, yx)
			eye.Show("Frame", frame)
		}
	}
	eye.Show("Control", "End")
}

func (Flows) CognizeFrame(eye *be.Eye, value interface{}) {}

func (Flows) CognizeControl(eye *be.Eye, value interface{}) {}
