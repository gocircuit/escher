// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/see"
)

// Fix creates a gate that waits until all fix valves are set and
// then outputs a singular conjunction of all values.
func MaterializeFix(fwd string, fix ...string) be.Reflex {
	reflex, eye := plumb.NewEye(append(fix, fwd)...)
	go func() {
		conj := Make()
		for {
			dvalve, dvalue := eye.See()
			if dvalve == fwd { // conjunction updated
				continue // ignore upstream updates
			} else { // field updated
				conj.Abandon(see.Name(dvalve)).Grow(see.Name(dvalve), dvalue)
				if conj.Len() == len(fix) {
					eye.Show(fwd, conj)
					eye.Drain() // As soon as the conjunction is output, this gate is done.
				}
			}
		}
	}()
	return reflex
}
