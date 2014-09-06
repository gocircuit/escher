// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package i

import (
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/plumb"
	es "github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/be"
)

func init() {
	ns := faculty.Root.Refine("i")
	ns.AddTerminal(es.Name("See"), See{})
	ns.AddTerminal(es.Name("Understand"), Understand{})
	ns.AddTerminal(es.Name("Memory"), Memory{})
	// ns.AddTerminal(es.Name("Materialize"), Materialize{})
}

// See
type See struct{}

func (See) Materialize() be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(
		func(eye *plumb.Eye, dvalve string, dvalue interface{}) {
			if dvalve != "Source" {
				return
			}
			eye.Show("Seen", es.SeeCircuit(es.NewSrcString(plumb.AsString(dvalue))))
		}, 
		"Source", "Seen",
	)
	return reflex
}
