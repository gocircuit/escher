// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package spin

import (
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// Move
type Move struct{}

func (Move) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Move) CognizeView(eye *be.Eye, v interface{}) {
	w := v.(Circuit)
	w.Include("Position", w.ComplexAt("Position") + w.ComplexAt("Orientation"))
	eye.Show(DefaultValve, w)
}

func (Move) Cognize(eye *be.Eye, v interface{}) {}
