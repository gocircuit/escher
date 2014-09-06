// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal(see.Name("Strobe"), Strobe{})
}

// Strobe ...
type Strobe struct{}

func (Strobe) Materialize() be.Reflex {
	x := &strobe{
		when: make(chan interface{}, 1), // whens and charges cannot be out of order by more than one slot
		charge: make(chan interface{}, 1),
	}
	reflex, _ := plumb.NewEyeCognizer(x.Cognize, "Charge", "When", "Strobe")
	return reflex
}

type strobe struct {
	when, charge chan interface{}
}

func (x *strobe) Cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	switch dvalve {
	case "Charge":
		select {
		case y := <- x.when: // if a when is already waiting, couple it with the charge and send a strobe pair
			eye.Show("Strobe", Make().Grow("When", y).Grow("Charge", dvalue))
		default: // otherwise remember the charge
			x.charge <- dvalue
		}
	case "When":
		select {
		case y := <- x.charge: // if a charge is already waiting, couple it with the when and send a strobe pair
			eye.Show("Strobe", Make().Grow("When", dvalue).Grow("Charge", y))
		default: // otherwise remember the when
			x.when <- dvalue
		}
	}
}
