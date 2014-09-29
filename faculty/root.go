// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"strings"
	"sync"

	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/memory"
)

// Root is the global faculties memory where Go packages add gate designs as side-effect of being imported.
var root = Memory(New())
var lk sync.Mutex

func Root() Memory {
	return root
}

func Register(name string, v interface{}) {
	lk.Lock()
	defer lk.Unlock()
	a := strings.Split(name, ".")
	if len(a) == 0 {
		panic(1)
	}
	//
	x := root
	for i, g := range a {
		if i+1 == len(a) {
			break
		}
		x = x.Refine(g)
	}
	if x.Include(a[len(a)-1], v) != nil {
		panic("overwriting builtin")
	}
}
