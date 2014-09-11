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
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	ns := faculty.Root.Refine("escher")
	ns.AddTerminal("CircuitSourceDir", CircuitSourceDir{})
	ns.AddTerminal("Lookup", Lookup{})
}

type Lookup struct{}

func (Lookup) Materialize() be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(
		func(eye *plumb.Eye, dvalve string, dvalue interface{}) {
			if dvalve != "Address" {
				return
			}
			_, r := faculty.Root.LookupAddress(dvalue.(string))
			eye.Show("Circuit", r.(Circuit))
		}, 
		"Address", "Circuit",
	)
	return reflex
}

// CircuitSourceDir
type CircuitSourceDir struct{}

func (CircuitSourceDir) Materialize(matter *be.Matter) be.Reflex {
	return be.NewNounReflex(matter.Design.At(faculty.Genus_{}).(*faculty.CircuitGenus).Dir)
}
