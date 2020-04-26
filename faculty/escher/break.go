// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
)

type Breakpoint struct{ be.Sparkless }

func (Breakpoint) OverCognize(eye *be.Eye, valve cir.Name, value interface{}) {
	panic("Escher breakpoint")
}
