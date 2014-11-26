// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package view

import (
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	// "github.com/gocircuit/escher/kit/plumb"
)

// Focus
type Focus struct{}

func (Focus) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Focus) CognizeView(eye *be.Eye, v interface{}) {
	w := v.(Circuit)
	eye.Show(DefaultValve, w.CircuitAt("Root").At(w.NameAt("Name")))
}

func (Focus) Cognize(eye *be.Eye, v interface{}) {}
