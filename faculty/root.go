// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"sync"

	. "github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/see"
)

var lk sync.Mutex
var root = NewIdiom()

func Root() Idiom {
	lk.Lock()
	defer lk.Unlock()
	return root
}

func Register(name string, v interface{}) {
	lk.Lock()
	defer lk.Unlock()
	root.Memorize(v, ParseAddress(name).Path...)
}
