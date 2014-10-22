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
	faculty.Register("escher.Parse", be.NewNativeMaterializer(Parse{}))
	// reservoir 
	faculty.Register("escher.Reservoir", be.NewNativeMaterializer(&ReservoirNoun{}))
	faculty.Register("escher.Put", be.NewNativeMaterializer(&ReservoirVerb{}))
	faculty.Register("escher.Get", be.NewNativeMaterializer(&ReservoirVerb{}))
	faculty.Register("escher.Forget", be.NewNativeMaterializer(&ReservoirVerb{}))
	//
	faculty.Register("escher.Shadow", be.NewNativeMaterializer(ReservoirShadow{}))
	// memory
	faculty.Register("Memory", be.NewNativeMaterializer(&ReservoirNoun{}, faculty.Root()))
	faculty.Register("Put", be.NewNativeMaterializer(&ReservoirVerb{}, faculty.Root()))
	faculty.Register("Get", be.NewNativeMaterializer(&ReservoirVerb{}, faculty.Root()))
	faculty.Register("Forget", be.NewNativeMaterializer(&ReservoirVerb{}, faculty.Root()))
}
