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
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	ns := faculty.Root.Refine("escher")
	ns.AddTerminal("CircuitDesignDir", Builtin{})
}

// Builtin
type Builtin struct{}

func (Builtin) Materialize() be.Reflex {
	reflex, _ := plumb.NewEyeCognizer((&builtin{}).Cognize, "_")
	return reflex
}

type builtin struct {}

func (x *builtin) Cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	panic("placeholder for compile-time builtin circuit")
}
