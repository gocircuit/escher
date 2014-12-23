// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&System{}), "escher", "Materialize")
	faculty.Register(be.NewMaterializer(Index{}), "escher", "Index")
	faculty.Register(be.NewMaterializer(Parse{}), "escher", "Parse")
	faculty.Register(be.NewMaterializer(&Help{}), "help")
	faculty.Register(be.NewMaterializer(&Help{}), "Help")
}
