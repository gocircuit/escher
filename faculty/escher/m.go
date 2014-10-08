// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// M
type M struct{}

func (M) Spark(*be.Matter, ...interface{}) Value {
	return M{}
}

// In: { Memory Circuit; Design * }
func (M) CognizeIn(eye *be.Eye, v interface{}) {
	x := v.(Circuit)
	reflex, residual := be.Materialize(x.CircuitAt("Memory"), x.At("Design"))
	eye.Show(DefaultValve, New().Grow("Reflex", reflex).Grow("Residual", residual))
}

// In: ignored
// Out: { Reflex Reflex; Residual Circuit }
func (M) Cognize(*be.Eye, interface{}) {}
