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
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("e.Materialize", be.NewNativeMaterializer(Materialize{}))
	faculty.Register("e.Parse", be.NewNativeMaterializer(Parse{}))
	// shortcuts
	faculty.Register("e.M", be.NewNativeMaterializer(Materialize{}))
	faculty.Register("e.P", be.NewNativeMaterializer(Parse{}))
}
