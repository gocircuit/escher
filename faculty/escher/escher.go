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
	faculty.Register("escher.Materialize", be.NewNativeMaterializer(Materialize{}))
	faculty.Register("escher.Index", be.NewNativeMaterializer(Index{}))
	faculty.Register("escher.Parse", be.NewNativeMaterializer(Parse{}))
	// faculty.Register("escher.Flatten", be.NewNativeMaterializer(Flatten{}))
}
