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
	"github.com/gocircuit/escher/fs"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("escher.CircuitSourceDir", CircuitSourceDir{})
	faculty.Register("escher.Lookup", Lookup{})
	faculty.Register("escher.Memory", Memory{})
	faculty.Register("escher.Embody", &Embody{})
	faculty.Register("escher.Connect", &Connect{})
}

// Lookup
type Lookup struct{}

func (Lookup) Materialize() (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, dvalve string, dvalue interface{}) {
			if dvalve != "Address" {
				return
			}
			r := faculty.Root().Lookup(NewAddressParse(dvalue.(string)))
			eye.Show("Circuit", r.(Circuit))
		}, 
		"Address", "Circuit",
	)
	return reflex, Lookup{}
}

// Memory
type Memory struct{}

func (Memory) Materialize() (be.Reflex, Value) {
	return be.MaterializeNoun(faculty.Root())
}

// CircuitSourceDir
type CircuitSourceDir struct{}

func (CircuitSourceDir) Materialize(matter *be.Matter) (be.Reflex, Value) {
	return be.MaterializeNoun(matter.Super.Design.(Circuit).At(fs.Source{}).(Circuit).CircuitAt(0).StringAt("Dir"))
}
