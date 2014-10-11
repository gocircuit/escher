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
	"github.com/gocircuit/escher/kit/fs"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("escher.Memorize", be.NewGateMaterializer(&Memorize{}))
	faculty.Register("escher.Faculties", Faculties)
	faculty.Register("escher.Materialize", be.NewGateMaterializer(&Materialize{}))
	faculty.Register("escher.Shell", be.NewGateMaterializer(&Shell{}))
	// faculty.Register("escher.CircuitSourceDir", CircuitSourceDir)
}

// Faculties
func Faculties() (be.Reflex, Value) {
	return be.MaterializeNoun(Circuit(faculty.Root()))
}

// CircuitSourceDir
func CircuitSourceDir(matter *be.Matter) (be.Reflex, Value) {
	return be.MaterializeNoun(matter.Super.Design.(Circuit).At(fs.Source{}).(Circuit).CircuitAt(0).StringAt("Dir"))
}
