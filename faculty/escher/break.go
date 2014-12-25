// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"
	// "log"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Breakpoint struct{ be.Sparkless }

func (Breakpoint) OverCognizeView(eye *be.Eye, valve Name, value interface{}) {
	panic("Escher breakpoint")
}
