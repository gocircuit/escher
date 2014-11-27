// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package path

import (
	// "fmt"
	"path"
	"sync"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("path.Join", Join{})
}

// Join
type Join struct{ Sparkless }

func (Join) Materialize(*be.Matter) (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer((&join{}).Cognize, DefaultValve, "Head", "Tail")
	return reflex, Join{}
}

type join struct {
	sync.Mutex
	head *string
	tail *string
}

func (x *join) Cognize(eye *be.Eye, dvalve Name, dvalue interface{}) {
	x.Lock()
	defer x.Unlock()
	switch dvalve {
	case "Head":
		head := dvalue.(string)
		x.head = &head
	case "Tail":
		tail := dvalue.(string)
		x.tail = &tail
	default:
		return
	}
	if x.head == nil || x.tail == nil {
		return
	}
	eye.Show(DefaultValve, path.Join(*x.head, *x.tail))
}
