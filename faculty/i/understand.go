// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package i

import (
	"github.com/gocircuit/escher/plumb"
	es "github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/be"
	eu "github.com/gocircuit/escher/understand"
)

// Understand
type Understand struct{}

func (Understand) Materialize() be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(
		func(eye *plumb.Eye, dvalve string, dvalue interface{}) {
			if dvalve != "Seen" {
				return
			}
			switch t := dvalue.(type) {
			case *es.Circuit:
				eye.Show("Understood", eu.Understand(t))
			}
			panic("nil or unknown seen")
		}, 
		"Seen", "Understood",
	)
	return reflex
}
