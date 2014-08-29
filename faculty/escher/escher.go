// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	ns := faculty.Root.Refine("escher")
	ns.AddTerminal("CircuitDesignDir", CircuitDesignDir{})
}

// CircuitDesignDir
type CircuitDesignDir struct{}

func (CircuitDesignDir) Materialize(super *be.Super) be.Reflex {
	if super.Faculty.Genus == nil {
		panic("citcuit not from source directory")
	}
	return be.NewNounReflex(super.Faculty.Genus().Dir)
}
