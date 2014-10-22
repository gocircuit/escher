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
	. "github.com/gocircuit/escher/kit/reservoir"
)

// ReservoirShadow
type ReservoirShadow struct{}

func (ReservoirShadow) Spark(eye *be.Eye, _ *be.Matter, aux ...interface{}) Value {
	return nil
}

func (ReservoirShadow) CognizeView(eye *be.Eye, value interface{}) {
	v := value.(Circuit)
	eye.Show(DefaultValve, Shadow(v.At("Reservoir").(Reservoir), v.CircuitAt("Shadow")))
}

func (ReservoirShadow) Cognize(eye *be.Eye, value interface{}) {}
