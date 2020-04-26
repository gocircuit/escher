// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package yield

import (
	//"fmt"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
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

	:End = frame
*/
type Flows struct{ be.Sparkless }

func sanitizeValue(u cir.Circuit, name cir.Name) cir.Value {
	value, ok := u.Gate[name]
	if !ok {
		return cir.NewAddress("missing")
	}
	if value == nil {
		return cir.NewAddress("nil")
	}
	return value
}

func (Flows) Cognize(eye *be.Eye, value interface{}) {
	u := value.(cir.Circuit)
	for xname, xview := range u.Flow {
		for xvalve, xvec := range xview {
			yname, yvalve := xvec.Gate, xvec.Valve
			//
			xvalue, yvalue := sanitizeValue(u, xname), sanitizeValue(u, yname)
			//
			frame := cir.New()
			xy := cir.New().Grow("Name", xname).Grow("Value", xvalue).Grow("Valve", xvalve)
			yx := cir.New().Grow("Name", xname).Grow("Value", yvalue).Grow("Valve", yvalve)
			frame.Grow(0, xy).Grow(1, yx)
			eye.Show("Frame", frame)
		}
	}
	eye.Show("End", value)
}

func (Flows) CognizeFrame(eye *be.Eye, value interface{}) {}

func (Flows) CognizeEnd(eye *be.Eye, value interface{}) {}
