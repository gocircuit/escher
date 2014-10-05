// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package think provides a basis of four fundamental gates for manipulating thinkable images.
package think

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

func init() {
	//
	faculty.Register("think.Associate", Associate{})
	faculty.Register("think.Remember", Remember{})
	faculty.Register("think.Choose", Choose{})
	faculty.Register("think.Focus", Focus{})
	//
	faculty.Register("think.A", Associate{})
	faculty.Register("think.R", Remember{})
	faculty.Register("think.C", Choose{})
	faculty.Register("think.F", Focus{})
}

// Focus
type Focus struct{}

func (Focus) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer((&focus{}).Cognize, "From", "On", "When", DefaultValve)
	return reflex
}

type focus struct {
	sync.Mutex
	from Circuit
	on string
	when interface{}
}

func (x *focus) Cognize(eye *be.Eye, dvalve string, dvalue interface{}) {
	x.Lock()
	defer x.Unlock()
	switch dvalve {
	case "From":
		x.from = dvalue.(Circuit)
	case "On":
		x.on = dvalue.(string)
	case "When":
		x.when = dvalue
	case DefaultValve:
	default:
		panic("eh")
	}
	eye.Show(
		DefaultValve, 
		Circuit{
			"Focus": x.from.Copy().Cut(x.on),
			"When": x.when,
		},
	)
}
