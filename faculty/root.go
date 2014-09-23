// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"strings"

	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/memory"
)

// Root is the global faculties memory where Go packages add gate designs as side-effect of being imported.
var root Memory = NewMemory()

func Root() Memory {
	return root
}

func Register(name string, v interface{}) {
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
		x.IncludeIfNot(g, New())
		x = x.Goto(g)
	}
	if x.Include(a[len(a)-1], v) != nil {
		panic("overwriting builtin")
	}
}
