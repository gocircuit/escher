// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	. "github.com/gocircuit/escher/see"
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/kit/reservoir"
)

var root = NewReservoir()

func Root() Reservoir {
	return root
}

func Register(name string, v interface{}) {
	root.Put(ParseAddress(name), v)
}
