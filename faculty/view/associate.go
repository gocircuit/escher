// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package view

import (
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	// "github.com/gocircuit/escher/kit/plumb"
)

// Associate
type Associate struct{}

func (Associate) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Associate) CognizeView(eye *be.Eye, v interface{}) {
	w := v.(Circuit)
	eye.Show(DefaultValve, w.CircuitAt("Root").Clone().ReGrow(w.NameAt("Name"), w.At("Value")))
}

func (Associate) Cognize(eye *be.Eye, v interface{}) {}
