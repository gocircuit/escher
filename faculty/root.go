// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/kit/memory"
)

var root = NewReservoir()

func Root() Circuit {
	return root.Copy()
}

func Register(name string, v interface{}) {
	root.Register(name, v)
}
