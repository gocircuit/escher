// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package yield

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("yield.Gates", be.NewMaterializer(Gates{}))
	faculty.Register("yield.Flows", be.NewMaterializer(Flows{}))
	faculty.Register("yield.DepthFirst", be.NewMaterializer(DepthFirst{}))
}
