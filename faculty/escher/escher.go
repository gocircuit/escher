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
	"github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	ns := faculty.Root.Refine("escher")
	ns.AddTerminal(see.Name("CircuitDesignDir"), CircuitDesignDir{})
}

// CircuitDesignDir
type CircuitDesignDir struct{}

func (CircuitDesignDir) Materialize(matter *be.Matter) be.Reflex {
	return be.NewNounReflex(matter.Circuit.SourceDir)
}
