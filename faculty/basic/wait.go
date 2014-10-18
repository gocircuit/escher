// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/kit/plumb"
)

// Wait receives a stream of values on valve View.
// If a value matches the value given on valve For,
// it emits that value to the default valve.
type Wait struct { // : :For :View
	forValue plumb.Given
}

func (w *Wait) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	w.forValue.Init()
	return nil
}

func (w *Wait) CognizeFor(eye *be.Eye, value interface{}) {
	w.forValue.Fix(value)
}

func (w *Wait) CognizeView(eye *be.Eye, value interface{}) {
	if Same(value, w.forValue.Use()) {
		eye.Show(DefaultValve, value)
	}
}

func (w *Wait) Cognize(eye *be.Eye, value interface{}) {}
