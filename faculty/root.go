// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"sync"

	. "github.com/gocircuit/escher/see"
	. "github.com/gocircuit/escher/circuit"
)

var lk sync.Mutex
var root = New()

func Root() Circuit {
	lk.Lock()
	defer lk.Unlock()
	return root.DeepCopy()
}

func Register(name string, v interface{}) {
	lk.Lock()
	defer lk.Unlock()
	root.Place(ParseAddress(name), v)
}
