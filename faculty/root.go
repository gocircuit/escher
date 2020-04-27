// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"sync"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
)

var lk sync.Mutex
var root = be.NewIndex()

func Root() be.Index {
	lk.Lock()
	defer lk.Unlock()
	return root
}

func Register(v be.Materializer, addr ...cir.Name) {
	lk.Lock()
	defer lk.Unlock()
	root.Memorize(v, addr...)
}
